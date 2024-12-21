package cli

func footer() string {
	return subtleStyle.Render("\n\nДля выхода нажмите ") + keywordStyle.Render("ctrl+c") + subtleStyle.Render("\nДля сброса нажмите ") + keywordStyle.Render("esc")
}

func err(s string) string {
	return keywordStyle.Render(s)
}

func (m Model) header() string {
	return subtleStyle.Render("Текущий экран: ") + keywordStyle.Render(m.CurrentScreen) + "\n\n"
}
