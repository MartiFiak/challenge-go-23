package main

import (
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	} else if len(arg[0]) >= 19 || len(arg[2]) >= 19 {
		return
	}
	if arg[1] == "+" {
		nb1 := atoi(arg[0])
		nb2 := atoi(arg[2])
		if nb1 == 0 && arg[0] != "0" || nb2 == 0 && arg[2] != "0" {
			return
		}
		if nb1 < 0 && nb2 < 0 {
			os.Stdout.WriteString("-" + int_to_str(-nb1+-nb2) + "\n")
		} else if nb2 < 0 && -nb2 > nb1 {
			os.Stdout.WriteString("-" + int_to_str(-nb1-nb2) + "\n")
		} else if nb1 < 0 {
			os.Stdout.WriteString("-" + int_to_str(-nb1-nb2) + "\n")
		} else {
			os.Stdout.WriteString(int_to_str(nb1+nb2) + "\n")
		}
	} else if arg[1] == "-" {
		nb1 := atoi(arg[0])
		nb2 := atoi(arg[2])
		if nb1 == 0 && arg[0] != "0" || nb2 == 0 && arg[2] != "0" {
			return
		}
		if nb1 < 0 && nb2 < 0 {
			if nb1 < nb2 {
				os.Stdout.WriteString("-" + int_to_str(nb2-nb1) + "\n")
			} else {
				os.Stdout.WriteString(int_to_str(nb1+-nb2) + "\n")
			}
		} else if nb1 < nb2 {
			os.Stdout.WriteString("-" + int_to_str(nb2-nb1) + "\n")
		} else {
			os.Stdout.WriteString(int_to_str(nb1-nb2) + "\n")
		}
	} else if arg[1] == "*" {
		nb1 := atoi(arg[0])
		nb2 := atoi(arg[2])
		if nb1 == 0 && arg[0] != "0" || nb2 == 0 && arg[2] != "0" {
			return
		}
		if nb1 < 0 && nb2 < 0 {
			os.Stdout.WriteString(int_to_str(-nb1*-nb2) + "\n")
		} else if nb2 < 0 {
			os.Stdout.WriteString("-" + int_to_str(nb1*-nb2) + "\n")
		} else if nb1 < 0 {
			os.Stdout.WriteString("-" + int_to_str(-nb1*nb2) + "\n")
		} else {
			os.Stdout.WriteString(int_to_str(nb1*nb2) + "\n")
		}
	} else if arg[1] == "/" {
		if arg[2] == "0" {
			os.Stdout.WriteString("No division by 0\n")
		} else {
			nb1 := atoi(arg[0])
			nb2 := atoi(arg[2])
			if nb1 == 0 && arg[0] != "0" || nb2 == 0 && arg[2] != "0" {
				return
			}
			if nb1 < 0 && nb2 < 0 {
				os.Stdout.WriteString(int_to_str(-nb1/-nb2) + "\n")
			} else if nb2 < 0 {
				os.Stdout.WriteString("-" + int_to_str(nb1/-nb2) + "\n")
			} else if nb1 < 0 {
				os.Stdout.WriteString("-" + int_to_str(-nb1/nb2) + "\n")
			} else {
				os.Stdout.WriteString(int_to_str(nb1/nb2) + "\n")
			}
		}
	} else if arg[1] == "%" {
		if arg[2] == "0" {
			os.Stdout.WriteString("No modulo by 0\n")
		} else {
			nb1 := atoi(arg[0])
			nb2 := atoi(arg[2])
			if nb1 == 0 && arg[0] != "0" || nb2 == 0 && arg[2] != "0" {
				return
			}
			if nb1 < 0 && nb2 < 0 {
				os.Stdout.WriteString(int_to_str(-nb1%-nb2) + "\n")
			} else if nb2 < 0 {
				os.Stdout.WriteString("-" + int_to_str(nb1%-nb2) + "\n")
			} else if nb1 < 0 {
				os.Stdout.WriteString("-" + int_to_str(-nb1%nb2) + "\n")
			} else {
				os.Stdout.WriteString(int_to_str(nb1%nb2) + "\n")
			}
		}
	}
}

func atoi(s string) int {
	srune1 := []rune(s)
	srune := []rune{}
	neg := false
	pos := false
	for v, c := range srune1 {
		if int(c) == 45 {
			if v == 0 {
				if neg == true {
					return 0
				} else {
					neg = true
				}
			} else {
				return 0
			}
		} else if int(c) == 43 {
			if v == 0 {
				if neg == true {
					return 0
				}
				if pos == true {
					return 0
				}
				pos = true
			} else {
				return 0
			}
		} else {
			srune = append(srune, c)
		}
	}
	sint := []int{}
	for _, c := range srune {
		if int(c)-48 <= 9 && int(c)-48 >= 0 {
			sint = append(sint, int(c)-48)
		} else {
			return 0
		}
	}
	result := 0
	for _, c := range sint {
		result = result*10 + c
	}
	if neg == true {
		return result / -1
	}
	return result
}

func int_to_str(nb int) string {
	if nb/10 == 0 {
		return string(rune(nb) + 48)
	}
	return int_to_str(nb/10) + string(rune(nb%10)+48)
}
