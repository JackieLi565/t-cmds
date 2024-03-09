package main

import (
	"errors"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	list     list.Model
	help     help.Model
	selected list.Item
	err      error
	modified bool
	copy     bool
}

type initOutput struct{}

func NewModel() *Model {
	cmds := getCommands()
	items := make([]list.Item, 0, len(cmds))

	for _, cmd := range cmds {
		items = append(items, cmd)
	}

	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	list.Title = "My Cmds"
	list.Styles.Title = title
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
		case key.Matches(msg, ModelKeys.Quit):
			if m.modified {
				saveToDisk(m.list.Items())
			}
			return m, tea.Quit

		case key.Matches(msg, ModelKeys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, ModelKeys.NewCmd):
			return NewForm().Update(nil)

		case key.Matches(msg, ModelKeys.Enter):
			selectedItem := m.list.SelectedItem()
			switch item := selectedItem.(type) {
			case CmdItem:
				return NewOutput(item).Update(initOutput{})
			}

		case key.Matches(msg, ModelKeys.Delete):
			m.modified = true
			m.list.RemoveItem(m.list.Index())

		case key.Matches(msg, ModelKeys.CopyCmd):
			selectedItem := m.list.SelectedItem()
			if selectedItem == nil { 
				m.err = errors.New("empty list") 
				return m, nil
			}

			m.selected = selectedItem
			m.copy = true
			switch c := selectedItem.(type) {
			case CmdItem:
				err := clipboard.WriteAll(cmdWithArgs(c.Cmd, c.Args))
				if err != nil {
					m.err = err
				}
			}
		}

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width-2, msg.Height-10)

	case Form:
		var (
			name = msg.inputs[name].Value()
			cmd  = msg.inputs[command].Value()
			args = msg.inputs[arguments].Value()
		)

		if !(isEmpty(name) || isEmpty(cmd)) {
			m.modified = true
			m.list.InsertItem(-1, NewCmd(name, cmd, parseArgs(args)))
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	var (
		position = lipgloss.Left
		main = border.MarginBottom(1).Width(60)
		help =  m.help.View(ModelKeys)
	)

	if m.err != nil {
		m.err = nil
		return lipgloss.JoinVertical(position, main.BorderForeground(lipgloss.Color(red)).Render(m.list.View()), help)
	}

	if m.copy {
		m.copy = false
		return lipgloss.JoinVertical(position, main.BorderForeground(lipgloss.Color(green)).Render(m.list.View()), help)
	}

	return lipgloss.JoinVertical(position, main.BorderForeground(lipgloss.Color(white)).Render(m.list.View()), help)
}

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
