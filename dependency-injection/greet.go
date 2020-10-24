package dependency_injection

import (
	"bytes"
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(&bytes.Buffer{}, "today")
}