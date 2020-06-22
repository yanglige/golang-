package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const N  = 18333811


type uuu9 struct {
	user string
	md5 string
	email string
	password string
}


func main() {


	alldata := make([]uuu9, N, N)

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


		i++
	}
	for {
		fmt.Println("请输入要搜索的用户")
		var inputstr string
		fmt.Scanln(&inputstr)
		starttime := time.Now()

		for i := 0; i < N; i++ {
			if alldata[i].user == inputstr {
				fmt.Println(alldata[i])
				fmt.Println(time.Since(starttime))
			}
		}

	}

}
