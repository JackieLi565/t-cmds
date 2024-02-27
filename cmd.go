package main

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/list"
)

const CMD_FILE = "commands.json"

type CmdJSON struct {
	Name 		string `json:"name"`
	Cmd 		string `json:"cmd"`
	Args 		[]string `json:"args"`
}

type CmdItem struct {
	Name    string
	Cmd   	string
	Args 		[]string
}

func NewCmd(name string, cmd string, args []string) CmdItem {
	return CmdItem{
		Name: name,
		Cmd: cmd,
		Args: args,
	}
}

func getCommands() []CmdItem {
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
		cmds = append(cmds, NewCmd(cmd.Name, cmd.Cmd, cmd.Args))
	}

	return cmds
}

func (c CmdItem) Title() string {
	return c.Name
}

func (c CmdItem) Description() string {
	return cmdWithArgs(c.Cmd, c.Args)
}

func (c CmdItem) FilterValue() string {
	return c.Name
}

func cmdWithArgs(cmd string, args []string) string {
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

func saveToDisk(items []list.Item) error {
	jsonItems := make([]CmdJSON, 0, len(items))

	for _, item := range items {
		if c, ok := item.(CmdItem); ok {
			jsonItems = append(jsonItems, CmdJSON{
				Name: c.Name,
				Cmd: c.Cmd,
				Args: c.Args,
			})
		} else {
			return errors.New("foreign type within cmd list")
		}
	}

	jsonStr, err := json.Marshal(jsonItems)
	if err != nil {
		return err
	}

	if err := os.WriteFile(CMD_FILE, jsonStr, 6044); err != nil {
		return err
	}

	return nil
}

func parseArgs(args string) []string {
	var parsedArgs []string
	var arg strings.Builder
	var insideQuote bool

	for _, char := range args {
		if unicode.Is(unicode.Quotation_Mark, char) {
			insideQuote = !insideQuote
			arg.WriteRune(char)
			continue
		}

		if insideQuote {
			arg.WriteRune(char)
		} else {
			if unicode.Is(unicode.Space, char) && arg.Len() > 0 {
				parsedArgs = append(parsedArgs, arg.String())
				arg.Reset()
			} else {
				arg.WriteRune(char)
			}
		}
	}

	if arg.Len() > 0 {
		parsedArgs = append(parsedArgs, arg.String())
	}

	return parsedArgs
} 