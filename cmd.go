package main

import (
	"encoding/json"
	"os"
	"strings"
)

const CMD_FILE = "commands.json"

type CmdJSON struct {
	Name string `json:"name"`
	Cmd string `json:"cmd"`
	Args []string `json:"args"`
}

type CmdItem struct {
	Name      string
	Cmd   string
	Args []string
}

func NewCmd(name string, cmd string, args []string) CmdItem {
	return CmdItem{
		Name: name,
		Cmd: cmd,
		Args: args,
	}
}

func GetCommands() []CmdItem {
	jsonCmds := make([]CmdJSON, 0)

	dat, err := os.ReadFile(CMD_FILE)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(dat, &jsonCmds); err != nil {
		panic(err)
	}

	cmds := make([]CmdItem, 0, len(jsonCmds))
	for _, cmd := range jsonCmds {
		cmds = append(cmds, NewCmd(cmd.Name, cmd.Name, cmd.Args))
	}

	return cmds
}

func (c CmdItem) Title() string {
	return c.Name
}

func (c CmdItem) Description() string {
	return CmdWithArgs(c.Cmd, c.Args)
}

func (c CmdItem) FilterValue() string {
	return c.Name
}

func CmdWithArgs(cmd string, args []string) string {
	var str strings.Builder
	
	str.WriteString(cmd)
	str.WriteString(" ")

	for i, arg := range args {
		str.WriteString(arg)

		if i != len(args) - 1 {
			str.WriteString(" ")
		} 
	}

	return str.String()
}

// TODO parse a string of arguments into a slice
func ParseArgs(args string) []string {
	return make([]string, 0)
} 