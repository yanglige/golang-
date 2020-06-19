package main

import "fmt"

func merge(arr []int, arr1 []int) []int {
	leftindex := 0
	rightindex := 0
	lastarr := []int {}
	for leftindex < len(arr) && rightindex < len(arr1) {
		if arr[leftindex] < arr1[rightindex] {
			lastarr = append(lastarr, arr[leftindex])
			leftindex++
		}else if arr[leftindex] > arr1[rightindex] {
				lastarr = append(lastarr, arr1[rightindex])
				rightindex++
		}else {
			lastarr = append(lastarr, arr[leftindex])
			lastarr = append(lastarr, arr1[rightindex])
			leftindex++
			rightindex++
		}
	}
	for leftindex < len(arr) {
		lastarr = append(lastarr, arr[leftindex])
		leftindex++
	}
	for rightindex < len(arr1) {
		lastarr = append(lastarr, arr1[rightindex])
		rightindex++
	}
	return lastarr
}
func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}else {
		mid := length/2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])
		return merge(leftarr, rightarr)
	}
}
func main() {
	arr := []int{1, 9, 3, 8, 2, 6, 4, 5, 7, 99, 212, 33, 23}
	fmt.Println(MergeSort(arr))

}
