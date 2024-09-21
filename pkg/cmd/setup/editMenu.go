package setup

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/scorify/scorify/pkg/config"
	"github.com/sirupsen/logrus"
)

type editItem struct {
	label   string
	value   string
	private bool
	prev    string
}

type editModel struct {
	itemCursor int
	editCursor int
	items      []editItem
	editting   bool
}

var selected = color.New(color.FgBlack, color.BgWhite).SprintFunc()

func newEditModel() editModel {
	return editModel{
		items: []editItem{
			{label: "Domain", value: config.Domain, prev: config.Domain},
			{label: "Port", value: fmt.Sprintf("%d", config.Port), prev: fmt.Sprintf("%d", config.Port)},
			{label: "Interval", value: config.IntervalStr, prev: config.IntervalStr},
			{label: "JWT Timeout", value: config.JWT.TimeoutStr, prev: config.JWT.TimeoutStr},
			{label: "JWT Secret", value: config.JWT.Secret, private: true, prev: config.JWT.Secret},
			{label: "Postgres Host", value: config.Postgres.Host, prev: config.Postgres.Host},
			{label: "Postgres Port", value: fmt.Sprintf("%d", config.Postgres.Port), prev: fmt.Sprintf("%d", config.Postgres.Port)},
			{label: "Postgres User", value: config.Postgres.User, prev: config.Postgres.User},
			{label: "Postgres Password", value: config.Postgres.Password, private: true, prev: config.Postgres.Password},
			{label: "Postgres DB", value: config.Postgres.DB, prev: config.Postgres.DB},
			{label: "Redis Host", value: config.Redis.Host, prev: config.Redis.Host},
			{label: "Redis Port", value: fmt.Sprintf("%d", config.Redis.Port), prev: fmt.Sprintf("%d", config.Redis.Port)},
			{label: "Redis Password", value: config.Redis.Password, private: true, prev: config.Redis.Password},
			{label: "RabbitMQ Host", value: config.RabbitMQ.Host, prev: config.RabbitMQ.Host},
			{label: "RabbitMQ Port", value: fmt.Sprintf("%d", config.RabbitMQ.Port), prev: fmt.Sprintf("%d", config.RabbitMQ.Port)},
			{label: "RabbitMQ Admin User", value: config.RabbitMQ.Server.User, prev: config.RabbitMQ.Server.User},
			{label: "RabbitMQ Admin Password", value: config.RabbitMQ.Server.Password, private: true, prev: config.RabbitMQ.Server.Password},
			{label: "RabbitMQ Minion User", value: config.RabbitMQ.Minion.User, prev: config.RabbitMQ.Minion.User},
			{label: "RabbitMQ Minion Password", value: config.RabbitMQ.Minion.Password, private: true, prev: config.RabbitMQ.Minion.Password},
		},
	}
}

func (m editModel) Init() tea.Cmd {
	return nil
}

func (m editModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if !m.editting {
				m.itemCursor = max(m.itemCursor-1, 0)
			}
		case "down":
			if !m.editting {
				m.itemCursor = min(m.itemCursor+1, len(m.items))
			}
		case "left":
			if len(m.items) == m.itemCursor {
				m.editCursor = 0
			} else {
				if m.editting {
					m.editCursor = max(m.editCursor-1, 0)
				}
			}
		case "right":
			if len(m.items) == m.itemCursor {
				m.editCursor = 1
			} else {
				if m.editting {
					m.editCursor = min(m.editCursor+1, len(m.items[m.itemCursor].value))
				}
			}
		case "enter":
			m.editting = !m.editting
			if len(m.items) == m.itemCursor {
				if m.editCursor == 0 {
					return m, tea.Quit
				} else {
					err := saveConfig(m)
					if err != nil {
						logrus.WithError(err).Fatal("failed to save config")
					}
					return m, tea.Quit
				}
			} else {
				m.editCursor = len(m.items[m.itemCursor].value)
			}
		case "ctrl+c":
			return m, tea.Quit
		case "backspace", "delete":
			if m.editting && m.editCursor != len(m.items) {
				if m.editCursor < len(m.items[m.itemCursor].value) {
					m.items[m.itemCursor].value = m.items[m.itemCursor].value[:m.editCursor-1] + m.items[m.itemCursor].value[m.editCursor:]
				} else if m.editCursor == len(m.items[m.itemCursor].value) {
					m.items[m.itemCursor].value = m.items[m.itemCursor].value[:m.editCursor-1]
				}
				m.editCursor--
			}
		case "esc":
			if m.editting {
				m.editting = false
			} else {
				return m, tea.Quit
			}
		default:
			if m.editting && len(m.items) != m.editCursor {
				m.items[m.itemCursor].value = m.items[m.itemCursor].value[:m.editCursor] + msg.String() + m.items[m.itemCursor].value[m.editCursor:]
				m.editCursor++
			} else {
				switch msg.String() {
				case "j":
					m.itemCursor = min(m.itemCursor+1, len(m.items))
				case "k":
					m.itemCursor = max(m.itemCursor-1, 0)
				case "q":
					return m, tea.Quit
				}
			}
		}
	}
	return m, nil
}

func (m editModel) View() string {
	s := ""
	for i, item := range m.items {
		prefix := " "
		if item.value != item.prev {
			prefix = "+"
		}
		if i == m.itemCursor {
			if m.editting {
				prefix = "*"
			} else {
				prefix = ">"
			}
		}

		if m.editting && i == m.itemCursor {
			if m.editCursor == len(item.value) {
				s += fmt.Sprintf("%s %s: %s%s\n", prefix, item.label, item.value, selected(" "))
			} else if m.editCursor == len(item.value)-1 {
				s += fmt.Sprintf("%s %s: %s\n", prefix, item.label, item.value[:m.editCursor]+selected(string(item.value[m.editCursor])))
			} else {
				s += fmt.Sprintf("%s %s: %s\n", prefix, item.label, item.value[:m.editCursor]+selected(string(item.value[m.editCursor]))+item.value[m.editCursor+1:])
			}
		} else {
			if item.private {
				s += fmt.Sprintf("%s %s: %s\n", prefix, item.label, strings.Repeat("*", len(item.value)))
			} else {
				s += fmt.Sprintf("%s %s: %s\n", prefix, item.label, item.value)
			}
		}
	}
	exit := " EXIT "
	save := " SAVE "

	if len(m.items) == m.itemCursor {
		if m.editCursor == 0 {
			exit = selected(exit)
		} else {
			save = selected(save)
		}
	}

	s += fmt.Sprintf("\n[%s] [%s]\n", exit, save)
	return s
}

func editMenu() error {
	p := tea.NewProgram(newEditModel())
	_, err := p.Run()
	return err
}

func saveConfig(m editModel) error {
	domain := m.items[0].value
	port, err := strconv.Atoi(m.items[1].value)
	if err != nil {
		return err
	}
	interval, err := time.ParseDuration(m.items[2].value)
	if err != nil {
		return err
	}
	jwtTimeout, err := time.ParseDuration(m.items[3].value)
	if err != nil {
		return err
	}
	jwtSecret := m.items[4].value
	postgresHost := m.items[5].value
	postgresPort, err := strconv.Atoi(m.items[6].value)
	if err != nil {
		return err
	}
	postgresUser := m.items[7].value
	postgresPassword := m.items[8].value
	postgresDB := m.items[9].value
	redisHost := m.items[10].value
	redisPort, err := strconv.Atoi(m.items[11].value)
	if err != nil {
		return err
	}
	redisPassword := m.items[12].value
	rabbitMQHost := m.items[13].value
	rabbitMQPort, err := strconv.Atoi(m.items[14].value)
	if err != nil {
		return err
	}
	rabbitMQAdminUser := m.items[15].value
	rabbitMQAdminPassword := m.items[16].value
	rabbitMQMinionUser := m.items[17].value
	rabbitMQMinionPassword := m.items[18].value

	return writeConfig(
		domain,
		port,
		interval,
		jwtTimeout,
		jwtSecret,
		postgresHost,
		postgresPort,
		postgresUser,
		postgresPassword,
		postgresDB,
		redisHost,
		redisPort,
		redisPassword,
		rabbitMQHost,
		rabbitMQPort,
		rabbitMQAdminUser,
		rabbitMQAdminPassword,
		rabbitMQMinionUser,
		rabbitMQMinionPassword,
	)
}
