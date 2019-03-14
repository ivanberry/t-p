package main

import "golang.org/x/tour/reader"

type MyReader struct {

}

// TODO: Add a Read([]byte) (int, error) method to MyReader
// emits an infinite stream of the ASCII character 'A';
func (r MyReader) Read(buf []byte) (n int, err error)  {
	buf[0] = 'A'
	return 1, nil
}


func main() {
	reader.Validate(MyReader{})
}