package main

import (
	"github.com/charmbracelet/lipgloss"
)

var notificationStyle = lipgloss.NewStyle().
				MarginTop(1).
        PaddingLeft(1).
        PaddingRight(1).
        Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#FFFFFF"))

var borderStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#FFFFFF")).
        PaddingTop(1).
        PaddingLeft(2).
        PaddingRight(2).
        PaddingBottom(1)