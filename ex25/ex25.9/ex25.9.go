//ex25/ex25.9/ex25.9.go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "number", 9) // ❶ 컨텍스트에 값을 추가한다
	square(ctx)
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil { // ❷ 컨텍스트에서 값을 읽어온다.
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
}
