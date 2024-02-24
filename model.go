package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type Model struct {
	list list.Model
	help help.Model
}

func NewModel() *Model {
	cmds := GetCommands()
	items := make([]list.Item, 0, len(cmds))

	for _, cmd := range cmds {
		items = append(items, cmd)
	}

	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	list.Title = "My Cmds"
	list.SetShowHelp(false)

	help := help.New()
	help.ShowAll = false

	return &Model{
		list: list,
		help: help,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
				case key.Matches(msg, Keys.Quit):
					return m, tea.Quit
				case key.Matches(msg, Keys.Help):
					m.help.ShowAll = !m.help.ShowAll
				}

		case tea.WindowSizeMsg:
			m.list.SetSize(msg.Width - 2, msg.Height - 10)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.list.View(), m.help.View(Keys))
}