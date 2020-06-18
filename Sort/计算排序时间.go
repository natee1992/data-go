package Sort

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func GetSortTime(filename string) {
	t1 := time.Now()
	const N = 6428632 // 需要开辟的内存

	myStrs := make([]string, N, N)
	fmt.Println(myStrs)

	fi, err := os.Open(filename)
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
		i++
		myStrs = append(myStrs, string(line))
	}
	time.Sleep(time.Second * 100)

	BubbleSort(myStrs)


	//wi,err := os.Create("CsdnSort.txt")
	//bw := bufio.NewWriter(wi)
	//sortMyStrs := InsertSortAll(myStrs)
	//for i:=0;i<len(sortMyStrs);i++{
	//	bw.WriteString(sortMyStrs[i])
	//}
	//fmt.Println(i)
	//bw.Flush()
	used := time.Since(t1)
	fmt.Println(used)
}
