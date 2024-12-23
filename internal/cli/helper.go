package cli

func footer() string {
	return subtleStyle.Render("\n\nДля выхода нажмите ") +
		keywordStyle.Render("ctrl+c") +
		subtleStyle.Render("\nДля сброса нажмите ") +
		keywordStyle.Render("esc")
}

func footerTable() string {
	return subtleStyle.Render("\n\nДля того, чтобы обновить данные, нажмите ") +
		keywordStyle.Render("r") +
		subtleStyle.Render("\nДля того, чтобы создать новую запись, нажмите ") +
		keywordStyle.Render("с")
}

func err(s string) string {
	return keywordStyle.Render(s)
}

func (m Model) header() string {
	return subtleStyle.Render("Текущий экран: ") + keywordStyle.Render(m.CurrentScreen) + "\n\n"
}
