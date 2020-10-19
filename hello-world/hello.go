package main

import "fmt"

//敞亮可提高应用程序的性能，它避免了每次使用Hello时创建“Hello， ”字符串
//创建常亮的价值还在雨可以快速理解值的含义
const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}


	return greetinfPrefix(language) + name
}

func greetinfPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("world",""))
}

