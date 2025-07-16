package method

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	dimmedStyle      = lipgloss.NewStyle().Faint(true)
	normalStyle      = lipgloss.NewStyle()
	descriptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	cursorStyle      = lipgloss.NewStyle().Bold(true)
)

type toggleDelegate struct {
	selected map[string]bool
}

func (d toggleDelegate) Height() int                               { return 2 }
func (d toggleDelegate) Spacing() int                              { return 0 }
func (d toggleDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d toggleDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	if m.Width() <= 0 {
		// short-circuit
		return
	}

	it, ok := item.(fileItem)
	if !ok {
		return
	}

	cursor := "  "
	if index == m.Index() {
		cursor = cursorStyle.Render("> ")
	}

	// Apply dimming if unselected
	var line string
	if d.selected[it.path] {
		line = normalStyle.Render(it.path)
	} else {
		line = dimmedStyle.Render(it.path)
	}

	desc := descriptionStyle.Render(it.Description())

	fmt.Fprintf(w, "%s%s\n  %s\n", cursor, line, desc)
}

func (d toggleDelegate) ShortHelp() []key.Binding {
	return []key.Binding{
		toggleKey,
		selectKey,
	}
}

func (d toggleDelegate) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{toggleKey, selectKey},
	}
}
