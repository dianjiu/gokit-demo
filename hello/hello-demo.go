package main

import (
	"github.com/dianjiu/gokit/hello"
	"log"
)

func main() {
	hi := hello.SayHello()
	log.Printf("hello.SayHello() Result: %v\n", hi)
}
