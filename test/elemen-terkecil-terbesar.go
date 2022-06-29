package test

import (
	"sort"
	"strconv"
)

var temp []string
var min int

func ElemenKecilBesar(array []int) []string {
	sortArr := make([]int, len(array))
	copy(sortArr, array)
	sort.Ints(sortArr)
	for _, value := range array {
		index := indexOf(value, sortArr)
		if index == len(array)-1 {
			temp = append(temp, "^")
			continue
		}
		for _, value2 := range sortArr {
			if value2 == sortArr[index] {
				index += 1
				if value2 < sortArr[index] {
					break
				}
			}
		}
		val := sortArr[index]
		temp = append(temp, strconv.Itoa(val))
	}

	return temp
}

func indexOf(word int, data []int) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
