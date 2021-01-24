//ex25/ex25.9/ex25.9.go
package main

import (
	"context"
	"fmt"
)

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil { // ❷
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "number", 9) // ❶
	square(ctx)
}
