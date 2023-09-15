package beta

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

// 以下代码有什么问题，怎么解决？
func Test20230220_1(t *testing.T) {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("total:%d sum %d\n", total, sum)

}

// TestRegexp 测试regexp包的用法
func TestRegexp(t *testing.T) {
	p := regexp.MustCompile(`a.`)
	//fmt.Println(p.Find([]byte("ababab")))           // 匹配字节数组，返回第一个匹配的结果
	fmt.Println(p.FindString("ababab"))             // 匹配字符串，返回第一个匹配的结果
	fmt.Println(p.FindAllString("ababab", -1))      // 返回所有匹配的结果
	fmt.Println(p.FindAllStringIndex("ababab", -1)) // 一个二维数组，返回所有匹配结果的开始和结束位置

	q, _ := regexp.Compile(`^a(.*)(\d)b$`)
	//fmt.Println(q.FindAllSubmatch([]byte("ababab"), -1))
	// 返回一个二维数组，表示所有子匹配的结果，每个一维数组的第一个元素是全匹配的字符串，第二个元素是子匹配的字符串
	// 二维数组的二维表示多个全匹配的结果以及它们对应的子匹配结果
	// 子匹配说明一个正则表达式中有多个分组表达式，一般用括号括起来，例如：^a(.*)b$ 的子匹配表达式是：(.*)
	fmt.Println(q.FindAllStringSubmatch("ababaxx1b", -1))
	fmt.Println(q.FindAllStringSubmatchIndex("ababab", -1))

	fmt.Println("==========分界线===========")
	r := regexp.MustCompile(`(?m)(key\d+):\s+(value\d+)`)
	content := []byte(`
        # comment line
        key1: value1
        key2: value2
        key3: value3
    `)
	fmt.Println(string(r.Find(content)))
	for _, matched := range r.FindAll(content, -1) {
		fmt.Println(string(matched))
	}
	for _, mutiMatched := range r.FindAllSubmatch(content, -1) {
		for _, matched := range mutiMatched {
			fmt.Println(string(matched))
		}
	}
}
