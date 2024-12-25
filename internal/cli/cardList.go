package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// CardListScreen данные для отображения карт пользователя
type CardList struct {
	table table.Model
}

func initCardList() CardList {
	columns := []table.Column{
		{Title: "№", Width: 3},
		{Title: "Номер", Width: 20},
		{Title: "Дата", Width: 4},
		{Title: "Год", Width: 3},
		{Title: "Имя владельца", Width: 20},
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

	return CardList{table: t}
}

func updateCardList(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			return m, m.listCard
		case "c":
			m.CardCreate = initCardCreate()
			m.CurrentScreen = "cardCreate"
			return m, nil
		case "ctrl+z":
			if m.CardList.table.Focused() {
				m.CardList.table.Blur()
			} else {
				m.CardList.table.Focus()
			}
		case "enter":
			return m, tea.Batch(
				tea.Printf("Выбрана карта номер %s!", m.CardList.table.SelectedRow()[1]),
			)
		}
	case cardListMsg:
		rows := make([]table.Row, 0)
		for i, v := range msg.cards {
			r := table.Row{fmt.Sprint(i), v.CardNumber, fmt.Sprint(v.ExpYear), fmt.Sprint(v.ExpMonth), v.OwnerName, v.Meta}
			rows = append(rows, r)
		}

		m.CardList.table.SetRows(rows)
		m.CardList.table.SetHeight(len(rows) + 2)
		m.CardList.table.Focus()
		m.CardList.table.GotoTop()
		m.CardList.table.UpdateViewport()
	}
	m.CardList.table, cmd = m.CardList.table.Update(msg)
	return m, cmd
}

func viewCardList(m *Model) string {
	return m.header() + baseStyle.Render(m.CardList.table.View()) + "\n" + footerTable() + "\n" + footer()
}
