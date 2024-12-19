package cli

import (
	"fmt"
	"strings"
)

// View (Представление) — это функция, которая отображает модель в виде, понятном пользователю.
func (m Model) View() string {
	s := ""

	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	switch m.CurrentScreen {
	case "logo":
		s = logoView(&m)
	case "reg":
		s = regView(&m)
	case "auth":
		s = authView(&m)
	case "wait":
		s = WaitForOTPView(&m)
	}
	return mainStyle.Render("\n" + s + "\n\n")
}

// Sub-views
func WaitForOTPView(m *Model) string {
	tpl := "Введите отправленный одноразовый пароль\n\n"
	tpl += m.User.Login
	tpl += ":"
	tpl += m.User.Password
	return tpl
}

// logoView - первое представление, где начинается работа в программе
func logoView(m *Model) string {
	c := m.Choice //  Забираем значение из модели, что выбрано
	var choices, tpl string

	tpl = "Добро пожаловать. Что вы хотите сделать?\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("esc: quit")
	if !m.Authenticated {
		choices = fmt.Sprintf(
			"%s\n%s\n",
			checkbox("Зарегистрироваться", c == 0), // Проставляем галочку на выбранном значении
			checkbox("Авторизоваться", c == 1),     // Проставляем галочку на выбранном значении
		)
	} else {
		choices = fmt.Sprintf(
			"%s\n%s\n%s\n%s\n%s\n",
			checkbox("Логины-пароли", c == 0),
			checkbox("Текст", c == 1),
			checkbox("Кредитные карты", c == 2),
			checkbox("Текст", c == 1),
			checkbox("карты", c == 2),
		)
	}

	return fmt.Sprintf(tpl, choices)
}

// regView - второе представление: либо для регистрации, либо для авторизации
func regView(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Хотите зарегистрироваться?\n\nОтлично! Просто введите  %s, %s и %s.\n",
		keywordStyle.Render("логин"),
		keywordStyle.Render("пароль"),
		keywordStyle.Render("адрес электронной почты"))

	for i := range m.RegScreen.Inputs {
		b.WriteString(m.RegScreen.Inputs[i].View())
		if i < len(m.RegScreen.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.RegScreen.FocusIndex == len(m.RegScreen.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.RegScreen.CursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))
	msg += "\n" + b.String()
	msg += "\n" + subtleStyle.Render("Для выхода нажмите esc или ctrl + c")
	return msg
}

// authView - второе представление: либо для регистрации, либо для авторизации
func authView(m *Model) string {
	var msg string
	var b strings.Builder

	msg = fmt.Sprintf("Уже зарегистрированы? \n\nОтлично! Просто введите %s и %s...\n\n", keywordStyle.Render("логин"), keywordStyle.Render("пароль"))

	msg += subtleStyle.Render("Для выхода нажмите esc")

	for i := range m.AuthScreen.Inputs {
		b.WriteString(m.AuthScreen.Inputs[i].View())
		if i < len(m.AuthScreen.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.RegScreen.FocusIndex == len(m.AuthScreen.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.AuthScreen.CursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))
	msg += "\n" + b.String()
	msg += "\n" + subtleStyle.Render("Для выхода нажмите esc или ctrl + c")

	return msg
}
