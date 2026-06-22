package ascii

import (
	"strings"
)

func GenerateArt(s string, banner map[rune][]string) string {
	if s == "" {
		return ""
	}

	line := strings.ReplaceAll(s, "\r\n", "\n")

	for _, rn := range line {
		if (rn < 32 || rn > 126) && rn != '\n' {
			return "Error: Invalid Character"
		}
	}

	sp := strings.Split(line, "\n")

	for _, ch := range sp {
		if ch != "" {
			break
		}
	}

	var result strings.Builder
	for _, st := range sp {
		if st == "" {
			result.WriteString("\n")
		} else {
			r := RenderLine(st, banner)
			for _, l := range r {
				result.WriteString(l)
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
