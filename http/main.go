package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("Error : ", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
