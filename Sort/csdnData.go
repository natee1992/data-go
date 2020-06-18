package Sort

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func SortCsdn(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close()

	path := "CSDNmail.txt"

	writerFi, err := os.Create(path)
	defer writerFi.Close()
	save := bufio.NewWriter(writerFi)

	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Println(string(line))
		lineStr := string(line)
		myStrs := strings.Split(lineStr, "#")
		fmt.Fprintln(save, myStrs[2])
	}
	save.Flush()
}
