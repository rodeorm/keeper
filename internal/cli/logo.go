package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Logo struct {
	Choice int  // Текущий выбор
	Chosen bool // Cделан выбор или нет
}

func initLogo() Logo {
	return Logo{}
}

// updateLogo loop для первого представления
func updateLogo(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Logo.Choice++
			if m.Logo.Choice > 3 {
				m.Logo.Choice = 3
			}
		case "k", "up":
			m.Logo.Choice--
			if m.Logo.Choice < 0 {
				m.Logo.Choice = 0
			}
		case "enter":
			m.Logo.Chosen = true
			switch m.Logo.Choice {
			case 0:
				m.CurrentScreen = "reg"
				return m, textinput.Blink
			case 1:
				m.CurrentScreen = "auth"
				return m, textinput.Blink
			}
		}
	}

	return m, nil
}

func viewLogo(m *Model) string {
	c := m.Logo.Choice //  Забираем значение из модели, что выбрано
	var choices, tpl string

	tpl = "Добро пожаловать. Что вы хотите сделать?\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("up/down: перемещение по меню") + dotStyle +
		subtleStyle.Render("enter: выбрать") + dotStyle

	choices = fmt.Sprintf(
		"%s\n%s\n",
		checkbox("Зарегистрироваться", c == 0), // Проставляем галочку на выбранном значении
		checkbox("Авторизоваться", c == 1),     // Проставляем галочку на выбранном значении
	)

	return m.header() + fmt.Sprintf(tpl, choices) + footer()
}
