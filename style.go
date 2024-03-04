package main

import (
	"github.com/charmbracelet/lipgloss"
)

type color string

const (
	red   color = "#FF6961"
	green color = "#77DD77"
	white color = "#FFFFFF"
	blue  color = "#57A0D2"
)

func notification(color lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().
		MarginTop(1).
		PaddingLeft(1).
		PaddingRight(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(color).
		Foreground(color)
}

func border(color lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(color)).
		PaddingTop(1).
		PaddingLeft(2).
		PaddingRight(2).
		PaddingBottom(1)
}

var title = lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color(blue)).Padding(0, 1, 0)
