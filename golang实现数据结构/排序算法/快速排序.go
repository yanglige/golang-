package main

import "fmt"

func QuickSort(arr []int) []int {
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
	//fmt.Println(low)
	//	//fmt.Println(high)
	//	//fmt.Println(mid)
	low, high = QuickSort(low), QuickSort(high)
	fmt.Println("--", low)
	fmt.Println("--", high)
	fmt.Println(mid)

	arr = append(append(low, mid...), high...)
	fmt.Println(append(low, mid...))

	return arr
}
func main() {
	arr := []int{1, 9, 3, 8, 2, 6, 4, 5, 7}
	fmt.Println(QuickSort(arr))
}


