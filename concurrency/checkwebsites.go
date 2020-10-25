package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		//对于每次迭代来说，url指向的内存地址都是一样的
		//而在例子里，主线程的for跑完了后协程才开始跑里面的语句，导致大家的url都指向最后一个url了
		//解决方法：为每个url创建副本
		go func(u string) {
			results[u] = wc(u)
		}(url) //这个()使得该匿名函数在声明的同时执行
	}

	time.Sleep(2 * time.Second)

	return results
}


