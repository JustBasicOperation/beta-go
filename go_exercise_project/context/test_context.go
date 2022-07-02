package context

import (
	"context"
	"fmt"
)

func main() {
	oriCtx := context.Background()
	_ = context.WithValue(oriCtx, "key1", "value1")
	_ = context.WithValue(oriCtx, "key2", "value2")
	value := oriCtx.Value("key1")
	fmt.Println(value.(string))
}
