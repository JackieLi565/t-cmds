package main

import (
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

var MainModel = NewModel()

func main() {
	p := tea.NewProgram(MainModel, tea.WithAltScreen())
	
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
	
	if c, ok := MainModel.selected.(CmdItem); ok {
		// Command is prefixed with bash as exec does not invoke a shell
		// rather it calls the command like a binary
		args := append([]string{"-c", c.Cmd}, c.Args...)
		cmd := exec.Command("bash", args...)

		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}
	
		fmt.Println(string(output))
	}
}