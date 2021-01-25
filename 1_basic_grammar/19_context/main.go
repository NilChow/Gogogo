package main

import (
	"context"
	"fmt"
)

func GetContextWithValue() context.Context {
	ctx := context.WithValue(context.Background(), "Test-Key", 55555)
	return ctx
}

func ParseContextValue(ctx context.Context) {
	val := ctx.Value("Test-Key")
	if num, ok := val.(int); ok {
		fmt.Printf("[Parse context Success]: [%d]", num)
	} else {
		fmt.Printf("[Parse context Failed]")
	}
}

func main() {
	fmt.Println("Context Test Start...")

	ctx := GetContextWithValue()
	ParseContextValue(ctx)
}