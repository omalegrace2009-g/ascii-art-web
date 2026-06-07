package ascii

func RenderLine(text string, banner map[rune][]string) []string {
	word := make([]string, 8)
	for i := 0; i <= 7; i++ {
		for _, r := range text {
			word[i] += banner[r][i]
		}
	}
	return word
}
