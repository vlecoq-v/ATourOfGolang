//Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func MyReader Read(b []byte) (int,error) {
	i:= 0
	for i = range b; {
		b[i] = byte('A')
	}
	return i, nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}