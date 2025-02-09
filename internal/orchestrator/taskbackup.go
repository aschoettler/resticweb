package orchestrator

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/garethgeorge/resticui/gen/go/v1"
	"github.com/garethgeorge/resticui/internal/protoutil"
	"github.com/garethgeorge/resticui/pkg/restic"
	"github.com/gitploy-io/cronexpr"
	"go.uber.org/zap"
)

// BackupTask is a scheduled backup operation.
type BackupTask struct {
	name string
	TaskWithOperation
	plan      *v1.Plan
	scheduler func(curTime time.Time) *time.Time
}

var _ Task = &BackupTask{}

func NewScheduledBackupTask(orchestrator *Orchestrator, plan *v1.Plan) (*BackupTask, error) {
	sched, err := cronexpr.ParseInLocation(plan.Cron, time.Now().Location().String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse schedule %q: %w", plan.Cron, err)
	}

	return &BackupTask{
		name: fmt.Sprintf("backup for plan %q", plan.Id),
		TaskWithOperation: TaskWithOperation{
			orch: orchestrator,
		},
		plan: plan,
		scheduler: func(curTime time.Time) *time.Time {
			next := sched.Next(curTime)
			return &next
		},
	}, nil
}

func NewOneofBackupTask(orchestrator *Orchestrator, plan *v1.Plan, at time.Time) *BackupTask {
	didOnce := false
	return &BackupTask{
		name: fmt.Sprintf("onetime backup for plan %q", plan.Id),
		TaskWithOperation: TaskWithOperation{
			orch: orchestrator,
		},
		plan: plan,
		scheduler: func(curTime time.Time) *time.Time {
			if didOnce {
				return nil
			}
			didOnce = true
			return &at
		},
	}
}

func (t *BackupTask) Name() string {
	return t.name
}

func (t *BackupTask) Next(now time.Time) *time.Time {
	next := t.scheduler(now)
	if next == nil {
		return nil
	}

	if err := t.setOperation(&v1.Operation{
		PlanId:          t.plan.Id,
		RepoId:          t.plan.Repo,
		UnixTimeStartMs: timeToUnixMillis(*next),
		Status:          v1.OperationStatus_STATUS_PENDING,
		Op:              &v1.Operation_OperationBackup{},
	}); err != nil {
		zap.S().Errorf("task %v failed to add operation to oplog: %v", t.Name(), err)
	}

	return next
}

func (t *BackupTask) Run(ctx context.Context) error {
	return t.runWithOpAndContext(ctx, func(ctx context.Context, op *v1.Operation) error {
		return backupHelper(ctx, t.orch, t.plan, op)
	})
}

// backupHelper does a backup.
func backupHelper(ctx context.Context, orchestrator *Orchestrator, plan *v1.Plan, op *v1.Operation) error {
	startTime := time.Now()
	backupOp := &v1.Operation_OperationBackup{
		OperationBackup: &v1.OperationBackup{},
	}
	op.Op = backupOp

	zap.L().Info("Starting backup", zap.String("plan", plan.Id), zap.Int64("opId", op.Id))
	repo, err := orchestrator.GetRepo(plan.Repo)
	if err != nil {
		return fmt.Errorf("couldn't get repo %q: %w", plan.Repo, err)
	}

	lastSent := time.Now() // debounce progress updates, these can endup being very frequent.
	summary, err := repo.Backup(ctx, plan, func(entry *restic.BackupProgressEntry) {
		if time.Since(lastSent) < 250*time.Millisecond {
			return
		}
		lastSent = time.Now()

		backupOp.OperationBackup.LastStatus = protoutil.BackupProgressEntryToProto(entry)
		if err := orchestrator.OpLog.Update(op); err != nil {
			zap.S().Errorf("failed to update oplog with progress for backup: %v", err)
		}
	})
	if err != nil {
		return fmt.Errorf("repo.Backup for repo %q: %w", plan.Repo, err)
	}

	op.SnapshotId = summary.SnapshotId
	backupOp.OperationBackup.LastStatus = protoutil.BackupProgressEntryToProto(summary)
	if backupOp.OperationBackup.LastStatus == nil {
		return fmt.Errorf("expected a final backup progress entry, got nil")
	}

	zap.L().Info("Backup complete", zap.String("plan", plan.Id), zap.Duration("duration", time.Since(startTime)), zap.Any("summary", summary))

	// schedule followup tasks
	at := time.Now()
	if plan.Retention != nil {
		orchestrator.ScheduleTask(NewOneofForgetTask(orchestrator, plan, op.SnapshotId, at), TaskPriorityForget)
	}

	orchestrator.ScheduleTask(NewOneofIndexSnapshotsTask(orchestrator, plan, at), TaskPriorityIndexSnapshots)

	return nil
}
