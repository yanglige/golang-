package main

import "fmt"


func SelectSortMax(arr[]int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	}else {
		max := arr[0]
		for i:=0; i<length; i++ {
			if arr[i]>max {
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSort(arr []int) []int {
	length := len(arr)
	if length<=1 {
		return arr
	}else {
		for i:=0; i<length-1; i++ {
			min := i
			for j:=i+1; j<length; j++ {
				if arr[min]<arr[j] {
					min = j
				}
			}
			if i!=min {
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 3, 8, 2, 6, 4, 5, 7}
	fmt.Println(SelectSort(arr))
}