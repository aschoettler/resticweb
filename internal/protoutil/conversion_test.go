package protoutil

import (
	"testing"

	v1 "github.com/garethgeorge/resticui/gen/go/v1"
	"github.com/garethgeorge/resticui/pkg/restic"
	"google.golang.org/protobuf/proto"
)

func TestSnapshotToProto(t *testing.T) {
	snapshot := &restic.Snapshot{
		Id:       "db155169d788e6e432e320aedbdff5a54cc439653093bb56944a67682528aa52",
		Time:     "2023-11-10T19:14:17.053824063-08:00",
		Tree:     "3e2918b261948e69602ee9504b8f475bcc7cdc4dcec0b3f34ecdb014287d07b2",
		Paths:    []string{"/resticui"},
		Hostname: "pop-os",
		Username: "dontpanic",
		Tags:     []string{},
		Parent:   "",
	}

	want := &v1.ResticSnapshot{
		Id:         "db155169d788e6e432e320aedbdff5a54cc439653093bb56944a67682528aa52",
		UnixTimeMs: 1699672457053,
		Tree:       "3e2918b261948e69602ee9504b8f475bcc7cdc4dcec0b3f34ecdb014287d07b2",
		Paths:      []string{"/resticui"},
		Hostname:   "pop-os",
		Username:   "dontpanic",
		Tags:       []string{},
		Parent:     "",
	}

	got := SnapshotToProto(snapshot)

	if !proto.Equal(want, got) {
		t.Errorf("wanted %+v, got: %+v", want, got)
	}
}

func TestBackupProgressEntryToProto(t *testing.T) {
	cases := []struct {
		name  string
		entry *restic.BackupProgressEntry
		want  *v1.BackupProgressEntry
	}{
		{
			name: "summary",
			entry: &restic.BackupProgressEntry{
				MessageType:         "summary",
				FilesNew:            1,
				FilesChanged:        2,
				FilesUnmodified:     3,
				DirsNew:             4,
				DirsChanged:         5,
				DirsUnmodified:      6,
				DataBlobs:           7,
				TreeBlobs:           8,
				DataAdded:           9,
				TotalFilesProcessed: 10,
				TotalBytesProcessed: 11,
				TotalDuration:       12.0,
				SnapshotId:          "db155169d788e6e432e320aedbdff5a54cc439653093bb56944a67682528aa52",
				PercentDone:         13.0, // should be ignored.
			},
			want: &v1.BackupProgressEntry{
				Entry: &v1.BackupProgressEntry_Summary{
					Summary: &v1.BackupProgressSummary{
						FilesNew:            1,
						FilesChanged:        2,
						FilesUnmodified:     3,
						DirsNew:             4,
						DirsChanged:         5,
						DirsUnmodified:      6,
						DataBlobs:           7,
						TreeBlobs:           8,
						DataAdded:           9,
						TotalFilesProcessed: 10,
						TotalBytesProcessed: 11,
						TotalDuration:       12.0,
						SnapshotId:          "db155169d788e6e432e320aedbdff5a54cc439653093bb56944a67682528aa52",
					},
				},
			},
		},
		{
			name: "status",
			entry: &restic.BackupProgressEntry{
				MessageType: "status",
				PercentDone: 13.0,
				TotalFiles:  14,
				FilesDone:   15,
				TotalBytes:  16,
				BytesDone:   17,
			},
			want: &v1.BackupProgressEntry{
				Entry: &v1.BackupProgressEntry_Status{
					Status: &v1.BackupProgressStatusEntry{
						PercentDone: 13.0,
						TotalFiles:  14,
						FilesDone:   15,
						TotalBytes:  16,
						BytesDone:   17,
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := BackupProgressEntryToProto(c.entry)
			if !proto.Equal(got, c.want) {
				t.Errorf("wanted: %+v, got: %+v", c.want, got)
			}
		})
	}
}
