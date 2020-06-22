package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)
func QuickSortStr(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]
	low := make([]string, 0, 0)
	high := make([]string, 0, 0)
	mid := make([]string, 0, 0)
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

	low, high = QuickSortStr(low), QuickSortStr(high)
	arr = append(append(low, mid...), high...)

	return arr
}
func main() {
	t1 := time.Now()
	const N=6428633
	allstrs := make([]string, N, N)
	fi, err := os.Open("D:\\数据结构\\day2\\CSDNpassword.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		linestr := string(line)
		allstrs[i] = linestr
		i++
	}
	allstrs = QuickSortStr(allstrs)
	path := "D:\\数据结构\\day2\\CSDNpasswordsort.txt"
	savefile, err := os.Create(path)
	defer savefile.Close()
	save := bufio.NewWriter(savefile)

	for i := 0; i < len(allstrs); i++ {
		fmt.Fprintln(save, allstrs[i])
	}
	save.Flush()
	usedTime := time.Since(t1)
	fmt.Println(usedTime)
}
