package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
	focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle   = focusedStyle
	noStyle       = lipgloss.NewStyle()
	//helpStyle           = blurredStyle
	//cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	docStyle      = lipgloss.NewStyle().Margin(1, 2)
	focusedButton = focusedStyle.Render("[ Отправить ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Отправить"))
	baseStyle     = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
)

const (
	dotChar = " • "
)

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
