package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	length := 0
	for i := range a {
		length = i + 1
	}

	ascendant := true
	descendant := true

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			descendant := false
		}
	}
	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			ascendant = false
		}
	}
	return ascendant || descendant
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
