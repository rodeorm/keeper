package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type MainScreen struct {
	Choice int  // Текущий выбор
	Chosen bool // Cделан выбор или нет
}

func initMainScreen() MainScreen {
	return MainScreen{}
}

// updateMain loop для экрана Main
func updateMain(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.MainScreen.Choice++
			if m.MainScreen.Choice > 3 {
				m.MainScreen.Choice = 3
			}
		case "k", "up":
			m.MainScreen.Choice--
			if m.MainScreen.Choice < 0 {
				m.MainScreen.Choice = 0
			}
		case "enter":
			m.MainScreen.Chosen = true
			switch m.MainScreen.Choice {
			case 0:
				m.CurrentScreen = "card"
				return m, textinput.Blink
			case 1:
				m.CurrentScreen = "couple"
				return m, textinput.Blink
			case 2:
				m.CurrentScreen = "text"
				return m, textinput.Blink
			case 3:
				m.CurrentScreen = "bin"
				return m, textinput.Blink
			}
		}
	}

	return m, nil
}

// mainView - первое представление, где начинается работа в программе
func mainView(m *Model) string {
	c := m.MainScreen.Choice //  Забираем значение из модели, что выбрано
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
