package piscine

func UltimatePointOne(n ***int) {
	var a int
	var b *int
	a = 1
	*b = a
	**n = &*b
}
