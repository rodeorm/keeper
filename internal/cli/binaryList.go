package cli

import (
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rodeorm/keeper/internal/core"
)

// BinaryListScreen данные для отображения карт пользователя
type BinaryList struct {
	table table.Model
}

func initBinaryList() BinaryList {
	columns := []table.Column{
		{Title: "Имя", Width: 50},
		{Title: "Мета", Width: 1000},
		{Title: "Сам бинарник", Width: 0}, // Бинарник не показываем, он нужен только для быстрого возврата. Большие файлы так лучше не хранить, конечно
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
			m.BinaryPick = initBinaryPick()
			m.CurrentScreen = "binaryPick"
			return m, m.BinaryPick.filepicker.Init()
		case "ctrl+z":
			if m.BinaryList.table.Focused() {
				m.BinaryList.table.Blur()
			} else {
				m.BinaryList.table.Focus()
			}
		case "enter":
			bin := core.Binary{
				Name:  m.BinaryList.table.SelectedRow()[0],
				Value: []byte(m.BinaryList.table.SelectedRow()[2]),
			}
			dir, _ := os.UserHomeDir()
			svCmd := saveBinary(bin, dir)

			return m, tea.Batch(
				tea.Printf("Файл %s будет сохранен в папку %s", bin.Name, dir),
				svCmd,
			)
		}
	case saveBinaryMsg:
		if msg.err != nil {
			return m, tea.Printf("Ошибка при сохранении %s", msg.err.Error())
		} else {
			return m, tea.Printf("Удачно сохранили")
		}
	case binaryListMsg:
		rows := make([]table.Row, 0)
		for _, v := range msg.bins {
			r := table.Row{v.Name, v.Meta, string(v.Value)}
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
