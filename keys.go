package main

import "github.com/charmbracelet/bubbles/key"

type ModelKey struct {
	NewCmd    	key.Binding
	Delete 			key.Binding
	Up     			key.Binding
	Down   			key.Binding
	Right  			key.Binding
	Left   			key.Binding
	Enter  			key.Binding
	Help   			key.Binding
	Quit   			key.Binding
}

func (k ModelKey) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k ModelKey) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, 
		{k.NewCmd, k.Delete},
		{k.Help, k.Quit},                
	}
}

type FormKey struct {
	Enter    		key.Binding
	Up     			key.Binding
	Down   			key.Binding
	Help   			key.Binding
	Cancel   	  key.Binding
}

func (k FormKey) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Cancel}
}

func (k FormKey) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter}, 
		{k.Help, k.Cancel},                
	}
}

var FormKeys = FormKey {
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithKeys("enter", "continue"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "cancel"),
	),
}

var ModelKeys = ModelKey{
	NewCmd: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new cmd"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "right"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/l", "left"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "enter"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}