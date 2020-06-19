package main

import "fmt"

func InsertTest(arr []int) []int {
	if len(arr)<=1 {
		return arr
	}
	for i:=1; i<len(arr); i++ {
		backup := arr[i]
		j := i - 1
		for j>=0 && backup<arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = backup
	}// 从上一个位置循环找到位置插入

	return arr
}


func main() {
	arr := []int{9, 1, 3, 8, 2, 6, 4, 5, 7}
	fmt.Println(InsertTest(arr))
}