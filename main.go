package main

import tea "github.com/charmbracelet/bubbletea"

var MainModel = NewModel()

func main() {
	p := tea.NewProgram(MainModel, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}