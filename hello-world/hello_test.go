package main

import "testing"

/*
重构不 仅仅 是针对程序的代码！
重要的是，你的测试 清楚地说明 了代码需要做什么。
 */
func TestHello(t *testing.T)  {

	//我们将断言重构为函数。这减少了重复并且提高了测试的可读性。
	assertCorrectMessage := func(t *testing.T, got, want string) {
		//t.Helper()将错误定位到具体的行号
		//如果不加，则错误显示为17行
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris","")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jack", "French")
		want := "Bonjour, Jack"
		assertCorrectMessage(t, got, want)
	})
}
