package main

import (
	"os"
	"testing"
)

func TestLoadTasks(t *testing.T) {
	tasks, err := loadTasks("nonexistent.json")
	if err != nil {
		t.Errorf("Expected no error for nonexistent file, got %v", err)
	}
	if len(tasks.Tasks) != 0 {
		t.Errorf("Expected empty task list, got %v", tasks.Tasks)
	}
}

func TestSaveTasks(t *testing.T) {
	tasks := TaskList{
		Tasks: []Task{{ID: 1, Title: "Test", Done: false}},
	}
	err := saveTasks(tasks, "test.json")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	defer os.Remove("test.json")
	loaded, err := loadTasks("test.json")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(loaded.Tasks) != 1 || loaded.Tasks[0].Title != "Test" {
		t.Errorf("Expected one task 'Test', got %v", loaded.Tasks)
	}
}
