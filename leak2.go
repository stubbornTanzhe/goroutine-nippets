package goroutine_nippets

import (
	"fmt"
	"runtime"
	"time"
)
// 接收不发送，指的是channel接收到了值，但是不发送
func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		ch <- struct{}{}
	}()

	time.Sleep(time.Second)
}
