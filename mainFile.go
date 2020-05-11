package main

import (
	"./StackArray"
	"errors"
	"fmt"
	"io/ioutil"
	"./Queue"
)

// GetALl: 递归遍历文件
func GetAll(path string, files [] string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹读取失败")
	}
	for _, fi := range read { // 循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fullDir := path + "/" + fi.Name() // 构造新的路径
			files = append(files, fullDir)    //追加路径
			files, _ = GetAll(fullDir, files) //文件夹递归处理

		} else {
			files = append(files, fi.Name())
		}
	}
	return files, nil

}

func mainx() {
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
				mystack.Push(fullDir)
			} else {
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

//func main() {
//	newList := []string{}
//	files, _ := GetAll("/Users/Natee/Documents", newList)
//	for i:=0;i<len(files);i++{
//		fmt.Println(files[i])
//	}
//}

func main()  {
	path := "/Users/Natee/Documents"
	files := []string{} //数组字符串
	mystack := Queue.NewQueue()
	mystack.EnQueue(path)
	for  {
		path := mystack.DeQueue()
		if path == nil{
			break
		}
		files = append(files, path.(string))     //加入列表
		read, _ := ioutil.ReadDir(path.(string)) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path.(string) + "/" + fi.Name()
				mystack.EnQueue(fullDir)
			} else {
				fullDir := path.(string) + "/" + fi.Name()

				files = append(files, fullDir)
			}

		}

	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
	fmt.Println(len(files))
}