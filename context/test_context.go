package context

import (
	"context"
	"fmt"
	"time"
)

// 用法1：使用WithCancel函数，当主协程退出时，控制子协程也退出
func testWithCancel01() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		monitorCancelSignal(ctx)
	}()
	// 一定要手动调用cancel函数来取消context，否则这个context永远都不会被取消
	cancel()
	fmt.Printf("main goroutine end\n")
	time.Sleep(3 * time.Second) // sleep是为了观察子协程的执行结果
}

func monitorCancelSignal(ctx context.Context) {
	time.Sleep(1 * time.Second)
	// 方法1：如果ctx被取消，退出协程
	if ctx.Err() != nil {
		fmt.Printf("canceled: %v\n", ctx.Err())
		return
	}
	// 方法2：也可以用ctx.Done()方法来监听取消信号，这里是阻塞式监听，也可以用for+select的非阻塞式监听
	//select {
	//case a := <-ctx.Done(): // channel关闭，读取到零值
	//	fmt.Printf("receive cancel signal: %v\n", a)
	//}
}

// 用法2：使用WithCancel函数，当子协程退出时，主协程才能退出
func testWithCancel02() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		controlMainRoutine(ctx, cancel)
	}()

	// 阻塞监听
	select {
	case <-ctx.Done():
		fmt.Printf("receive sub routine cancel signal\n")
	}

	fmt.Printf("main goroutine end\n")
}

func controlMainRoutine(ctx context.Context, cancel context.CancelFunc) {
	time.Sleep(1 * time.Second)
	fmt.Printf("sub routine end\n")
	cancel()
}

// 使用WithDeadline()函数，控制子协程超时退出或者手动退出
func testWithDeadline() {
	rootCtx := context.Background()
	deadline := time.Now().Add(1 * time.Second)
	ctx1, cancel1 := context.WithDeadline(rootCtx, deadline)
	ctx2, cancel2 := context.WithDeadline(rootCtx, deadline)
	go func() {
		monitorTimeoutCancel(ctx1) // 超时退出
		monitorManualCancel(ctx2)  // cancel()方法退出
	}()
	cancel2()
	fmt.Printf("main goroutine end\n")
	time.Sleep(3 * time.Second)
	cancel1() // 放到最后是为了
}

func monitorTimeoutCancel(ctx context.Context) {
	// 方法1：非阻塞，只判断一次
	time.Sleep(1 * time.Second)
	deadline, ok := ctx.Deadline()
	if ok && time.Now().After(deadline) {
		fmt.Printf("method1: receive cancel signal of timeout\n")
	}

	// 方法2：这里是阻塞式监听
	//select {
	//case <-ctx.Done():
	//	fmt.Printf("method2: receive cancel signal of timeout\n")
	//}

	// 方法3：这里是非阻塞式监听，每隔1s监听一次
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Printf("method3: receive cancel signal of timeout\n")
	//	default:
	//		time.Sleep(1 * time.Second)
	//		fmt.Printf("method3: not received, continue")
	//	}
	//}

	// 方法4：非阻塞，只判断一次
	//select {
	//case <-ctx.Done():
	//	fmt.Printf("method4: receive cancel signal of timeout\n")
	//default:
	//	fmt.Printf("method4: ignore cancel signal, continue\n")
	//}

	// 一般在执行某个函数前判断下当前ctx是否已经被取消，提前返回不浪费资源
	// do something...
}

func monitorManualCancel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Printf("receive cancel signal of timeout\n")
	}
}

func testWithValue() {
	rootCtx := context.Background()
	ctx := context.WithValue(rootCtx, "key", "value")
	subCtx := context.WithValue(ctx, "key1", "value1")
	fmt.Printf("get value form ctx: %v\n", ctx.Value("key").(string))
	fmt.Printf("get value from subCtx: %v\n", subCtx.Value("key").(string)) // 会向上回溯直到找到或者返回nil
}
