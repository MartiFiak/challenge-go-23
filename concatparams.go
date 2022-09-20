package piscine

func ConcatParams(args []string) string {
	str := ""
	for _, i := range args[0:] {
		str += i
		str += string("\n")
	}
	return str
}
