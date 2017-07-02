package main

import "fmt"

func main() {
	array := []int{5, 9, 3, 1, 2, 8, 4, 7, 6}
	var temp int

	fmt.Println(array)

	// iterate over full array
	for j := 0; j < len(array); j++ {
		// iterate over array - 1 so I can use i+1
		for i := 0; i < len(array)-1; i++ {
			fmt.Println(array[i], ">", array[i+1])
			// 9 > 3
			if array[i] > array[i+1] {
				// temp = 3
				temp = array[i+1]
				fmt.Println("temp: ", temp)
				array[i+1] = array[i]
				// 3 = 9
				fmt.Println("array[i+1]: ", array[i+1])
				array[i] = temp
				// 9 = 3
				fmt.Println("array[i]: ", array[i])
			}
		}
	}
	fmt.Println(array)
}
