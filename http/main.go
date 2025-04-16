package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if (err != nil) {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("just wrote %v byte", len(bs))
	return len(bs), nil
}

