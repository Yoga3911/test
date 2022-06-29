package test

import (
	"fmt"
	"strings"
)

var array []int
var first int
var action string

func Oscillated(value int) string {
	if len(array) == 0 {
		array = append(array, value)
		first = value
		action = "-"
	}

	var result int
	if value > 0 && action == "-" {
		result = value - 5
		array = append(array, result)
	} else {
		action = "+"
		result = value + 5
		array = append(array, result)
	}

	if result != first {
		Oscillated(result)
	}

	return strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", array), "[", ""), "]", "")
}
