package cli

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

// BinaryCreate данные карты
type BinaryCreate struct {
	filepicker   filepicker.Model
	selectedFile string

	err error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func updateBinaryCreate(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case clearErrorMsg:
		m.BinaryCreate.err = nil
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		m.selectedFile = path
	}
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		m.BinaryCreate.err = errors.New(path + " - неправильный")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func viewBinaryCreate(m *Model) string {
	var s strings.Builder
	s.WriteString("\n  ")
	if m.BinaryCreate.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.BinaryCreate.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Выберите файл:")
	} else {
		s.WriteString("Выбран файл: " + m.filepicker.Styles.Selected.Render(m.selectedFile))

		return m.header() + s.String() + footer()
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	return m.header() + s.String() + footer()
}

func initBinaryCreate() BinaryCreate {
	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.UserHomeDir()

	m := BinaryCreate{
		filepicker: fp,
	}

	return m
}
