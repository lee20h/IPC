package main

import (
	"fmt"
	util "github.com/lee20h/IPC"
	"github.com/lee20h/IPC/named_pipe/ipc"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		util.InvalidArg("Usage: ./write named_pipe")
	}
	namedPipe := os.Args[1]
	fmt.Println("Opening named pipe for writing")
	ipc.Write(namedPipe, "inter process communication")
}
