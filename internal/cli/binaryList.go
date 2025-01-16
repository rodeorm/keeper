package cli

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/client"
)

// BinaryListScreen данные для отображения карт пользователя
type BinaryList struct {
	table table.Model
	Msg   string
}

type binaryListMsg struct {
	bins []core.Binary
	err  error
}

func (m *Model) listBinary() tea.Msg {
	ctx := context.TODO()

	bs, err := client.ReadAllBinaries(ctx, m.Token, m.sc)

	msg := binaryListMsg{}
	if err != nil {
		msg.err = err
		return msg
	}

	msg.bins = make([]core.Binary, 0)

	for _, v := range bs {
		msg.bins = append(msg.bins,
			core.Binary{
				Value: v.Value,
				Name:  v.Name,
				Meta:  v.Meta,
				ID:    int(v.Id),
			})
	}

	return msg
}

func initBinaryList() BinaryList {
	columns := []table.Column{
		{Title: "Имя", Width: 50},
		{Title: "Мета", Width: 100},
		{Title: "ID", Width: 0}, // Бинарник не показываем, он нужен только для быстрого возврата. Большие файлы так лучше не хранить, конечно
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
			id, _ := strconv.Atoi(m.BinaryList.table.SelectedRow()[2])
			bin := core.Binary{
				Name: m.BinaryList.table.SelectedRow()[0],
				ID:   id,
			}
			if m.FilePath == "" {
				m.FilePath, _ = os.UserHomeDir()
			}
			svCmd := m.saveBinary(bin, m.FilePath)
			m.BinaryList.Msg = fmt.Sprintf("Файл %s (id %d) будет сохранен в папку %s", bin.Name, bin.ID, m.FilePath)
			return m, tea.Batch(svCmd)
		}
	case saveBinaryMsg:
		if msg.err != nil {
			m.BinaryList.Msg = fmt.Sprintf("Ошибка при сохранении %s", msg.err.Error())
			return m, nil
		} else {
			m.BinaryList.Msg = fmt.Sprintf("Файл удачно сохранили в папку %s", m.FilePath)
			return m, nil
		}
	case binaryListMsg:
		rows := make([]table.Row, 0)
		for _, v := range msg.bins {
			r := table.Row{v.Name, v.Meta, strconv.Itoa(v.ID)}
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
	if m.Msg != "" {
		return m.header() + m.Msg + "\n" + baseStyle.Render(m.BinaryList.table.View()) + "\n" + footerTable() + "\n" + footer()
	}
	return m.header() + baseStyle.Render(m.BinaryList.table.View()) + "\n" + footerTable() + "\n" + footer()
}
