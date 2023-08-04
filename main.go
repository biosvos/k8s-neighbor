package main

import (
	"fmt"
	"github.com/biosvos/k8s-neighbor/client"
	"github.com/biosvos/k8s-neighbor/dresource"
	"github.com/biosvos/k8s-neighbor/flow"
	"github.com/biosvos/k8s-neighbor/printer"
	"log"
	"os"
	"strings"
)

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println(file)
	}
}

func setupLog() func() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logFile := openFile("log.log")
	log.SetOutput(logFile)
	return func() {
		closeFile(logFile)
	}
}

func openFile(filename string) *os.File {
	ret, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	return ret
}

func main() {
	closeFn := setupLog()
	defer closeFn()
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	list, err := cli.List()
	if err != nil {
		panic(err)
	}
	for _, resource := range list {
		one(resource, cli)
	}
}

func one(resource []byte, cli *client.Client) {
	newResource, err := dresource.NewResource[dresource.Normal](resource)
	if err != nil {
		panic(err)
	}
	group, version := parseAPIVersion(newResource.APIVersion)
	filename := fmt.Sprintf("all/%v_%v_%v_%v_%v.d2", group, version, newResource.Kind, newResource.Metadata.Namespace, newResource.Metadata.Name)
	file := openFile(filename)
	defer closeFile(file)

	flo := flow.NewFlow(cli, &printer.D2{
		Writer: file,
	})
	flo.GetWorkloadResources(group, version, newResource.Kind, newResource.Metadata.Namespace, newResource.Metadata.Name)
}

func parseAPIVersion(apiVersion string) (string, string) {
	split := strings.Split(apiVersion, "/")
	if len(split) == 1 {
		return "", split[0]
	}
	return split[0], split[1]
}
