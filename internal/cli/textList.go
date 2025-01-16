package cli

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/client"
)

// TextListScreen
type TextList struct {
	table table.Model
}

type TextListMsg struct {
	Texts []core.Text
	err   error
}

func (m *Model) listText() tea.Msg {
	msg := TextListMsg{}
	ctx := context.TODO()
	cds, err := client.ReadAllTexts(ctx, m.Token, m.sc)
	if err != nil {
		msg.err = err
		return msg
	}

	msg.Texts = make([]core.Text, 0)

	for _, v := range cds {
		msg.Texts = append(msg.Texts,
			core.Text{
				Value: v.Text,
				Meta:  v.Meta,
			})
	}

	return msg
}

func initTextList() TextList {
	columns := []table.Column{
		{Title: "Номер", Width: 3},
		{Title: "Текст", Width: 100},
		{Title: "Мета", Width: 100},
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

	return TextList{table: t}
}

func updateTextList(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			return m, m.listText
		case "c":
			m.TextCreate = initTextCreate()
			m.CurrentScreen = "textCreate"
			return m, nil
		case "ctrl+z":
			if m.TextList.table.Focused() {
				m.TextList.table.Blur()
			} else {
				m.TextList.table.Focus()
			}
		case "enter":
			return m, tea.Batch(
				tea.Printf("Выбран текст %s!", m.TextList.table.SelectedRow()[1]),
			)
		}
	case TextListMsg:
		rows := make([]table.Row, 0)
		for i, v := range msg.Texts {
			r := table.Row{fmt.Sprint(i), v.Value, v.Meta}
			rows = append(rows, r)
		}

		m.TextList.table.SetRows(rows)
		m.TextList.table.SetHeight(len(rows) + 2)
		m.TextList.table.Focus()
		m.TextList.table.GotoTop()
		m.TextList.table.UpdateViewport()
	}
	m.TextList.table, cmd = m.TextList.table.Update(msg)
	return m, cmd
}

func viewTextList(m *Model) string {
	return m.header() + baseStyle.Render(m.TextList.table.View()) + "\n" + footerTable() + "\n" + footer()
}
