package channel

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_channel01(t *testing.T) {
	ch := make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for c := range ch {
				fmt.Printf("get int: %v\n", c)
			}
			wg.Done()
		}()
	}
	//time.Sleep(3*time.Second)
	wg.Wait()
}

func Test_channel02(t *testing.T) {
	go func() {
		go produce()
		go consumer()
	}()
	time.Sleep(10 * time.Second)
}

// GenBaseHeader 生成基础 Header
func GenBaseHeader(privateToken, paasID, paasToken string) map[string]string {
	// TOF 鉴权签名处理
	timestamp := fmt.Sprintf("%d", time.Now().Unix()) // 生成时间戳，注意服务器的时间与标准时间差不能大于180秒
	rand.Seed(time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().Unix()))
	nonce := strconv.Itoa(r.Intn(4096)) // 随机字符串，十分钟内不重复即可
	signStr := fmt.Sprintf("%s%s%s%s", timestamp, paasToken, nonce, timestamp)
	sign := fmt.Sprintf("%X", sha256.Sum256([]byte(signStr)))
	header := make(map[string]string)
	header["PRIVATE-TOKEN"] = privateToken
	header["x-rio-paasid"] = paasID
	header["x-rio-nonce"] = nonce
	header["x-rio-timestamp"] = timestamp
	header["x-rio-signature"] = sign
	fmt.Println(nonce)
	fmt.Println(timestamp)
	fmt.Println(sign)
	return header
}
