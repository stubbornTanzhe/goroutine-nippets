package goroutine_nippets

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)


func main() {
	for i := 0; i < 2; i++ {
		queryAll()
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}

// 其重点在于调用 query 方法后的结果会写入 ch 变量中，接收成功后再返回 ch 变量。
// 最后可看到输出的 goroutines 数量是在不断增加的，每次多 2 个。也就是每调用一次，都会泄露 Goroutine。
// 原因在于 channel 均已经发送了（每次发送 3 个），但是在接收端并没有接收完全（只返回 1 个 ch），所诱发的 Goroutine 泄露。
func queryAll() int {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func() {
			ch <- query()
		}()
		fmt.Printf("queryAll goroutines: %d\n", runtime.NumGoroutine())
	}

	return <-ch
}

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

