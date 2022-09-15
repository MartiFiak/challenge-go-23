package piscine

func TrimAtoi(s string) int {
	intr := 0
	neg := false
	for _, c := range s {
		if string(c) == "-" {
			if neg {
				neg = false
			} else if !neg && intr == 0 {
				neg = true
			}
		}
		if IsNumeric(string(c)) {
			intr *= 10
			intr += int(c) - 48
		}
	}
	if neg {
		return -intr
	} else {
		return intr
	}
}
