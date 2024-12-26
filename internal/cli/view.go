package cli

// View (Представление) — это функция, которая отображает модель в виде, понятном пользователю.
func (m Model) View() string {
	var s string

	if m.Quitting {
		return "\n  До свидания!\n\n"
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
	case "coupleCreate":
		s = viewCoupleCreate(&m)
	case "coupleList":
		s = viewCoupleList(&m)
	case "textCreate":
		s = viewTextCreate(&m)
	case "textList":
		s = viewTextList(&m)
	case "binaryPick":
		s = viewBinaryPick(&m)
	case "binaryAdd":
		s = viewBinaryAdd(&m)
	case "binaryList":
		s = viewBinaryList(&m)
	}
	return mainStyle.Render(s)
}
