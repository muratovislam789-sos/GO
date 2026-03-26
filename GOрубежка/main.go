package main

import (
	"fmt"
)

type Task struct {
	ID          int
	Description string
	IsDone      bool
}

type TaskList struct {
	Tasks []Task
}

func (tl *TaskList) AddTask(desc string) {
	newID := len(tl.Tasks) + 1
	tl.Tasks = append(tl.Tasks, Task{ID: newID, Description: desc})
	fmt.Println("✅ Задача добавлена!")
}

func (tl *TaskList) ShowTasks() {
	fmt.Println("\n--- Список задач ---")
	for _, t := range tl.Tasks {
		status := "[ ]"
		if t.IsDone {
			status = "[X]"
		}
		fmt.Printf("%d. %s %s\n", t.ID, status, t.Description)
	}
}

func main() {
	myList := TaskList{}
	var choice int
	var text string

	for {
		fmt.Println("\n1. Добавить | 2. Показать | 3. Выход")
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Что нужно сделать? ")
			fmt.Scan(&text)
			myList.AddTask(text)
		case 2:
			myList.ShowTasks()
		case 3:
			fmt.Println("Пока!")
			return
		default:
			fmt.Println("Ошибка выбора")
		}
	}
}
