package main

import (
	"fmt"
	"strings"
)

func SelectSortMaxString(arr[]string) string {
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

func SelectSortString(arr []string) []string {
	length := len(arr)
	if length<=1 {
		return arr
	}else {
		for i:=0; i<length-1; i++ {
			min := i
			for j:=i+1; j<length; j++ {
				if strings.Compare(arr[min], arr[j])>0 {
					min=j

				}
				//if arr[min]<arr[j] {
				//	min = j
				//}
			}
			if i!=min {
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []string{"a", "c", "j", "d", "e", "n", "o"}
	fmt.Println(SelectSortString(arr))

}
