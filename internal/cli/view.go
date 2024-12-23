package cli

// View (Представление) — это функция, которая отображает модель в виде, понятном пользователю.
func (m Model) View() string {
	s := ""

	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	switch m.CurrentScreen {
	case "logo":
		s = viewLogo(&m)
	case "reg":
		s = viewReg(&m)
	case "auth":
		s = viewAuth(&m)
	case "verify":
		s = viewVerify(&m)
	case "main":
		s = viewMain(&m)
	case "cardCreate":
		s = viewCardCreate(&m)
	case "cardList":
		s = viewCardList(&m)
	}
	return mainStyle.Render("\n" + s + "\n\n")
}
