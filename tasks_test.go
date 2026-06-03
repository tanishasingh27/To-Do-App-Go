package main

import "testing"

func TestCreateTask(t *testing.T) {

	task := createTask("LearnGo")

	if task.Title != "LearnGo" {
		t.Errorf("Expected title 'Learn Go', got '%s' ", task.Title)
	}

	if task.Completed != false {
		t.Errorf("Expected Completed = fasle")
	}
}

func TestMarkComplete(t *testing.T) {
	tasks := []Task{
		{
			Title:     "Learn Go",
			Completed: false,
		},
	}

	tasks = markTaskComplete(tasks, 0)
	if !tasks[0].Completed {
		t.Errorf("Expected task to be completed")
	}
}
