package piscine

func Index(s string, toFind string) int {
	startOccur := -1
	if len(s) == 0 || len(toFind) == 0 {
		return 0
	}
	for i, c := range []rune(s) {
		if c == rune(toFind[0]) {
			startOccur := i
			for j, v := range []rune(toFind) {
				if v != rune(s[startOccur+j]) {
					startOccur = -1
				}
			}
			if startOccur != -1 {
				return startOccur
			}
		}
	}
	return startOccur
}
