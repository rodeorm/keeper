package cli

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
	case "verify":
		s = verifyView(&m)
	case "main":
		s = mainView(&m)
	}
	return mainStyle.Render("\n" + s + "\n\n")
}
