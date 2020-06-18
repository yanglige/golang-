package main

import (
	"errors"
	"fmt"
	"golang实现数据结构/Queue"
	"io/ioutil"
)


func GetAl(path string, files[] string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _,fi := range read {
		if fi.IsDir() {
			fulldir := path+"\\"+fi.Name()
			files = append(files, fulldir)
			files,_ = GetAl(fulldir, files)
		}else {
			fulldir := path+"\\"+fi.Name()
			files = append(files, fulldir)
		}
	}
	return files, nil
}
func main() {
	path := "C:\\Windows"
	files := []string{}
	myq := Queue.NewQueue()
	myq.EnQueue(path)
	for  {
		path := myq.DeQueue()
		if path == nil {
			break
		}
		//files = append(files, path.(string))
		read, _ := ioutil.ReadDir(path.(string))
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path.(string)+"\\"+fi.Name()
				files = append(files, fulldir)
				myq.EnQueue(fulldir)
			}else {
				fulldir := path.(string)+"\\"+fi.Name()
				files = append(files, fulldir)
			}
		}
	}
	for i:=0; i<len(files); i++ {
		fmt.Println(files[i])
	}
}