package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lw := logWriter{}

	//_, _ = io.Copy(os.Stdout, resp.Body)
	_, _ = io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	 fmt.Println(string(bs))
	 fmt.Println("Wrote", len(bs), "bytes")
	 return len(bs), nil
}
