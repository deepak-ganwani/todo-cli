package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add Todo Item")
	flag.StringVar(&cf.Edit, "edit", "", "Edit Todo Item. Format -> ID:(todo)  ")
	flag.IntVar(&cf.Del, "del", -1, "Delete Todo Item using ID")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle Todo Item using ID ")
	flag.BoolVar(&cf.List, "list", false, "Todo List")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Use formate ID:(Todo)")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Invalid Index")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid Command")
	}

	todos.print()
}
