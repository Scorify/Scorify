package setup

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type actionChoice string

const (
	actionNone   actionChoice = ""
	actionCreate actionChoice = "create"
	actionUpdate actionChoice = "update"
	actionDelete actionChoice = "delete"
	actionView   actionChoice = "view"
)

type actionItem struct {
	action      actionChoice
	label       string
	description string
}

type actionModel struct {
	cursor int
	items  []actionItem
	action *actionChoice
}

func newActionModel(action *actionChoice) actionModel {
	return actionModel{
		cursor: 0,
		items: []actionItem{
			{action: actionCreate, label: "Create", description: "Create a new configuration"},
			{action: actionUpdate, label: "Update", description: "Update existing configuration"},
			{action: actionDelete, label: "Delete", description: "Delete existing configuration"},
			{action: actionView, label: "View", description: "View existing configuration"},
			{action: actionNone, label: "Exit", description: "Cancel and exit"},
		},
		action: action,
	}
}

func (m actionModel) Init() tea.Cmd {
	return nil
}

func (m actionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			*m.action = actionNone
			return m, tea.Quit
		case "up", "left", "k":
			m.cursor = max(m.cursor-1, 0)
		case "down", "right", "j":
			m.cursor = min(m.cursor+1, len(m.items)-1)
		case "enter":
			*m.action = m.items[m.cursor].action
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m actionModel) View() string {
	s := "Select action to setup Scorify Configuration\n\n"

	for i, item := range m.items {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s: %s\n", cursor, item.label, item.description)
	}

	s += "\nPress 'q' to exit or 'enter' to select\n"

	return s
}

func actionMenu() (actionChoice, error) {
	action := actionNone
	m := newActionModel(&action)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return actionNone, err
	}

	return *m.action, nil
}
