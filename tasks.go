package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func loadTasks(filename string) (TaskList, error) {
	var taskList TaskList
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return TaskList{}, nil
		}
		return TaskList{}, err

	}
	err = json.Unmarshal(data, &taskList)
	return taskList, err
}

func saveTasks(tasks TaskList, filename string) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
