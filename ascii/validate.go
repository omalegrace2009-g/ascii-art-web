package ascii

import "fmt"

func ValidateInput(s string) (rune, error) {
	for _, rn := range s {
		if rn < 32 || rn > 126 {
			return rn, fmt.Errorf("Inavlid Character %c", rn)
		}
	}
	return 0, nil
}