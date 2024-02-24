package main

import "github.com/charmbracelet/bubbles/key"

func (k Key) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k Key) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, 
		{k.NewTask, k.Delete},
		{k.Help, k.Quit},                
	}
}

type Key struct {
	NewTask    	key.Binding
	Delete 			key.Binding
	Up     			key.Binding
	Down   			key.Binding
	Right  			key.Binding
	Left   			key.Binding
	Enter  			key.Binding
	Help   			key.Binding
	Quit   			key.Binding
}

var Keys = Key{
	NewTask: key.NewBinding(
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