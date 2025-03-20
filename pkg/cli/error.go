package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/samber/oops"
)

func PrintFatalError(err error) {
	if oopsErr, ok := oops.AsOops(err); ok {
		sources := SourcesStyle().Render(oopsErr.Sources())
		stacktrace := StackTraceStyle().Render(oopsErr.Stacktrace())
		message := ErrorStyle().Render(oopsErr.Error())
		fmt.Printf("%s\n%s\n%s\n", sources, stacktrace, message)
	} else {
		message := ErrorStyle().Render(err.Error())
		fmt.Printf("%s\n", message)
	}
}

func SourcesStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")). // bright red
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(0, 1)
}

func StackTraceStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")). // bright red
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(0, 1)
}

func ErrorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")). // bright red
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(0, 1)
}
