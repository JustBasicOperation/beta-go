package algorithm

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func duplicateString() {
	// 统计所有文件中，特定数据项出现的次数，输出超过一次的数据
	dirName := "/Users/faust/Desktop/62mutants_remove_more_background_mutations"
	resMap := make(map[string][]string, 1000)
	// 获取目录下所有文件
	dirs, _ := ioutil.ReadDir(dirName)
	// 遍历所有文件去重
	for _, dir := range dirs {
		handleFileData(dirName, dir.Name(), resMap)
	}
	srcCount := len(resMap)
	outContent := ""
	count := 0
	for k, v := range resMap {
		if len(v) > 1 {
			count++
			res := fmt.Sprintf("key: %v, value: %v\n", k, v)
			outContent += res
			fmt.Printf(res)
		}
	}
	fmt.Println(count)
	// 输出到文本文件
	outContent += "total duplicated keys:" + strconv.FormatInt(int64(srcCount), 10) +
		", result keys: " + strconv.FormatInt(int64(count), 10)
	if err := ioutil.WriteFile("output.txt", []byte(outContent), 0644); err != nil {
		fmt.Printf("output to file err: %v", err)
	}
}

func handleFileData(dirName, fileName string, resMap map[string][]string) {
	file, err := os.Open(dirName + "/" + fileName)
	if err != nil {
		fmt.Printf("打开文件失败: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	}(file)
	reader := bufio.NewReader(file)
	i := 0
	clo9Slice := make([]string, 0, 1000)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		i++
		if i == 1 {
			// 忽略第一行
			continue
		}
		clo9Slice = append(clo9Slice, strings.Split(string(line), "\t")[9])
	}
	// 去重
	sliceStr := distinctSliceString(clo9Slice)
	// 统计出现次数并保存文件名
	for _, s := range sliceStr {
		if _, ok := resMap[s]; !ok {
			resMap[s] = []string{fileName}
			continue
		}
		resMap[s] = append(resMap[s], fileName)
	}
}

func distinctSliceString(slice []string) []string {
	if len(slice) == 0 {
		fmt.Println("slice is nil")
		return []string{}
	}
	var value int
	var res []string
	tempMap := make(map[string]int)
	for _, i := range slice {
		l := len(tempMap)
		tempMap[i] = value
		if len(tempMap) > l {
			res = append(res, i)
		}
	}
	return res
}
