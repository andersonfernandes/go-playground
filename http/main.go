package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

  // body := make([]byte, 99999)
  // resp.Body.Read(body)
  //
	// fmt.Println(string(body))

  io.Copy(os.Stdout, resp.Body)
}
