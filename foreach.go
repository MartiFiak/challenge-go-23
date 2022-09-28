package piscine

import (
	"fmt"
)

func ForEach(f func(int), a []int) {
	for index, element := range a {
		fmt.Println(index, element)
	}
}
