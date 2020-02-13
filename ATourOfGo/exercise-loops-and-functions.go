//Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	i := 0
	for i = range b {
		b[i] = byte('A')
	}
	fmt.Println(i, len(b), cap(b))
	return i, nil
}

func main() {
	reader.Validate(MyReader{})
}
