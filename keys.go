package main

import "github.com/charmbracelet/bubbles/key"

type ModelKey struct {
	NewCmd    	key.Binding
	CopyCmd			key.Binding
	Delete 			key.Binding
	Enter  			key.Binding
	Help   			key.Binding
	Quit   			key.Binding
}

func (k ModelKey) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k ModelKey) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NewCmd, k.CopyCmd, k.Delete},
		{k.Help, k.Quit},                
	}
}

type FormKey struct {
	Enter    		key.Binding
	Cancel   	  key.Binding
}

func (k FormKey) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Cancel}
}

func (k FormKey) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter, k.Cancel}, 
	}
}

var FormKeys = FormKey{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "continue"),
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
	CopyCmd: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "cpy cmd"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "run cmd"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "del cmd"),
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