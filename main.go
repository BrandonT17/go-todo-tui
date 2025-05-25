// TO-DO LIST USING BUBBLETEA

package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	tasks     []string
	cursor    int
	checked   map[int]struct{}
	addMode   bool
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter text here"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 40

	return model{
		tasks:     []string{},
		cursor:    0,
		checked:   make(map[int]struct{}),
		addMode:   false,
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.addMode {
			// Handle text input in add mode
			switch msg.String() {
			case "enter":
				trimmed := strings.TrimSpace(m.textInput.Value())
				if trimmed != "" {
					m.tasks = append(m.tasks, trimmed)
				}
				m.textInput.Reset()
				m.addMode = false
				return m, nil
			case "esc":
				m.textInput.Reset()
				m.addMode = false
				return m, nil
			default:
				var cmd tea.Cmd
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}
		}

		// Normal list interaction
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "n":
			m.addMode = true
			m.textInput.Focus()
			return m, nil
		case "enter", " ":
			if _, ok := m.checked[m.cursor]; ok {
				delete(m.checked, m.cursor)
			} else {
				m.checked[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.addMode {
		return fmt.Sprintf(
			"Add New Task:\n\n%s\n\n[Enter] Add • [Esc] Cancel\n",
			m.textInput.View(),
		)
	}

	s := "\nMY TO-DO LIST\n\n"
	for i, task := range m.tasks {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // current selection
		}

		checked := " " // not checked
		if _, ok := m.checked[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task)
	}

	s += "\n[n] New Task • [↑/↓] Navigate • [Space] Check/Uncheck • [q] Quit\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error Running Program: %v\n", err)
		os.Exit(1)
	}
}

