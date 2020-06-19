package main

import "fmt"


func HeapSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		lastlength := len(arr) - i
		HeapSortMax(arr, lastlength)
		if i < len(arr) {
			arr[0], arr[lastlength-1] = arr[lastlength-1], arr[0]
		}
	}
	return arr
}
func HeapSortMax(arr []int, length int) int {
	if length <= 1 {
		return arr[0]
	}
	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		tmpmax := i
		if left <= length - 1 && arr[left] > arr[tmpmax]  {
			tmpmax = left
		}
		if right <= length - 1 && arr[right] > arr[tmpmax] {
			tmpmax = right
		}
		if tmpmax != i {
			arr[i], arr[tmpmax] = arr[tmpmax], arr[i]
		}
	}

	return arr[0]
}

func main() {
	arr := []int{1, 9, 3, 8, 2, 6, 4, 5, 7}
	fmt.Println(HeapSort(arr))
}