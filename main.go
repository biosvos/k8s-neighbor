package main

import (
	"github.com/biosvos/k8s-neighbor/internal"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := internal.Work()
	log.Println(err)
}
