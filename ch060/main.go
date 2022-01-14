package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println(runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancelF := context.WithCancel(ctx)

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println("running",n)
			for {
				select {
				case <- ctx.Done():
					return
				default:
					fmt.Println("still working", n)
					time.Sleep(100 * time.Millisecond)
				}
				fmt.Println("GORUTINES RUNNING:", runtime.NumGoroutine())
			}
		}(i)
	}
	time.Sleep(5 * time.Second)
	cancelF()
	fmt.Println("Sleeping fot 5 seconds")
	time.Sleep(time.Second)
	fmt.Println("GOROUTINES RUNNING:", runtime.NumGoroutine())

}