package main

import "fmt"

func OddEven(arr []int) []int {
	isSorted := false
	for ; isSorted == false; {
		isSorted = true
		for i := 1; i < len(arr) - 1; i = i + 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		for i := 0; i < len(arr) - 1; i = i + 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr
}
func main() {
	arr := []int{1, 9, 3, 8, 2, 6, 4, 5, 7}
	fmt.Println(OddEven(arr))
}


