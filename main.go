package main

import (
	"fmt"
	"os"

	list "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

type Task struct {
	status status
	title string
	description string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

type Model struct {
	list list.Model 
	err error
}

// TODO: call this on tea.WindowSizeMsg
func (m *Model) initList(width, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(),width,height)
	m.list.Title = "Todo List"
	m.list.SetItems([]list.Item{
		Task{status: todo, title: "coding", description: "write some code"},
		Task{status: todo, title: "reading", description: "read some books"},
		Task{status: todo, title: "running", description: "run 5km"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg: m.initList(msg.Width, msg.Height)
	}
	var cmd tea.Cmd	
	m.list,cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View()
}

func main() {
	p := tea.NewProgram(Model{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}