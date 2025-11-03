package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestTaks(t *testing.T) {
	scheduler := NewScheduler()

	// 示例1: 添加一次性任务
	scheduler.AddTask("once-job1", TaskTypeOnce, 2*time.Second, func() {
		fmt.Println("[一次性] Job 1 执行:", time.Now().Format("15:04:05"))
	})

	// 示例2: 使用便捷方法添加一次性任务
	scheduler.AddOnceTask("once-job2", 4*time.Second, func() {
		fmt.Println("[一次性] Job 2 执行:", time.Now().Format("15:04:05"))
	})

	// 示例3: 添加重复执行任务
	scheduler.AddTask("repeat-job1", TaskTypeRepeat, 1*time.Second, func() {
		fmt.Println("[重复] Job 1 执行:", time.Now().Format("15:04:05"))
	})

	// 示例4: 使用便捷方法添加重复任务
	scheduler.AddRepeatTask("repeat-job2", 1500*time.Millisecond, func() {
		fmt.Println("[重复] Job 2 执行:", time.Now().Format("15:04:05"))
	})

	fmt.Println("初始任务数:", scheduler.TaskCount())
	fmt.Println("任务列表:", scheduler.ListTasks())

	// 等待一次性任务执行
	time.Sleep(3 * time.Second)
	fmt.Println("\n3秒后任务数:", scheduler.TaskCount())
	fmt.Println("任务列表:", scheduler.ListTasks())

	// 停止一个重复任务
	time.Sleep(2 * time.Second)
	scheduler.RemoveTask("repeat-job1")
	fmt.Println("\n停止 repeat-job1")

	time.Sleep(3 * time.Second)
	fmt.Println("\n最终任务数:", scheduler.TaskCount())

	// 停止所有任务
	scheduler.StopAll()
	fmt.Println("已停止所有任务")
}
