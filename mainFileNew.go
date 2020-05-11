package main

import (
	"fmt"
	"io/ioutil"
)



// GetALl: 递归遍历文件生成Tree
func GetAllX(path string, level int) {
	fmt.Println("level", level)
	leveStr := ""
	if level == 1 {
		leveStr = "+"
	} else {
		for ; level > 1; level-- {
			leveStr += "|--"
		}
		leveStr += "+"
	}

	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return
	}
	for _, fi := range read { // 循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fullDir := path + "/" + fi.Name()      // 构造新的路径
			fmt.Println(leveStr+fullDir)
			//newLevel := level + 1
			fmt.Println("level in for ", level)
			GetAllX(fullDir,level + 1) //文件夹递归处理

		} else {
			fullDir := path + "/" + fi.Name() // 构造新的路径
			fmt.Println(leveStr+fullDir)
		}
	}
	return
}

func main() {
	path := "/Users/Natee/Documents"
	GetAllX(path, 1)
}

func mians() {
	// 深度

}
