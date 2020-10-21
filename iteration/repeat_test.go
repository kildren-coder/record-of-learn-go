package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {

	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but get '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {

	var times int
	fmt.Scan(&times)

	for i := 0; i < b.N; i++ {
		Repeat("a", times)
	}
}

func ExampleRepeat() {
	var temp string
	temp = Repeat("b", 6)
	fmt.Println(temp)
	// Output: bbbbbb
}
