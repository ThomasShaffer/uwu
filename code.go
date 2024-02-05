package main

import (
    tea "github.com/charmbracelet/bubbletea"
    viewport "github.com/charmbracelet/bubbles/viewport"
    lipgloss "github.com/charmbracelet/lipgloss" 
    "os"
    "strings"
    "strconv"
)

type code struct {
    model viewport.Model
    style lipgloss.Style
}

func GetCode() code {
    return code {
        model: viewport.New(0, 20),
        style: lipgloss.NewStyle().Border(lipgloss.NormalBorder()),
    }
}

func readFile(filepath string, line string) string {
    mid, _ := strconv.Atoi(line)
    file, _ := os.ReadFile(filepath)
    lines := string(file)
    split := strings.Split(lines, "\n")
    split[mid-1] = lipgloss.NewStyle().Bold(true).Render(split[mid-1])
    result := split[max(0, mid-3): min(len(split), mid+9)]
    return strings.Join(result,"\n")
}

func (c *code) SetFile(todo entry) {
    c.model.SetContent(c.style.Render(readFile(todo.filepath, todo.line)))
}

func (c code) Init() tea.Cmd {
    return nil
}

func (c code) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return c, nil
}

//TODO: Make this prettier, still very ugly
func (c code) View() string {
    return c.model.View()
}

