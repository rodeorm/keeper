package cli

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

// BinaryPick экран выбора файла
type BinaryPick struct {
	filepicker   filepicker.Model
	selectedFile string

	err error
}

func initBinaryPick() BinaryPick {
	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.UserHomeDir()
	return BinaryPick{filepicker: fp}
}

func updateBinaryPick(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.BinaryPick.filepicker, cmd = m.BinaryPick.filepicker.Update(msg)

	// Пользователь выбрал файл?
	if didSelect, path := m.BinaryPick.filepicker.DidSelectFile(msg); didSelect {
		// Получаем путь выбранного файла
		m.selectedFile = path
		m.CurrentScreen = "binaryAdd" // Переходим на экран добавления дополнительных параметров к выбранному файлу
	}

	return m, cmd
}

func viewBinaryPick(m *Model) string {
	var s strings.Builder
	s.WriteString("\n  ")
	if m.selectedFile == "" {
		s.WriteString("Выберите файл для сохранения:")
	} else {
		s.WriteString("Выбран файл: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	return s.String()
}
