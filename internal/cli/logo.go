package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LogoScreen struct {
	Choice int  // Текущий выбор
	Chosen bool // Cделан выбор или нет
}

func initLogoScreen() LogoScreen {
	return LogoScreen{}
}

// updateLogo loop для первого представления
func updateLogo(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.LogoScreen.Choice++
			if m.LogoScreen.Choice > 3 {
				m.LogoScreen.Choice = 3
			}
		case "k", "up":
			m.LogoScreen.Choice--
			if m.LogoScreen.Choice < 0 {
				m.LogoScreen.Choice = 0
			}
		case "enter":
			m.LogoScreen.Chosen = true
			switch m.LogoScreen.Choice {
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

// logoView - первое представление, где начинается работа в программе
func logoView(m *Model) string {
	c := m.LogoScreen.Choice //  Забираем значение из модели, что выбрано
	var choices, tpl string

	tpl = "Добро пожаловать. Что вы хотите сделать?\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("esc: quit")

	choices = fmt.Sprintf(
		"%s\n%s\n",
		checkbox("Зарегистрироваться", c == 0), // Проставляем галочку на выбранном значении
		checkbox("Авторизоваться", c == 1),     // Проставляем галочку на выбранном значении
	)
	return fmt.Sprintf(tpl, choices)
}
