package cli

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
)

// Update (Обновление) — это функция, которая принимает текущее состояние (модель) и сообщение (например, событие, вызванное пользователем),
// и возвращает новое состояние.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Всегда перехватывем q esc ctl+c
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	switch m.CurrentScreen {
	case "logo":
		return updateLogo(msg, &m)
	case "reg":
		return updateRegScreen(msg, &m)
	case "auth":
		return updateAuthScreen(msg, &m)
	case "wait":
		return updateWaitScreen(msg, &m)
	}
	return updateLogo(msg, &m)
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

// updateWait
func updateWaitScreen(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return m, nil
}

// updateLogo loop для первого представления
func updateLogo(msg tea.Msg, m *Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
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
				m.Chosen = false
				m.OTPMessageSended = true
				m.CurrentScreen = "wait"
				m.User = core.User{Login: m.RegScreen.Inputs[0].View(), Password: m.RegScreen.Inputs[1].Value()}
				return m, nil //tea.Quit
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
				m.Chosen = false
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
