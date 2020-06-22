package main

import (
	"fmt"

)
func QuickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, splitData)
	for i := 1; i < len(arr); i++ {
		if arr[i] > splitData {
			high = append(high, arr[i])
		}else if arr[i] < splitData{
			low = append(low, arr[i])
		}else {
			mid = append(mid, arr[i])
		}
	}
	low, high = QuickSort2(low), QuickSort2(high)

	arr = append(append(low, mid...), high...)


	return arr
}
func bin_search(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > data {
			high = mid - 1
		}else if arr[mid] < data {
			low = mid + 1
		}else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{9, 1, 3, 8, 2, 6, 4, 5, 7}
	arr1 := QuickSort2(arr)
	fmt.Println(bin_search(arr1, 4))
}