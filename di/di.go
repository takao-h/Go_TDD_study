package di

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
    fmt.Printf("Hello, %s", name)
}
