package main

import "fmt"

type Task struct {
	id           int
	name         string
	cost         int
	taskType     string
	taskCallback func(input ...any) error
}

type Callback struct {
	callbackTaskId int
	callbackFunc   func(input ...any) error
}

func execute(task Task, notifiers ...chan int) {
	fmt.Printf("executing %s task %d\n", task.taskType, task.id)

	for i := task.cost; i > 0; i-- {
	}

	if task.taskType == "async" && notifiers != nil && len(notifiers) > 0 {
		notifiers[0] <- 1
		close(notifiers[0])
	}

}

func main() {

	tasks := []Task{
		{
			id:       1,
			name:     "task1",
			cost:     10,
			taskType: "async",
			taskCallback: func(input ...any) error {
				fmt.Println("hello from task 1 callback")
				return nil
			},
		},
		{
			id:       2,
			name:     "task2",
			cost:     10,
			taskType: "sync",
		},
	}

	callbackQueue := []Callback{}
	notifiers := make([]chan int, 0)

	for _, task := range tasks {
		if task.taskType == "async" {
			callbackQueue = append(callbackQueue, Callback{
				callbackTaskId: task.id,
				callbackFunc:   task.taskCallback,
			})
			notifier := make(chan int)
			notifiers = append(notifiers, notifier)
			go execute(task, notifier)
		} else {

			execute(task)
		}
	}

	for _, notifier := range notifiers {
		completedTaskid := <-notifier
		var callback Callback
		for _, cb := range callbackQueue {
			if cb.callbackTaskId == completedTaskid {
				callback = cb
				break
			}
		}
		callback.callbackFunc()
	}

}
