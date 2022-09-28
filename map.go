package piscine

func Map(f func(int) bool, a []int) []bool {
	st := []bool{}
	for _, i := range a {
		st = append(st, f(i))
	}
	return st
}
