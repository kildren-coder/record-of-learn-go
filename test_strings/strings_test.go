package test_strings

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	res := strings.Compare("a", "b")
	want := -1


	if res != want {
		t.Errorf("excepted '%q' but get '%q'", want, res)
	}

	res = strings.Compare("a", "a")
	want = 0

	if res != want {
		t.Errorf("excepted '%q' but get '%q'", want, res)
	}

	res = strings.Compare("b", "a")
	want = 1
	if res != want {
		t.Errorf("excepted '%q' but get '%q'", want, res)
	}
}

func TestCount(t *testing.T)  {
	res := strings.Count("cheese eee", "ee")
	want := 2

	if res != want {
		t.Errorf("excepted '%q' but get '%q'", want, res)
	}
}


func Benchmark_Compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Compare("a","b")
		strings.Compare("a","a")
		strings.Compare("b","a")
	}
}

func Benchmark_compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compare("a", "b")
		compare("a", "a")
		compare("b", "a")
	}
}
