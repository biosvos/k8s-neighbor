package main

import (
	"github.com/biosvos/k8s-neighbor/client"
	"github.com/biosvos/k8s-neighbor/flow"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	flo := flow.NewFlow(cli)
	flo.GetWorkloadResources("", "v1", "Pod", "kube-system", "coredns-59b4f5bbd5-kbrvk")
}
