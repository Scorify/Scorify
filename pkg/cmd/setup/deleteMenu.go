package setup

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type deleteActionChoice string

const (
	deleteYes deleteActionChoice = "yes"
	deleteNo  deleteActionChoice = "no"
)

type deleteItem struct {
	choice      deleteActionChoice
	label       string
	description string
}

type deleteModel struct {
	cursor int
	items  []deleteItem
	action *deleteActionChoice
}

func newDeleteModel(action *deleteActionChoice) deleteModel {
	return deleteModel{
		cursor: 0,
		items: []deleteItem{
			{choice: deleteYes, label: "Yes", description: "Delete the configuration"},
			{choice: deleteNo, label: "No", description: "Do not delete the configuration"},
		},
		action: action,
	}
}

func (m deleteModel) Init() tea.Cmd {
	return nil
}

func (m deleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			*m.action = deleteNo
			return m, tea.Quit
		case "up", "k":
			m.cursor = max(m.cursor-1, 0)
		case "down", "j":
			m.cursor = min(m.cursor+1, len(m.items)-1)
		case "enter":
			*m.action = m.items[m.cursor].choice
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m deleteModel) View() string {
	s := "Are you sure you want to delete the configuration?\n\n"

	for i, item := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s: %s\n", cursor, item.label, item.description)
	}

	s += "\nPress 'Enter' to confirm or 'q' to cancel"

	return s
}

func deleteMenu() error {
	action := deleteNo
	m := newDeleteModel(&action)

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}

	if *m.action == deleteYes {
		return os.Remove(".env")
	}

	return nil
}
