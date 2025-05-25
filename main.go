// TO-DO LIST USING BUBBLETEA

package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/textinput"
)

type model struct {
    tasks       []string 
    cursor      int 
    checked     map[int]struct{}
    addMode     bool
    textInput   textinput.Model
    err error
}

func initialModel() model {
    ti := textinput.New()
    ti.Placeholder = "Enter text here"
    ti.Focus()
    ti.CharLimit = 50
    ti.Width = 40

    return model {
        tasks: []string{},
        cursor: 0,
        checked: make(map[int]struct{}),
        addMode: false,
        textInput: ti,
        err: nil,
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}

func (m model) View() string {

}

func main() {
    p := tea.NewProgram(initialModel())
    if err := p.Start; err != nil {
        fmt.Println("Error Running Program: &v", err)
        os.Exit(1)
    }
}
