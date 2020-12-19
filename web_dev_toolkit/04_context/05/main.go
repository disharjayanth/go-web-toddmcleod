package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Second)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case ch <- n:
				n++
			case <-ctx.Done():
				fmt.Println("Goroutine returned.")
				return
			}
		}
	}()
	fmt.Println("Gen function returned.")
	return ch
}
