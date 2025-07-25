package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}


func sliceLoop() {
	anotherSlice := []string{"Jason", "Greg", "Steven", "Fernando"}

		for i, name := range anotherSlice {
			if (anotherSlice[i] == "J") {
			fmt.Printf("The name is %v", name)
			}
		}
}