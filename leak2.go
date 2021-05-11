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
		ch <- struct{}{} // 这里给一个nil的channel发送数据，会永远阻塞
	}()
	//<- ch     // 如果这里是从一个空的channel里读取数据，会导致panic

	//go func() {   // 这里如果是启动一个goroutine来获取channel里的值，不会panic，会阻塞
	//	<- ch
	//}()

	//Given a nil channel c:
	//<-c receiving from c blocks forever
	//c <- v sending into c blocks forever
	//close(c) closing c panics

	time.Sleep(time.Second)
}
