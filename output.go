package main

import (
	"os/exec"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type result []byte

func executeCmd(c CmdItem) tea.Cmd {
	return func () tea.Msg {
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
	return Output{
		cmd: item,
		spinner: spinner.New(),
	}
}

type Output struct {
	cmd CmdItem
	spinner spinner.Model
	output []byte
	done bool
}

func (o Output) Init() tea.Cmd {
	return tea.Batch(
		o.spinner.Tick,
		executeCmd(o.cmd),
	)
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
	}

	return o, o.Init()
}

func (o Output) View() string {
	if o.done {
		return borderStyle.Render(string(o.output))
	}

	return o.spinner.View()
}