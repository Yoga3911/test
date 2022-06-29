package main

import (
	"coba/test"
	"fmt"
)

func main() {
	var array []int
	array = append(array, 3, 5, 6, 2, 3, 9, 12, 4, 3, 7)
	fmt.Println(test.ElemenKecilBesar(array))
}
