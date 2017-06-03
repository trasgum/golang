package main

import (
	"strings"
	"strconv"
	"fmt"
	"errors"
)

var (
	NegativeNumber = errors.New("negatives not allowed")
)

func Add(numbers string) (suma int64, err error) {
	var sumandos[]string
	var sep string
	var neg_numbers[]int64

	nums := strings.Split(numbers, "\n")
	if strings.Contains(nums[0], "//") {
		sep = strings.Split(nums[0], "//")[1]
		nums = nums[1:]
	} else {
		sep = ","
	}

	for _, str_num := range nums {
		sumandos = append(sumandos, strings.Split(str_num, sep)...)
	}

	for _, sumando := range sumandos {
		sumando_int, _ := strconv.ParseInt(sumando, 10, 64)
		if sumando_int < 0 {
			neg_numbers = append(neg_numbers, sumando_int)
		}
		suma += sumando_int
	}
	if len(neg_numbers) > 0 {
		panic(fmt.Sprintf("%v", neg_numbers))

	}
	return
}

func main() {
	//fmt.Printf("Result: %d\n", Add(""))
	//fmt.Printf("Result: %d\n", Add("1,2"))
	//fmt.Printf("Result: %d\n", Add("1,\n"))
	sum, _ := Add("//|\n100|10")
	fmt.Printf("Result: %d\n", sum)
	//Add("-1")
}