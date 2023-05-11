package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world")
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("can not load env:port")
	}
}
