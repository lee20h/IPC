package main

import (
	util "github.com/lee20h/IPC"
	"github.com/lee20h/IPC/named_pipe/ipc"
	"os"
	"os/exec"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	namedPipe := ipc.CreateNamedPipe()

	waitGroup.Add(1)

	go execute("./write", namedPipe)

	go ipc.Read(namedPipe)

	waitGroup.Wait()

	if err := os.Remove(namedPipe); err != nil {
		util.ExecutionError("error deleting tempDir", err)
	}

}

func execute(executablePath, namedPipe string) {
	defer waitGroup.Done()
	cmd := exec.Command(executablePath, namedPipe)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		util.ExecutionError("error during write process", err)
	}
}
