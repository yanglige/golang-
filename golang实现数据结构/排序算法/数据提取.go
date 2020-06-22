package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fi, err := os.Open("D:\\数据结构\\day2\\CSDN-中文IT社区-600万.sql")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close()

	path := "D:\\数据结构\\day2\\CSDNpassword.txt"
	savefile, err := os.Create(path)
	defer savefile.Close()
	save := bufio.NewWriter(savefile)



	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		linestr := string(line)
		mystrs := strings.Split(linestr, " # ")
		fmt.Fprintln(save, mystrs[1])
	}
	save.Flush()
}


