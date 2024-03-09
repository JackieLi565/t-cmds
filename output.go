package main

import (
	"os/exec"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type result []byte

func executeCmd(c CmdItem) tea.Cmd {
	return func() tea.Msg {
		args := append([]string{"-c", c.Cmd}, c.Args...)
		cmd := exec.Command("bash", args...)

		output, err := cmd.Output()
		if err != nil {
			return err
		}
		return result(output)
	}
}

func NewOutput(item CmdItem) Output {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(blue)).MarginRight(1)

	help := help.New()
	help.ShowAll = false

	return Output{
		cmd:     item,
		spinner: s,
		help: 	 help,
	}
}

type Output struct {
	cmd     CmdItem
	spinner spinner.Model
	err     error
	output  []byte
	done    bool
	help		help.Model
}

func (o Output) Init() tea.Cmd {
	return nil
}

func (o Output) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if o.done {
			return MainModel.Update(nil)
		} else if o.err != nil {
			return o, tea.Quit
		}
	case result:
		o.done = true
		o.output = msg
	case spinner.TickMsg:
		var cmd tea.Cmd
		o.spinner, cmd = o.spinner.Update(msg)
		return o, cmd
	case error:
		o.err = msg
	case initOutput:
		return o, tea.Batch(
			o.spinner.Tick,
			executeCmd(o.cmd),
		)
	}

	return o, nil
}

func (o Output) View() string {
	b := border.BorderForeground(lipgloss.Color(white)).Width(60)
	textWhite := lipgloss.NewStyle().Foreground(lipgloss.Color(white))
	textGreen := lipgloss.NewStyle().Foreground(lipgloss.Color(green)).MarginTop(1)

	if o.err != nil {
		errMsg := textWhite.Render("Failed to execute command!")
		copyMsg := textGreen.Render("Command copied to clipboard!")

		return b.Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Center, errMsg, copyMsg))
	}

	if o.done {
		msg := textWhite.Render("Press any key to return")
		resTitle := title.MarginBottom(1).Render("Output Result")
		return b.Render(lipgloss.JoinVertical(lipgloss.Left, resTitle, string(o.output), msg))
	}

	msg := textWhite.Render("Executing command, press q to quit")
	outView := lipgloss.JoinHorizontal(lipgloss.Center, o.spinner.View(), msg)
	return b.Render(outView)
}
