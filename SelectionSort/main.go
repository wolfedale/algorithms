package main

import "fmt"

func main() {
	array := []int{6, 1, 7, 8, 9, 3, 5, 4, 2}
	var temp int
	fmt.Println(array)
	for i := 0; i < len(array)-1; i++ {
		// min = 0
		min := i
		for j := i + 1; j < len(array); j++ {
			// array[j] = 1
			// array[min] = 6
			// if 1 < 6
			fmt.Println(array[j], "<", array[min])
			if array[j] < array[min] {
				// 0 = 1
				min = j
				fmt.Println("min: ", min)
			}
		}
		// if 1 != 0
		if min != i {
			fmt.Println("o:", array[i], array[min])
			temp = array[i]
			// temp = 6
			array[i] = array[min]
			// 6 = 1
			array[min] = temp
			// 1 = 6
		}
	}
	fmt.Println(array)
}
