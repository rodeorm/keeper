package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update (Обновление) — это функция, которая принимает текущее состояние (модель) и сообщение (например, событие, вызванное пользователем),
// и возвращает новое состояние.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Всегда перехватывем ctl+c (выходим из приложения)
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}
	// // Всегда перехватывем esc (обнуляем приложение)
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "esc" {
			m = InitialModel(m.sc)
			return m, nil
		}
	}

	switch m.CurrentScreen {
	case "logo":
		return updateLogo(msg, &m)
	case "reg":
		return updateRegScreen(msg, &m)
	case "auth":
		return updateAuthScreen(msg, &m)
	case "verify":
		return updateVerifyScreen(msg, &m)
	case "main":
		return updateMain(msg, &m)
	default:
		return updateLogo(msg, &m)
	}
}
