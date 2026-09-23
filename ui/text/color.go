package text

func Color(text string, code string) string {
	return "\033[" + code + "m" + text + "\033[0m"
}

func Red(text string) string {
	return Color(text, "31")
}

func Green(text string) string {
	return Color(text, "32")
}

func Yellow(text string) string {
	return Color(text, "33")
}

func Blue(text string) string {
	return Color(text, "34")
}
