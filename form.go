package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))
	continueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#767676"))
)

const (
	name int = iota
	command
	arguments
)

type Form struct {
	inputs  []textinput.Model
	help    help.Model
	focused int
}

func NewForm() Form {
	inputs := make([]textinput.Model, 3)

	nameInput := textinput.New()
	nameInput.Placeholder = "Make A NextJS Project"
	nameInput.CharLimit = 300 // if you have a command longer then this please show me
	nameInput.Width = 30
	nameInput.Focus()

	cmd := textinput.New()
	cmd.Placeholder = "ls or pwd"
	cmd.CharLimit = 50
	cmd.Width = 20

	args := textinput.New()
	args.Placeholder = "-l | wc -l"
	args.CharLimit = 250
	args.Width = 40

	inputs[name] = nameInput
	inputs[command] = cmd
	inputs[arguments] = args

	help := help.New()
	help.ShowAll = false

	return Form{
		inputs:  inputs,
		focused: 0,
		help:    help,
	}
}

func (f Form) Init() tea.Cmd {
	return nil
}

func (f Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(f.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, FormKeys.Enter):
			return MainModel.Update(f)
		case key.Matches(msg, FormKeys.Cancel):
			return MainModel.Update(nil)
		case key.Matches(msg, FormKeys.Up):
			f.prevInput()
		case key.Matches(msg, FormKeys.Down):
			f.nextInput()
		case key.Matches(msg, FormKeys.Help):
			f.help.ShowAll = !f.help.ShowAll
		}

		for i := range f.inputs {
			f.inputs[i].Blur()
		}

		f.inputs[f.focused].Focus()
	}

	for i := range f.inputs {
		f.inputs[i], cmds[i] = f.inputs[i].Update(msg)
	}

	return f, tea.Batch(cmds...)
}

func (f Form) View() string {
	formTitle := title.Render("New Command")
	form := fmt.Sprintf(
		`%s
		
 %s
 %s

 %s %s

 %s %s

 %s
		`, formTitle, inputStyle.Width(30).Render("Command Name"),
		f.inputs[name].View(),
		inputStyle.Width(10).Render("Command"),
		f.inputs[command].View(),
		inputStyle.Width(10).Render("Arguments"),
		f.inputs[arguments].View(),
		continueStyle.Render("Press Enter to Submit"),
	)

	return border(lipgloss.Color(white)).Render(lipgloss.JoinVertical(lipgloss.Left, form, f.help.View(FormKeys)))
}

func (m *Form) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *Form) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
