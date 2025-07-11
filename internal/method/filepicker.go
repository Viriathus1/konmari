package method

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	fp           filepicker.Model
	selectedFile string
	quitting     bool
	err          error
}

func NewFilePicker() tea.Model {
	fp := filepicker.New()
	fp.AllowedTypes = []string{"*"} // allow all
	fp.CurrentDirectory, _ = os.Getwd()

	return model{fp: fp}
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return m.fp.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.fp, cmd = m.fp.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.fp.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.selectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.fp.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.fp.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Pick a file:")
	} else {
		s.WriteString("Selected file: " + m.fp.Styles.Selected.Render(m.selectedFile))
	}
	s.WriteString("\n\n" + m.fp.View() + "\n")
	return s.String()
}
