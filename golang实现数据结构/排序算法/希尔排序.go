package main

import "fmt"

func ShellSortSub(arr []int, start int, gap int) {
	for i := start + gap; i < len(arr); i += gap {
		backup := arr[i]
		j := i - gap
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j-=gap
		}
		arr[j+gap] = backup
	}
}

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}else {
		gap := length/2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				ShellSortSub(arr, i, gap)
			}
			gap--
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 3, 8, 2, 6, 4, 5, 7}
	fmt.Println(ShellSort(arr))
}
