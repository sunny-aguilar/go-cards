package main

import (
	"fmt"
)

func evenOdd() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range s {
		if val%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
