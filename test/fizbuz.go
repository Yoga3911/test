package test

import "fmt"

func FizBuz(value int) {
	for i := 0; i < value; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
