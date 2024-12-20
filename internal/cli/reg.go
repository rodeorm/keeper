package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
)

// RegScreen данные регистрации
type RegScreen struct {
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
}

func (m *Model) updateRegInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.RegScreen.Inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.RegScreen.Inputs {
		m.RegScreen.Inputs[i], cmds[i] = m.RegScreen.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// Update loop for the second view after a choice has been made
func updateRegScreen(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Переместить фокус на следующее поле
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Пользователь нажал на enter, когда выбрана кнопка Submit?
			// Если так, то отправляем сообщение на grpc и возвращаемся на лого форму.
			if s == "enter" && m.RegScreen.FocusIndex == len(m.RegScreen.Inputs) {
				//cmdNew := cmdWithArg(m.RegScreen.Inputs)
				m.CurrentScreen = "logo"
				m.OTPMessageSended = true
				m.CurrentScreen = "wait"
				m.User = core.User{Login: m.RegScreen.Inputs[0].Value(),
					Password: m.RegScreen.Inputs[1].Value(),
					Email:    m.RegScreen.Inputs[2].Value(),
					Name:     m.RegScreen.Inputs[3].Value(),
				}

				return m, m.RegUser
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.RegScreen.FocusIndex--
			} else {
				m.RegScreen.FocusIndex++
			}

			if m.RegScreen.FocusIndex > len(m.RegScreen.Inputs) {
				m.RegScreen.FocusIndex = 0
			} else if m.RegScreen.FocusIndex < 0 {
				m.RegScreen.FocusIndex = len(m.RegScreen.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.RegScreen.Inputs))
			for i := 0; i <= len(m.RegScreen.Inputs)-1; i++ {
				if i == m.RegScreen.FocusIndex {
					// Устанавливаем фокус
					cmds[i] = m.RegScreen.Inputs[i].Focus()
					m.RegScreen.Inputs[i].PromptStyle = focusedStyle
					m.RegScreen.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Убираем фокус
				m.RegScreen.Inputs[i].Blur()
				m.RegScreen.Inputs[i].PromptStyle = noStyle
				m.RegScreen.Inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateRegInputs(msg)
	return *m, cmd
}

// initRegScreen инцицилизирует форму для регистрации по умолчанию
func initRegScreen() RegScreen {
	m := RegScreen{
		Inputs: make([]textinput.Model, 4),
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
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 2:
			t.Placeholder = "Email"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 3:
			t.Placeholder = "Name"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.Inputs[i] = t
	}
	return m
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
	/*
		b.WriteString(helpStyle.Render("cursor mode is "))
		b.WriteString(cursorModeHelpStyle.Render(m.RegScreen.CursorMode.String()))
		b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))
	*/
	msg += "\n" + b.String()
	msg += "\n" + subtleStyle.Render("Для выхода нажмите esc или ctrl + c")
	return msg
}
