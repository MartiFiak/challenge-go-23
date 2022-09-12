package piscine

func UltimatePointOne(n ***int) {
	var a int = 1
	var b *int = &a
	***n = *b
}
