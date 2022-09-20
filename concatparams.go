package piscine

func ConcatParams(args []string) string {
	str := ""
	for a, i := range args[0:] {
		str += i
		if len(args)-1 != a {
			str += string("\n")
		}

	}
	return str
}
