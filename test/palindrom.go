package test

import (
	"fmt"
	"strings"
)

func Palindrom(text string) {
	var kata string = text
	real := make([]string, 0)
	palindrom := make([]string, 0)
	for i := len(kata) - 1; i > -1; i-- {
		palindrom = append(palindrom, string(kata[i]))
	}
	for j := 0; j < len(kata); j++ {
		real = append(real, string(kata[j]))
	}
	hasil1 := strings.Join(real, " ")
	hasil2 := strings.Join(palindrom, " ")
	if hasil1 == hasil2 {
		fmt.Println("Palindrom")
		return
	}
	fmt.Println("Bukan Palindrom")
}
