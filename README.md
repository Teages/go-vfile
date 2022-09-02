# go-vfile

A virtual file system for go in runtime. 

## Example

Easy

```go
package main

import (
	"embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Teages/go-vfile"
)

//go:embed files/*
var f embed.FS

func run() {
	fmt.Println(vfile.GetPath("."))
	err := vfile.JoinAll(".", f)
	fmt.Println(err)
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		vfile.Close()
		os.Exit(0)
	}()
	run()
	vfile.Close()
}

```