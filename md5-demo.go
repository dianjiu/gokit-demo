package main

import (
	"github.com/dianjiu/gokit/security"
	"log"
)

func main() {
	md5 := security.MD5("Hello World")
	log.Printf("MD5加密后: %v\n", md5)
}
