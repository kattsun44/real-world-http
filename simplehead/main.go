package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Fields:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
