package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// AuthScreen данные для авторизации
type AuthScreen struct {
	// Для ввода данных с группы textInput
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
}

// Update loop for the second view after a choice has been made
func updateAuthScreen(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Смена режима курсора
		case "ctrl+r":
			m.AuthScreen.CursorMode++
			if m.AuthScreen.CursorMode > cursor.CursorHide {
				m.AuthScreen.CursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.AuthScreen.Inputs))
			for i := range m.AuthScreen.Inputs {
				cmds[i] = m.AuthScreen.Inputs[i].Cursor.SetMode(m.AuthScreen.CursorMode)
			}
			return m, tea.Batch(cmds...)

		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.AuthScreen.FocusIndex == len(m.AuthScreen.Inputs) {
				//cmdNew := cmdWithArg(m.AuthScreen.Inputs)
				m.CurrentScreen = "logo"
				m.OTPMessageSended = true
				return m, nil //tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.AuthScreen.FocusIndex--
			} else {
				m.AuthScreen.FocusIndex++
			}

			if m.AuthScreen.FocusIndex > len(m.AuthScreen.Inputs) {
				m.AuthScreen.FocusIndex = 0
			} else if m.AuthScreen.FocusIndex < 0 {
				m.AuthScreen.FocusIndex = len(m.AuthScreen.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.AuthScreen.Inputs))
			for i := 0; i <= len(m.AuthScreen.Inputs)-1; i++ {
				if i == m.AuthScreen.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.AuthScreen.Inputs[i].Focus()
					m.AuthScreen.Inputs[i].PromptStyle = focusedStyle
					m.AuthScreen.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.AuthScreen.Inputs[i].Blur()
				m.AuthScreen.Inputs[i].PromptStyle = noStyle
				m.AuthScreen.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateRegInputs(msg)
	return *m, cmd
}

// initAuthScreen инициализирует по умолчанию
func initAuthScreen() AuthScreen {
	m := AuthScreen{
		Inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Login"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.Inputs[i] = t
	}

	return m
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
