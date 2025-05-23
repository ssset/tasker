package main

import (
	"flag"
	"fmt"
)

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", 0, "Mark task as done")
	flag.Parse()

	tasks, err := loadTasks("tasks.json")
	if err != nil {
		fmt.Println("Ошибка загрузки задач:", err)
		return
	}

	if *add != "" {
		newTask := Task{
			ID:    len(tasks.Tasks) + 1,
			Title: *add,
			Done:  false,
		}
		tasks.Tasks = append(tasks.Tasks, newTask)
		if err := saveTasks(tasks, "tasks.json"); err != nil {
			fmt.Println("Ошибка при сохранении задач:", err)
			return
		}
		fmt.Println("Добавлена задача", newTask.Title)
		fmt.Println("Все задачи", tasks.Tasks)
	} else if *list {
		if len(tasks.Tasks) == 0 {
			fmt.Println("Нет задач")
		} else {
			for _, task := range tasks.Tasks {
				status := " "
				if task.Done {
					status = "x"
				}
				fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
			}
		}
	} else if *done > 0 {
		for i, task := range tasks.Tasks {
			if task.ID == *done {
				tasks.Tasks[i].Done = true
				if err := saveTasks(tasks, "tasks.json"); err != nil {
					fmt.Println("Ошибка при сохранении задач:", err)
					return
				}
				fmt.Println("Задача отмечена как выполненная", task.Title)
			}
		}
	} else {
		fmt.Println("Используйте -add или -list")
	}

}
