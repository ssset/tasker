# Tasker

A simple CLI tool for managing tasks, written in Go.

## Installation
```bash
  go install github.com/ssset/tasker@latest
```
# Usage
```bash
tasker -add "Buy milk"  # Add a task
tasker -list           # List all tasks
tasker -done 1         # Mark task with ID 1 as done
tasker -delete 1       # Delete task with ID 1
```

# Features
* Add tasks and save to tasks.json
* List tasks
* Mark tasks as done
* Delete tasks
* Unit tests for JSON persistence

# Testing

Run tests with:
```bash
  go test
```