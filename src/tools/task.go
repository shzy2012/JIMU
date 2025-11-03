package tools

import (
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
	timer    *time.Timer
	ticker   *time.Ticker
	stop     chan struct{}
}

type Scheduler struct {
	tasks map[string]*Task
	mu    sync.RWMutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(map[string]*Task),
	}
}

// 添加任务
// taskType: TaskTypeOnce(一次执行) 或 TaskTypeRepeat(重复执行)
// interval: 延迟时间(一次执行) 或 执行间隔(重复执行)
func (s *Scheduler) AddTask(id string, taskType TaskType, interval time.Duration, fn func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 如果任务存在,先取消
	if task, exists := s.tasks[id]; exists {
		s.stopTask(task)
	}

	task := &Task{
		ID:       id,
		Type:     taskType,
		Interval: interval,
		Fn:       fn,
		stop:     make(chan struct{}),
	}

	if taskType == TaskTypeOnce {
		// 一次性任务
		task.timer = time.AfterFunc(interval, func() {
			fn()
			// 执行完后自动从任务列表中移除
			s.mu.Lock()
			delete(s.tasks, id)
			s.mu.Unlock()
		})
	} else {
		// 重复执行任务
		task.ticker = time.NewTicker(interval)
		go func() {
			for {
				select {
				case <-task.ticker.C:
					fn()
				case <-task.stop:
					return
				}
			}
		}()
	}

	s.tasks[id] = task
}

// 添加一次性任务的便捷方法
func (s *Scheduler) AddOnceTask(id string, delay time.Duration, fn func()) {
	s.AddTask(id, TaskTypeOnce, delay, fn)
}

// 添加重复任务的便捷方法
func (s *Scheduler) AddRepeatTask(id string, interval time.Duration, fn func()) {
	s.AddTask(id, TaskTypeRepeat, interval, fn)
}

// 停止单个任务(内部使用,不加锁)
func (s *Scheduler) stopTask(task *Task) {
	if task.Type == TaskTypeOnce {
		if task.timer != nil {
			task.timer.Stop()
		}
	} else {
		if task.ticker != nil {
			task.ticker.Stop()
		}
		close(task.stop)
	}
}

// 移除任务
func (s *Scheduler) RemoveTask(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, exists := s.tasks[id]; exists {
		s.stopTask(task)
		delete(s.tasks, id)
		return true
	}
	return false
}

// 停止所有任务
func (s *Scheduler) StopAll() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.tasks {
		s.stopTask(task)
	}
	s.tasks = make(map[string]*Task)
}

// 获取任务数量
func (s *Scheduler) TaskCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.tasks)
}

// 获取任务信息
func (s *Scheduler) GetTaskInfo(id string) (TaskType, time.Duration, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if task, exists := s.tasks[id]; exists {
		return task.Type, task.Interval, true
	}
	return 0, 0, false
}

// 列出所有任务ID
func (s *Scheduler) ListTasks() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]string, 0, len(s.tasks))
	for id := range s.tasks {
		tasks = append(tasks, id)
	}
	return tasks
}
