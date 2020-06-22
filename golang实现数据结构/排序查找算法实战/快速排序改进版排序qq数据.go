package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const N = 84331445

type QQ struct {
	QQuser int
	password string
}


func SortForMergeQQ(arr []QQ, left int, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i]
		var j int
		for j = i; j > left && arr[j-1].QQuser > temp.QQuser; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func swapQQ(arr []QQ, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSortXqq(arr []QQ, left int, right int) {
	if right - left < 15 {
		SortForMergeQQ(arr, left, right)
	}else {
		swapQQ(arr, left, rand.Int()%(right-left+1)+left)
		vdata := arr[left]
		It := left
		gt := right + 1
		i := left + 1
		for i < gt {
			if arr[i].QQuser < vdata.QQuser {
				swapQQ(arr, i, It+1)
				It++
				i++
			}else if arr[i].QQuser>vdata.QQuser{
					swapQQ(arr, i, gt-1)
					gt--
			}else {
						i++
			}
		}
		swapQQ(arr, left, It)
		QuickSortXqq(arr, left, It-1)
		QuickSortXqq(arr, gt, right)
	}
}

func QuickSortPlusQQ(arr []QQ) {
	QuickSortXqq(arr, 0, len(arr)-1)

}


func bin_searchQQX(arr[]QQ, data int) int {
	low := 0
	high := len(arr)-1
	for low <= high {
		mid := (low+high)/2
		if arr[mid].QQuser > data {
			high = mid - 1
		}else if arr[mid].QQuser < data {
			low = mid+1
		}else {
			return mid
		}
	}
	return -1

}


func main() {
	alldata := make([]QQ, N, N)
	path := "D:\\数据结构\\day3\\QQ.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件错误")
	}
	i := 0
	br := bufio.NewReader(file)
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}
		linestr := string(line)
		lines := strings.Split(linestr, "----")
		if len(lines) == 2 {
			alldata[i].QQuser,_ = strconv.Atoi(lines[0])
			alldata[i].password = lines[1]
		}
		i++
	}
	starttime := time.Now()
	QuickSortPlusQQ(alldata)
	fmt.Println("排序花了",time.Since(starttime))
	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr int
		fmt.Scanf("%d", &inputstr)
		starttime := time.Now()
		index := bin_searchQQX(alldata, inputstr)
		fmt.Println("index", index)
		if index == -1 {
			fmt.Println("找不到")
		}else {
			fmt.Println("找到", alldata[index])
		}
		fmt.Println("本次查询用了", time.Since(starttime))
	}

}

