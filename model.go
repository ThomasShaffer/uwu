package main

import (
    "os/exec"
    tea     "github.com/charmbracelet/bubbletea"
    lipgloss "github.com/charmbracelet/lipgloss"
)

type model struct {
    todos   []entry
    code    code
    cursor  int
}

func InitializeModel() model {
    return model {
        todos:  GetTodos(),
        code:   GetCode(),
        cursor: 0,
    }
}

func (m model) OpenEditor(filepath string, line string) tea.Cmd {
    cmd := exec.Command(config.Editor,  "+"+line, filepath)
    return tea.ExecProcess(cmd, func(err error) tea.Msg {
            return err
    })
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.code.SetFile(m.todos[m.cursor])
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl-c", "q":
            return m, tea.Quit
        case "j":
            if m.cursor != len(m.todos)-1 {
                m.cursor++
            }
            m.code.SetFile(m.todos[m.cursor])
            return m, nil
        case "k":
            if m.cursor != 0 {
                m.cursor--
            }
            m.code.SetFile(m.todos[m.cursor])
            return m, nil
        case "enter":
            return m, m.OpenEditor(m.todos[m.cursor].filepath, m.todos[m.cursor].line)
        case "R":
            m.todos = GetTodos()
            return m, nil
        }
    }
    return m, nil
}

//TODO: Make the TUI prettier. Pretty ugly 
func (m model) View() string {
    todos := "\n TODOS\n\n" 
    entries := FormatTodos(m.todos)
    for entry := range entries {
        todo := entries[entry]
        if entry == m.cursor {
            todos += ">"
        } else {
            todos += " "
        }
        todos += todo
    }
    return lipgloss.JoinVertical(lipgloss.Top, todos, m.code.View())
}
