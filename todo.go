package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(ind int) error {
	if ind < 0 || ind >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(ind int) error {
	t := *todos

	if err := t.validateIndex(ind); err != nil {
		return err
	}

	*todos = append(t[:ind], t[ind+1:]...)

	return nil
}

func (todos *Todos) toggle(ind int) error {
	t := *todos

	if err := t.validateIndex(ind); err != nil {
		return err
	}

	t[ind].Completed = !t[ind].Completed
	t[ind].CompletedAt = nil

	if t[ind].Completed {
		completionTime := time.Now()
		t[ind].CompletedAt = &completionTime
	}

	return nil
}

func (todos *Todos) edit(ind int, title string) error {
	t := *todos

	if err := t.validateIndex(ind); err != nil {
		return err
	}

	t[ind].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for ind, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			completedAt = t.CompletedAt.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(ind), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
