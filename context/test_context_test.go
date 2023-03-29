package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 模拟rpc框架超时控制
func Test_01(t *testing.T) {
	start := time.Now()
	pCtx := context.Background()
	ctx, cancelFunc := context.WithTimeout(pCtx, 2*time.Second)
	go func(ctx context.Context) {
		fmt.Printf("start sleep\n")
		time.Sleep(5 * time.Second)
		fmt.Printf("end sleep\n")
	}(ctx)
	select {
	case <-ctx.Done():
		fmt.Printf("receive ctx signal\n")
	}
	cancelFunc()
	time.Sleep(3 * time.Second)
	fmt.Printf("cost: %v\n", time.Since(start))
}

// 模拟手动cancel，结束外层协程
func Test_02(t *testing.T) {
	start := time.Now()
	pCtx := context.Background()
	ctx, cancelFunc := context.WithCancel(pCtx)
	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		fmt.Printf("start cancelFunc\n")
		cancelFunc()

		time.Sleep(5 * time.Second)
		fmt.Printf("end sleep\n")
	}(ctx, cancelFunc)
	select {
	case <-ctx.Done():
		fmt.Printf("receive ctx signal\n")
	}

	time.Sleep(8 * time.Second)
	fmt.Printf("cost: %v\n", time.Since(start))
}

type emptyCtx int

func Test_03(t *testing.T) {
	a := new(emptyCtx)
	b := new(emptyCtx)
	fmt.Printf("%p, %p\n", a, b)
}

func Test_testWithCancel01(t *testing.T) {
	testWithCancel01()
}

func Test_testWithCancel02(t *testing.T) {
	testWithCancel02()
}

func Test_testWithDeadline(t *testing.T) {
	testWithDeadline()
}

func Test_testWithValue(t *testing.T) {
	testWithValue()
}
