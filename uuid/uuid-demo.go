package main

import (
	"github.com/dianjiu/gokit/uuid"
	"log"
)

func main() {
	uid, _ := uuid.NewV4()
	log.Printf("uuid.NewV4() Result: %v\n", uid)
}
