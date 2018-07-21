package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
	custom writer interface
*/
type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	// logWriter satisfies Writer interface just by defining the method name and return types

	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*
		bs := make([]byte, 99999) // initialization
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
