package method

import "github.com/charmbracelet/lipgloss"

var (
	dimmedStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#999999", // light gray
		Dark:  "#666666", // darker gray
	}).Faint(true)
	normalStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#000000",
		Dark:  "#FFFFFF",
	})
	descriptionStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#444444",
		Dark:  "#BBBBBB",
	})
	colourStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#005FFF", // Strong blue for light backgrounds
		Dark:  "#5FAFFF", // Lighter blue for dark backgrounds
	}).Bold(true)
)
