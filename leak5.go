package goroutine_nippets

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func handle(v int) {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < v; i++ {
		fmt.Println("aaa")
		wg.Done()
	}
	wg.Wait()
}

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	go handle(3)
	time.Sleep(time.Second)
}
