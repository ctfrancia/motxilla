package main

import (
	"fmt"
	// "log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/huh"
	td "github.com/ctfrancia/motxilla/internal/capabilities/todo"
)

type app struct {
	selectedItem       string
	capabilities       []string
	selectedCapability string
	runs               []runbook

	todos      []todo
	deleteTodo bool
	createTodo bool
}

type todo struct {
	Name string
	Done bool
}

type runbook struct {
	cmd string
}

func main() {
	app := app{
		capabilities:       fetchWhatMotxillaCanDo(),
		selectedCapability: "",
	}

	/*
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Options(huh.NewOptions(app.capabilities...)...).
					Value(&app.selectedCapability).Title("What do you want to do?"),
			),
		)

		err := form.Run()
		if err != nil {
			fmt.Println("state", app.selectedCapability)
			log.Fatal(err)
		}
	*/

	// Where the progam starts based on the selected capability
	switch app.selectedCapability {
	case "todos":
		m := td.GetTodos()
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}

	case "runbook":
		fmt.Println("You selected runbook")
		os.Exit(0)

	default:
		fmt.Println("work in progress:", app.selectedCapability)
		os.Exit(1)
	}
}

func fetchWhatMotxillaCanDo() []string {
	// fetch database tables
	return []string{"todos", "runbook", "tbd"}
}

func fetchTodos() []todo {
	// this represents something we would get back from a database
	tds := []todo{
		{Name: "Buy milk", Done: false},
		{Name: "Buy eggs", Done: false},
		{Name: "Buy bread", Done: false},
		{Name: "Buy things", Done: false},
	}
	return tds
}

func fetchRunbook() []runbook {
	// this represents something we would get back from a database
	return []runbook{
		{cmd: "ls"},
		{cmd: "pwd"},
		{cmd: "cd /Users/ctfrancia/Projects/motxilla"},
	}
}
