package main

import (
	"fmt"
	"math/rand"
)

func SortForMerge(arr []int, left int, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i]
		var j int
		for j = i; j > left && arr[j - 1] > temp; j-- {
			arr[j] = arr[j - 1]
		}
		arr[j] = temp
	}
}


func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSortX(arr []int, left int, right int) {
	if right - left < 2 {
		SortForMerge(arr, left, right)
	}else {
		swap(arr, left, rand.Int()%(right-left+1) + left)
		vdata := arr[left]
		It := left
		gt := right + 1
		i := left + 1
		for i < gt {
			if arr[i] < vdata {
				swap(arr, i, It + 1)
				It++
				i++
			}else if arr[i] > vdata {
					swap(arr, i, gt-1)
					gt--
			}else {
						i++
			}
		}
		swap(arr, left, It)
		QuickSortX(arr, left, It-1)
		QuickSortX(arr, gt, right)
	}
}

func QuickSortPlus(arr []int) {
	QuickSortX(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 9, 3, 8, 2, 6, 4, 5, 7}
	fmt.Println(arr)
	QuickSortPlus(arr)
	fmt.Println(arr)
}