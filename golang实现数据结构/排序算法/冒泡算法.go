package main

import "fmt"


func bubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i:=0; i<len(arr)-1; i++ {
		isneedChange := false
		for j:=0; j<len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isneedChange = true
			}
		}
		if !isneedChange {
			break
		}
		fmt.Println(arr)

	}
	return arr
}


func main() {
	arr := []int{9, 1, 3, 8, 2, 6, 4, 5, 7}
	fmt.Println(bubbleSort(arr))
}