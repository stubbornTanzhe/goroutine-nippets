package goroutine_nippets

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan int
	go func() {
		<-ch
	}()

	time.Sleep(time.Second)
}
