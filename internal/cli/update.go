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

	// Всегда перехватывем esc (возвращаемя в основные меню)
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "esc" {
			if m.Authenticated {
				m.CurrentScreen = "main"
			} else {
				m.CurrentScreen = "logo"
			}

			return m, nil
		}
	}

	switch m.CurrentScreen {
	case "logo":
		return updateLogo(msg, &m)
	case "reg":
		return updateReg(msg, &m)
	case "auth":
		return updateAuth(msg, &m)
	case "verify":
		return updateVerify(msg, &m)
	case "main":
		return updateMain(msg, &m)
	case "cardCreate":
		return updateCardCreate(msg, &m)
	case "cardList":
		return updateCardList(msg, &m)
	case "coupleCreate":
		return updateCoupleCreate(msg, &m)
	case "coupleList":
		return updateCoupleList(msg, &m)
	case "textCreate":
		return updateTextCreate(msg, &m)
	case "textList":
		return updateTextList(msg, &m)
	case "binaryPick":
		return updateBinaryPick(msg, &m)
	case "binaryAdd":
		return updateBinaryAdd(msg, &m)
	case "binaryList":
		return updateBinaryList(msg, &m)
	default:
		return updateLogo(msg, &m)
	}
}
