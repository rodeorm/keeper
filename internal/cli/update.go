package cli

import (
	tea "github.com/charmbracelet/bubbletea"
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
	case "main":
		return updateMain(msg, &m)
	default:
		return updateLogo(msg, &m)
	}
}
