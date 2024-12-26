package cli

import (
	"context"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/client"
)

// CoupleList
type CoupleList struct {
	table table.Model
}

type coupleListMsg struct {
	couples []core.Couple
	err     error
}

func (m *Model) listCouple() tea.Msg {
	msg := coupleListMsg{}
	ctx := context.TODO()
	cds, err := client.ReadAllCouples(ctx, m.Token, m.sc)
	if err != nil {
		msg.err = err
		return msg
	}

	msg.couples = make([]core.Couple, 0)

	for _, v := range cds {
		msg.couples = append(msg.couples,
			core.Couple{
				Source:   v.Source,
				Login:    v.Login,
				Password: v.Password,
				Meta:     v.Meta,
			})
	}

	return msg
}

func initCoupleList() CoupleList {
	columns := []table.Column{
		{Title: "Источник", Width: 50},
		{Title: "Логин", Width: 50},
		{Title: "Пароль", Width: 50},
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

	return CoupleList{table: t}
}

func updateCoupleList(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			return m, m.listCouple
		case "c":
			m.CoupleCreate = initCoupleCreate()
			m.CurrentScreen = "coupleCreate"
			return m, nil
		case "ctrl+z":
			if m.CoupleList.table.Focused() {
				m.CoupleList.table.Blur()
			} else {
				m.CoupleList.table.Focus()
			}
		case "enter":
			return m, tea.Batch(
				tea.Printf("Выбрана пара логин-пароль%s/%s!", m.CoupleList.table.SelectedRow()[1], m.CoupleList.table.SelectedRow()[2]),
			)
		}
	case coupleListMsg:
		rows := make([]table.Row, 0)
		for _, v := range msg.couples {
			r := table.Row{v.Source, v.Login, v.Password, v.Meta}
			rows = append(rows, r)
		}

		m.CoupleList.table.SetRows(rows)
		m.CoupleList.table.SetHeight(len(rows) + 2)
		m.CoupleList.table.Focus()
		m.CoupleList.table.GotoTop()
		m.CoupleList.table.UpdateViewport()
	}
	m.CoupleList.table, cmd = m.CoupleList.table.Update(msg)
	return m, cmd
}

func viewCoupleList(m *Model) string {
	return m.header() + baseStyle.Render(m.CoupleList.table.View()) + "\n" + footerTable() + "\n" + footer()
}
