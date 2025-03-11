package main

import (
	"fmt"
)

type Task struct {
	id           int
	name         string
	cost         int
	taskType     string
	taskCallback func(input ...any) error
	status       string
}

func execute(task *Task) {

	task.status = "pending"

	fmt.Printf("executing %s task %d\n", task.taskType, task.id)

	for i := task.cost; i > 0; i-- {
	}

	task.status = "completed"

}

func main() {

	tasks := []Task{
		{
			id:       1,
			name:     "task1",
			cost:     100000000,
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
		{
			id:       3,
			name:     "task3",
			cost:     10,
			taskType: "async",
			taskCallback: func(input ...any) error {
				fmt.Println("hello from task 3 callback")
				return nil
			},
		},
	}

	scheduledTasks := []*Task{}

	for _, task := range tasks {
		if task.taskType == "async" {
			scheduledTasks = append(scheduledTasks, &task)
			go execute(&task)
		} else {
			execute(&task)
		}
	}

	for {
		taskTop := scheduledTasks[0]
		scheduledTasks = scheduledTasks[1:]

		if taskTop.status == "completed" {
			taskTop.taskCallback()
		} else {
			scheduledTasks = append(scheduledTasks, taskTop)
		}

		if len(scheduledTasks) == 0 {
			break
		}

	}

}
