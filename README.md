# Go CLI To-Do Application

A simple command-line To-Do application built using Go. This project allows users to manage tasks directly from the terminal with persistent storage using JSON files.

## Features

* Add new tasks
* View all tasks
* Mark tasks as completed
* Delete tasks
* Automatically save tasks to a JSON file
* Automatically load tasks when the application starts
* Support for multi-word task names

## Project Structure

```text
todo-app/

├── main.go
├── tasks.go
├── storage.go
├── go.mod
├── .gitignore
└── tasks.json
```

### File Responsibilities

* **main.go** - Application entry point and menu handling
* **tasks.go** - Task-related operations (Add, View, Complete, Delete)
* **storage.go** - JSON file storage and retrieval
* **tasks.json** - Persistent task data (ignored by Git)

## Technologies Used

* Go (Golang)
* JSON Serialization (`encoding/json`)
* File Handling (`os`)
* Buffered Input (`bufio`)
* Git & GitHub

## Concepts Practiced

* Structs
* Slices
* Functions
* Error Handling
* JSON Marshalling and Unmarshalling
* File Operations
* Modular Code Organization
* Git Version Control

## How to Run

### Clone the Repository

```bash
git clone <repository-url>
cd todo-app
```

### Run the Application

```bash
go run .
```

## Sample Usage

```text
1. Add Task
2. View Tasks
3. Complete Task
4. Delete Task
5. Exit
```

Example:

```text
Enter Task:
Learn Go Programming

Task added successfully
```

```text
1. [ ] Learn Go Programming
2. [x] Build CLI Project
```

## Future Improvements

* Task priorities
* Due dates
* Search functionality
* Task categories
* Unit testing
* Custom Go packages
* Cobra CLI support

## Learning Outcome

This project was built as part of learning Go and covers many core backend development concepts including data structures, file persistence, JSON processing, modular code organization, and version control with Git.

---

Built with Go 🚀
