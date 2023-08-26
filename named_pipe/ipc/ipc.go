package ipc

import (
	"fmt"
	util "github.com/lee20h/IPC"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

func CreateNamedPipe() string {
	tmpDir, err := os.MkdirTemp("", "named-pipes")

	namedPipe := filepath.Join(tmpDir, "stdout")

	if err != nil && !os.IsNotExist(err) {
		util.ExecutionError("error removing named pipe", err)
	}

	if err := syscall.Mkfifo(namedPipe, 0666); err != nil {
		util.ExecutionError("error creating named pipe", err)
	}

	fmt.Println("Named pipe created:", namedPipe)

	return namedPipe
}

func Read(namedPipe string) {
	file, err := os.OpenFile(namedPipe, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		util.ExecutionError("error opening named pipe", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		util.ExecutionError("error reading from named pipe", err)
	}

	fmt.Println("Received message:", string(content))
}

func Write(namedPipe string, content string) {
	file, err := os.OpenFile(namedPipe, os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		util.ExecutionError("error opening named pipe", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		util.ExecutionError("error writing from named pipe", err)
	}
}
