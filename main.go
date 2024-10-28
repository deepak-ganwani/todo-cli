package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("E:/Code/Dev/CLI_ToDo/todos.json")
	cmdFlags := NewCmdFlags()

	storage.Load(&todos)
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
