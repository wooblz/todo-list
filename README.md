Note: got chatgpt to write this read me...... because i didn't feel like it

Todo List CLI
A command-line application for managing to-do tasks, written in Go. This tool allows users to create, list, complete, and delete tasks stored locally in a JSON file.

Features
  Add new tasks via command line
  
  List pending or all tasks
  
  Mark tasks as completed
  
  Remove tasks by ID
  
  Persist data using a local JSON file
  
  Organized with subcommands using Cobra

Installation
git clone https://github.com/wooblz/todo-list.git
cd todo-list
go build -o tasks main.go
This will compile the CLI tool and produce an executable named tasks.

Usage
  ./tasks add "Read a book"
  ./tasks list
  ./tasks complete 2
  ./tasks delete 3
  ./tasks list --all
Available Commands
  add <description>: Adds a new task with the given description.
  
  list: Lists all incomplete tasks. Use --all to include completed tasks.
  
  complete <id>: Marks the task with the specified ID as complete.
  
  delete <id>: Deletes the task with the specified ID.

File Structure
  cmd/          // Cobra command implementations
  file/         // JSON file read/write logic
  model/        // Task struct and methods
  main.go       // Application entry point
Testing
To run tests for file and model functionality:
  go test ./file ./model
License
This project is licensed under the MIT License
