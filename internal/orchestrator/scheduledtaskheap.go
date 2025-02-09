package orchestrator

import (
	"container/heap"
	"context"
	"sync"
	"time"
)

type taskQueue struct {
	dequeueMu sync.Mutex
	mu        sync.Mutex
	heap      scheduledTaskHeapByTime
	notify    chan struct{}
	ready     scheduledTaskHeapByPriorityThenTime

	Now func() time.Time
}

func (t *taskQueue) curTime() time.Time {
	if t.Now != nil {
		return t.Now()
	}
	return time.Now()
}

func (t *taskQueue) Push(task scheduledTask) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if task.task == nil {
		panic("task cannot be nil")
	}

	heap.Push(&t.heap, &task)
	if t.notify != nil {
		t.notify <- struct{}{}
	}
}

func (t *taskQueue) Reset() []*scheduledTask {
	t.mu.Lock()
	defer t.mu.Unlock()

	oldTasks := t.heap.tasks
	oldTasks = append(oldTasks, t.ready.tasks...)
	t.heap.tasks = nil
	t.ready.tasks = nil

	if t.notify != nil {
		t.notify <- struct{}{}
	}
	return oldTasks
}

func (t *taskQueue) Dequeue(ctx context.Context) *scheduledTask {
	t.dequeueMu.Lock()
	defer t.dequeueMu.Unlock()

	t.notify = make(chan struct{}, 1)
	defer func() {
		t.notify = nil
	}()

	t.mu.Lock()
	for {
		first, ok := t.heap.Peek().(*scheduledTask)
		if !ok { // no tasks in heap.
			if t.ready.Len() > 0 {
				t.mu.Unlock()
				return heap.Pop(&t.ready).(*scheduledTask)
			}
			t.mu.Unlock()
			select {
			case <-ctx.Done():
				return nil
			case <-t.notify:
			}
			t.mu.Lock()
			continue
		}

		now := t.curTime()

		// if there's a task in the ready queue AND the first task isn't ready yet then immediately return the ready task.
		ready, ok := t.ready.Peek().(*scheduledTask)
		if ok && now.Before(first.runAt) {
			heap.Pop(&t.ready)
			t.mu.Unlock()
			return ready
		}

		t.mu.Unlock()
		timer := time.NewTimer(first.runAt.Sub(now))

		select {
		case <-timer.C:
			t.mu.Lock()
			if t.heap.Len() == 0 {
				break
			}

			for {
				first, ok := t.heap.Peek().(*scheduledTask)
				if !ok {
					break
				}
				if first.runAt.After(t.curTime()) {
					// task is not yet ready to run
					break
				}
				heap.Pop(&t.heap) // remove the task from the heap
				heap.Push(&t.ready, first)
			}

			if t.ready.Len() == 0 {
				break
			}
			t.mu.Unlock()
			return heap.Pop(&t.ready).(*scheduledTask)
		case <-t.notify: // new task was added, loop again to ensure we have the earliest task.
			t.mu.Lock()
			if !timer.Stop() {
				<-timer.C
			}
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return nil
		}
	}
}

type scheduledTask struct {
	task     Task
	runAt    time.Time
	priority int
}

type scheduledTaskHeap struct {
	tasks      []*scheduledTask
	comparator func(i, j *scheduledTask) bool
}

func (h *scheduledTaskHeap) Len() int {
	return len(h.tasks)
}

func (h *scheduledTaskHeap) Swap(i, j int) {
	h.tasks[i], h.tasks[j] = h.tasks[j], h.tasks[i]
}

func (h *scheduledTaskHeap) Push(x interface{}) {
	h.tasks = append(h.tasks, x.(*scheduledTask))
}

func (h *scheduledTaskHeap) Pop() interface{} {
	old := h.tasks
	n := len(old)
	x := old[n-1]
	h.tasks = old[0 : n-1]
	return x
}

func (h *scheduledTaskHeap) Peek() interface{} {
	if len(h.tasks) == 0 {
		return nil
	}
	return h.tasks[0]
}

type scheduledTaskHeapByTime struct {
	scheduledTaskHeap
}

var _ heap.Interface = &scheduledTaskHeapByTime{}

func (h *scheduledTaskHeapByTime) Less(i, j int) bool {
	return h.tasks[i].runAt.Before(h.tasks[j].runAt)
}

type scheduledTaskHeapByPriorityThenTime struct {
	scheduledTaskHeap
}

var _ heap.Interface = &scheduledTaskHeapByPriorityThenTime{}

func (h *scheduledTaskHeapByPriorityThenTime) Less(i, j int) bool {
	if h.tasks[i].priority != h.tasks[j].priority {
		return h.tasks[i].priority > h.tasks[j].priority
	}
	return h.tasks[i].runAt.Before(h.tasks[j].runAt)
}
