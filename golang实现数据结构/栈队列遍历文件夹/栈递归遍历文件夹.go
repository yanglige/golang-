package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// 递归文件夹

func GetAll(path string, files[] string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _,fi := range read {
		if fi.IsDir() {
			fulldir := path+"\\"+fi.Name()
			files = append(files, fulldir)
			files,_ = GetAll(fulldir, files)
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
	files,_ = GetAll(path, files)
	for i:=0; i<len(files); i++ {
		fmt.Println(files[i])
	}

}
//func main() {
//	path := "C:\\Windows"
//	files := []string{}
//	mystack := StackArray.NewStack()
//	mystack.Push(path)
//	for !mystack.IsEmpty() {
//		path := mystack.Pop().(string)
//		files = append(files, path)
//		read, _ := ioutil.ReadDir(path)
//		for _, fi := range read {
//			if fi.IsDir() {
//				fulldir := path+"\\"+fi.Name()
//				files = append(files, fulldir)
//				mystack.Push(fulldir)
//			}else {
//				fulldir := path+"\\"+fi.Name()
//				files = append(files, fulldir)
//			}
//		}
//	}
//	for i:=0; i<len(files); i++ {
//		fmt.Println(files[i])
//	}
//}