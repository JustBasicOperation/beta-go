package unit

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/gomodule/redigo/redis"
)

type PersonDetail struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 检查用户名是否非法
func checkUsername(username string) bool {
	const pattern = `^[a-z0-9_-]{3,16}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}

// 检查用户邮箱是否非法
func checkEmail(email string) bool {
	const pattern = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 通过 redis 拉取对应用户的资料信息
func getPersonDetailRedis(username string) (*PersonDetail, error) {
	result := &PersonDetail{}

	client, err := redis.Dial("tcp", ":6379")
	// 下面的client.close()方法和client.Do()方法都被内联了，会导致后面运行单测时报错：invalid memory Address or nil point reference
	// 解决办法：运行单测时禁止内联优化:  go test -v -cover -gcflags=all=-l -coverprofile=coverage.out
	// 生成覆盖率报告： go tool cover -html=coverage.out -o coverage.html
	defer client.Close()
	data, err := redis.Bytes(client.Do("GET", username))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 拉取用户资料信息并校验
func GetPersonDetail(username string) (*PersonDetail, error) {
	// 检查用户名是否有效
	if ok := checkUsername(username); !ok {
		return nil, errors.New("invalid username")
	}

	// 从 redis 接口获取信息
	detail, err := getPersonDetailRedis(username)
	if err != nil {
		return nil, err
	}

	// 校验
	if ok := checkEmail(detail.Email); !ok {
		return nil, errors.New("invalid email")
	}

	return detail, nil
}

var b = 10

func max(a int) (res int) {
	return a + b
}

func square(a int) (square int) {
	fun := Fun(a)
	return fun
}

var Fun = func(a int) (square int) {
	return a * a
}
