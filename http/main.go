package main

import (
	"os"
	"net/http"
	"io"
)

func main() {
	res, err := http.Get("https://www.google.com")
	
	if err != nil {
		panic(err)
	}


	io.Copy(os.Stdout, res.Body)

}