**CLI Todo List**

This is a simple command-line todo list application built using Go. 

**Getting Started**

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/deepak-ganwani/todo-cli.git
   cd todo-cli
   ```

2. **Initialize Go Modules:**
   ```bash
   go mod init todo
   ```
   This command initializes the `go.mod` file, which manages dependencies for your Go project.

3. **Build the Application:**
   ```bash
   go build -o todo.exe
   ```
   This command builds the application and creates an executable file named `todo.exe`.

**Running the Application**
Once built, you can run the `todo.exe` file from any directory:
```bash
todo -<command>
```

**Available Commands:**

* **`-help`:** Lists all Commands.
* **`-add <task>`:** Adds a new task to the list.
* **`-del <id>`:** Deletes a task by its ID.
* **`-edit <id>:<new_task>`:** Edits an existing task.
* **`-list`:** Lists all tasks.
* **`-toggle <id>`:** Toggles the completion status of a task.
* ![image](https://github.com/user-attachments/assets/6aeead60-7485-4172-aabe-8b50301fa4a4)

**Example Usage:**
```bash
todo -add "Buy groceries"
todo -add "Learn Go"
todo -toggle 1
todo -list
```
* ![image](https://github.com/user-attachments/assets/ecaf6d39-087c-431f-8fc2-49e6881132b8)

**How it Works**

The `storage.go` file handles the persistence of todo items to a JSON file named `todos.json`. The `todo.go` file defines the `Todo` struct and related functions. The `command.go` file implements the command-line interface, parsing arguments and executing the corresponding actions.

**Additional Notes**

* The `go.sum` file is automatically generated by the `go mod` tool and records the exact versions of dependencies used in your project. This ensures consistent builds across different environments.
* For more advanced usage and customization, feel free to explore the source code. 

**Contributing**
Contributions are welcome! Feel free to fork the repository, make changes, and submit a pull request.

**License**
This project is licensed under the MIT License.
