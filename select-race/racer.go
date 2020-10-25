package select_race

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	select {
	//两个chan的阻塞操作，先写入的那个chan会使得对应的代码先被执行
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	//这里采用struct{}而不是bool型的原因是struct{}占用的空间最小
	//且我们不需要往channel发送任何东西，自然不给channel分配任何空间
	ch := make(chan struct{}) //channel永远用make声明，若用var声明则chan初始化为nil，对nil的chan的任何操作都会导致阻塞
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}


/*
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
*/

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}