package tools

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TaskType int

const (
	TaskTypeOnce   TaskType = iota // 执行一次
	TaskTypeRepeat                 // 重复执行
)

type Task struct {
	ID       string
	Type     TaskType
	Interval time.Duration
	Fn       func()
	cancel   context.CancelFunc
}

type Scheduler struct {
	tasks  map[string]*Task
	mu     sync.RWMutex
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func NewScheduler() *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &Scheduler{
		tasks:  make(map[string]*Task),
		ctx:    ctx,
		cancel: cancel,
	}
}

// 添加任务
func (s *Scheduler) AddTask(id string, taskType TaskType, interval time.Duration, fn func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 停止旧任务
	if task, exists := s.tasks[id]; exists {
		task.cancel()
	}

	taskCtx, cancel := context.WithCancel(s.ctx)
	task := &Task{
		ID:       id,
		Type:     taskType,
		Interval: interval,
		Fn:       fn,
		cancel:   cancel,
	}

	s.tasks[id] = task
	s.wg.Add(1)

	if taskType == TaskTypeOnce {
		go s.runOnce(taskCtx, task)
	} else {
		go s.runRepeat(taskCtx, task)
	}
}

// 添加一次性任务
func (s *Scheduler) AddOnceTask(id string, delay time.Duration, fn func()) {
	s.AddTask(id, TaskTypeOnce, delay, fn)
}

// 添加重复任务
func (s *Scheduler) AddRepeatTask(id string, interval time.Duration, fn func()) {
	s.AddTask(id, TaskTypeRepeat, interval, fn)
}

func (s *Scheduler) runOnce(ctx context.Context, task *Task) {
	defer s.wg.Done()

	timer := time.NewTimer(task.Interval)
	defer timer.Stop()

	select {
	case <-timer.C:
		task.Fn()
		s.mu.Lock()
		delete(s.tasks, task.ID)
		s.mu.Unlock()
	case <-ctx.Done():
	}
}

func (s *Scheduler) runRepeat(ctx context.Context, task *Task) {
	defer s.wg.Done()

	ticker := time.NewTicker(task.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			task.Fn()
		case <-ctx.Done():
			return
		}
	}
}

// 移除任务
func (s *Scheduler) RemoveTask(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, exists := s.tasks[id]; exists {
		task.cancel()
		delete(s.tasks, id)
	}
}

// 优雅关闭
func (s *Scheduler) Shutdown(timeout time.Duration) error {
	s.cancel()

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("shutdown timeout")
	}
}

// 获取任务数量
func (s *Scheduler) TaskCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.tasks)
}
