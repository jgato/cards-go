package main

import "fmt"

func main() {
	numbers := [10]int{}

	for i := 0; i < 10; i++ {
		numbers[i] = i

	}

	for i := 0; i < 10; i++ {
		if numbers[i]%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}

	}
}
