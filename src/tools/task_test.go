package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestTaks(t *testing.T) {
	scheduler := NewScheduler()

	// 一次性任务
	scheduler.AddOnceTask("once", 2*time.Second, func() {
		fmt.Println("一次性任务:", time.Now().Format("15:04:05"))
	})

	scheduler.AddTask("task", TaskTypeOnce, 2*time.Second, func() {
		fmt.Println("任务:", time.Now().Format("15:04:05"))
	})

	// 重复任务
	scheduler.AddRepeatTask("repeat", 1*time.Second, func() {
		fmt.Println("重复任务:", time.Now().Format("15:04:05"))
	})

	time.Sleep(5 * time.Second)

	// 优雅关闭
	scheduler.Shutdown(3 * time.Second)
	fmt.Println("已关闭")
}
