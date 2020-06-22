package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)
const N  = 1833381
type uuu struct {
	user string
	md5 string
	email string
	password string
}

func QuickSort3(arr []uuu) []uuu {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]
	low := make([]uuu, 0, 0)
	high := make([]uuu, 0, 0)
	mid := make([]uuu, 0, 0)
	mid = append(mid, splitData)
	for i := 1; i < len(arr); i++ {
		if arr[i].user > splitData.user {
			high = append(high, arr[i])
		}else if arr[i].user < splitData.user {
			low = append(low, arr[i])
		}else {
			mid = append(mid, arr[i])
		}
	}
	low, high = QuickSort3(low), QuickSort3(high)
	arr = append(append(low, mid...), high...)
	return arr
}
func bin_searchu(arr []uuu, data string) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid].user > data {
			high = mid - 1
		}else if arr[mid].user < data {
			low = mid + 1
		}else {
			return mid
		}
	}
	return -1
}
func main() {
	alldata := make([]uuu, N, N)

	path := "D:\\数据结构\\day3\\uuu9.com.sql"
	sqlfile, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	br := bufio.NewReader(sqlfile)
	i := 0
	for {
		linestr, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		linestr1 := string(linestr)
		linestr2 := strings.Split(linestr1, " | ")
		//fmt.Println(linestr2)
		if len(linestr2) == 4 {
			alldata[i].user = linestr2[0]
			alldata[i].md5 = linestr2[1]
			alldata[i].email = linestr2[2]
			alldata[i].password = linestr2[3]
		}

		if i > N {
			break
		}
		i++

	}
	alldata = QuickSort3(alldata)
	for {
		fmt.Println("请输入要搜索的用户")
		var inputstr string
		fmt.Scanln(&inputstr)
		starttime := time.Now()
		fmt.Println(bin_searchu(alldata, inputstr))
		fmt.Println(alldata[bin_searchu(alldata, inputstr)])
		fmt.Println(time.Since(starttime))
	}
}
