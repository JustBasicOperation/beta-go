package iostream

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func readLineFromFile(dirName, fileName string) ([]string, error) {
	file, err := os.Open(dirName + "/" + fileName)
	if err != nil {
		fmt.Printf("打开文件失败: %v", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	}(file)
	reader := bufio.NewReader(file)
	var resLine []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		resLine = append(resLine, string(line))
	}
	return resLine, nil
}

func readBytesFromFile(dirName, fileName string) ([]byte, error) {
	file, err := os.Open(dirName + "/" + fileName)
	if err != nil {
		fmt.Printf("打开文件失败: %v", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	}(file)
	reader := bufio.NewReader(file)
	var resBytes []byte
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		resBytes = append(resBytes, line...)
	}
	return resBytes, nil
}

type Ratee struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

func extractKeyData() {
	var rateeList []*Ratee
	data, _ := readBytesFromFile("/Users/faust/Desktop", "test.txt")
	if err := json.Unmarshal(data, &rateeList); err != nil {
		fmt.Printf("unmarshal err: %v", err)
	}
	fmt.Printf("name,language\n")
	for _, ratee := range rateeList {
		fmt.Printf("%v,%v\n", ratee.Name, ratee.Language)
	}
}
