package piscine

func ForEach(f func(int), a []int) {
	for index, element := range a {
		println(index, element)
	}
}
