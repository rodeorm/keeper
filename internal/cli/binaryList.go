package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BinaryListScreen данные для отображения карт пользователя
type BinaryList struct {
	table table.Model
}

func initBinaryList() BinaryList {
	columns := []table.Column{
		{Title: "№", Width: 3},
		{Title: "Мета", Width: 50},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithHeight(0),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("225")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(s)

	return BinaryList{table: t}
}

func updateBinaryList(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			return m, m.listBinary
		case "c":
			m.BinaryCreate = initBinaryCreate()
			m.CurrentScreen = "binaryCreate"
			return m, tea.Batch(m.filepicker.Init(), textinput.Blink)
		case "ctrl+z":
			if m.BinaryList.table.Focused() {
				m.BinaryList.table.Blur()
			} else {
				m.BinaryList.table.Focus()
			}
		case "enter":
			return m, tea.Batch(
				tea.Printf("Выбран файл номер %s!", m.BinaryList.table.SelectedRow()[1]),
			)
		}
	case binaryListMsg:
		rows := make([]table.Row, 0)
		for i, v := range msg.bins {
			r := table.Row{fmt.Sprint(i), v.Meta}
			rows = append(rows, r)
		}
		m.BinaryList.table.SetRows(rows)
		m.BinaryList.table.SetHeight(len(rows) + 2)
		m.BinaryList.table.Focus()
		m.BinaryList.table.GotoTop()
		m.BinaryList.table.UpdateViewport()
	}
	m.BinaryList.table, cmd = m.BinaryList.table.Update(msg)
	return m, cmd
}

func viewBinaryList(m *Model) string {
	return m.header() + baseStyle.Render(m.BinaryList.table.View()) + "\n" + footerTable() + "\n" + footer()
}
