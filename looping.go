package main

import "fmt"

func main() {
	const MAX_ITERATION = 30

	for i := 0; i < MAX_ITERATION; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
