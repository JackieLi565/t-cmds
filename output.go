package main

import (
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type result []byte

func executeCmd(c CmdItem) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second * 3)
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

	return Output{
		cmd:     item,
		spinner: s,
	}
}

type Output struct {
	cmd     CmdItem
	spinner spinner.Model
	err     error
	output  []byte
	done    bool
}

func (o Output) Init() tea.Cmd {
	return nil
}

func (o Output) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return o, tea.Quit
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
	if o.err != nil {
		errMsg := notification(lipgloss.Color(red)).Render("Failed to execute command!")
		copyMsg := notification(lipgloss.Color(green)).Render("Command copied to clipboard!")

		return lipgloss.JoinVertical(lipgloss.Center, errMsg, copyMsg)
	}

	if o.done {
		resTitle := title.MarginBottom(1).Render("Output Result")
		return border(lipgloss.Color(white)).Render(lipgloss.JoinVertical(lipgloss.Left, resTitle, string(o.output)))
	}

	msg := lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Render("Executing command")
	return border(lipgloss.Color(white)).Render(lipgloss.JoinHorizontal(lipgloss.Center, o.spinner.View(), msg))
}
