package main

import (
    tea "github.com/charmbracelet/bubbletea"
)

func main() {
    p := tea.NewProgram(InitializeModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        panic(err)
    }
}
