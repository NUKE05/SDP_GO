package main

import "fmt"

type Command interface {
	Execute()
}

type Task struct {
	task string
}

func (t *Task) Execute() {
	fmt.Printf("Выполнить задачу: %s\n", t.task)
}

type TaskScheduler struct {
	tasks []Command
}

func (scheduler *TaskScheduler) AddTask(task Command) {
	scheduler.tasks = append(scheduler.tasks, task)
}

func (scheduler *TaskScheduler) RunTasks() {
	for _, task := range scheduler.tasks {
		task.Execute()
	}
}

func main() {
	// Создание задач на день
	task1 := &Task{task: "Закончить Assignment по SDP"}
	task2 := &Task{task: "Сделать ревью квиза по СДП"}
	task3 := &Task{task: "Приготовиться к рубежке по СДП"}

	scheduler := &TaskScheduler{}

	// Добавление задач на день в планировщик
	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)

	scheduler.RunTasks()
}
