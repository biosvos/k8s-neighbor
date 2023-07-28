package main

import (
	"github.com/biosvos/k8s-neighbor/client"
	"github.com/biosvos/k8s-neighbor/flow"
	"github.com/biosvos/k8s-neighbor/printer"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("filename.d2", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	flo := flow.NewFlow(cli, &printer.D2{
		Writer: file,
	})
	flo.GetWorkloadResources("", "v1", "Pod", "kube-system", "coredns-59b4f5bbd5-kbrvk")
}
