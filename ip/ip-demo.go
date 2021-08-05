package main

import (
	"github.com/dianjiu/gokit/ip"
	"log"
)

func main() {
	ip := ip.GetLocalIPv4()
	log.Printf("本机IP: %v\n", ip)
}
