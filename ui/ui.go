package ui

import (
	"strings"

	"github.com/benmatselby/prolificli/model"
	"github.com/charmbracelet/lipgloss"
)

const (
	// DarkBlue is a colour used in the UI.
	DarkBlue = "#083759"
	// LightBlue is a colour used in the UI.
	LightBlue = "#e3f3ff"
	// Green is a colour used in the UI.
	Green = "#008033"
)

// AppDateTimeFormat The format for date/times in the application.
const AppDateTimeFormat string = "02-01-2006 15:04"

// RenderTitle will render a nice coloured UI for a title based on status
func RenderTitle(title, status string) lipgloss.Style {

	var color = ""
	switch strings.ToLower(status) {
	case model.StatusActive:
	case model.StatusCompleted:
		color = Green
	case model.StatusAwaitingReview:
	case model.StatusScheduled:
	case model.StatusUnpublished:
		color = DarkBlue
	}

	var style = lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Background(lipgloss.Color(color)).
		MarginBottom(1).
		Padding(1).
		Align(lipgloss.Center).
		SetString(title)

	return style
}

// RenderHeading will render a heading in the output.
func RenderHeading(heading string) string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Background(lipgloss.Color(LightBlue)).
		MarginBottom(1).
		Align(lipgloss.Center)

	return style.Render(heading)
}
