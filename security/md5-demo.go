package main

import (
	"github.com/dianjiu/gokit/security"
	"log"
)

func main() {
	md5 := security.MD5("Hello World")
	log.Printf("security.MD5() Result: %v\n", md5)
}
