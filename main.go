package main

import (
	"fmt"
	"log"
)

func main() {
	storage, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", storage)
	server.Run()
	fmt.Println("hello")
}
