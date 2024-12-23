package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Main struct {
	Choice int  // Текущий выбор
	Chosen bool // Cделан выбор или нет
}

func initMain() Main {
	return Main{}
}

// updateMain loop для экрана Main
func updateMain(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Main.Choice++
			if m.Main.Choice > 3 {
				m.Main.Choice = 3
			}
		case "k", "up":
			m.Main.Choice--
			if m.Main.Choice < 0 {
				m.Main.Choice = 0
			}
		case "enter":
			m.Main.Chosen = true
			switch m.Main.Choice {
			case 0:
				m.CurrentScreen = "cardList"
				return m, tea.Batch(textinput.Blink, m.listCard)
			case 1:
				m.CurrentScreen = "coupleList"
				return m, textinput.Blink
			case 2:
				m.CurrentScreen = "textList"
				return m, textinput.Blink
			case 3:
				m.CurrentScreen = "binList"
				return m, textinput.Blink
			}
		}
	}

	return m, nil
}

func viewMain(m *Model) string {
	c := m.Main.Choice //  Забираем значение из модели, что выбрано
	var choices, tpl string

	tpl = fmt.Sprintf("Добро пожловать, %s!\n\nВыберите раздел данных:\n",
		keywordStyle.Render(m.User.Login))
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("esc: quit")

	choices = fmt.Sprintf("%s\n%s\n%s\n%s\n",
		checkbox("Карты", c == 0),
		checkbox("Пароли", c == 1),
		checkbox("Текст", c == 2),
		checkbox("Бинарники", c == 3))

	return m.header() + fmt.Sprintf(tpl, choices) + footer()
}
