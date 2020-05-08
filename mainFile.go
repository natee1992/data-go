package main

import (
	"./StackArray"
	"fmt"
	"io/ioutil"
)

func main() {
	path := "/Users/Natee/Documents"
	files := []string{} //数组字符串
	mystack := StackArray.NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		path := mystack.Pop().(string)
		files = append(files, path)     //加入列表
		read, _ := ioutil.ReadDir(path) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path + "/" + fi.Name()
				//fmt.Println(fullDir)
				mystack.Push(fullDir)
			}else {
				fullDir := path + "/" + fi.Name()
				files = append(files, fullDir)
			}

		}

	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
	fmt.Println(len(files))
}
