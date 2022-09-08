package piscine

func UltimatePointOne(n ***int) {
	var a int
	var b *int
	*b = 1
	a = *b
	**n = &a
}
