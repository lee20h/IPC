package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		invalidArg("Usage: ./pipe '<command1> | <command2>'")
	}

	commands := strings.Split(os.Args[1], "|")
	if len(commands) < 2 {
		invalidArg("You need to provide two commands separated by '|'")
	}

	command1 := strings.Fields(commands[0])
	command2 := strings.Fields(commands[1])

	cmd1 := exec.Command(command1[0], command1[1:]...)
	cmd2 := exec.Command(command2[0], command2[1:]...)

	cmd1Out, err := cmd1.StdoutPipe()
	if err != nil {
		executionError("Error creating stdout pipe for first program", err)
	}

	cmd2In, err := cmd2.StdinPipe()
	if err != nil {
		executionError("Error creating stdout pipe for second program", err)
	}

	cmd2.Stdout = os.Stdout

	go func() {
		defer cmd2In.Close()
		io.Copy(cmd2In, cmd1Out)
	}()

	if err := cmd1.Start(); err != nil {
		executionError("Error starting first program", err)
	}
	if err := cmd2.Start(); err != nil {
		executionError("Error starting second program", err)
	}

	cmd1.Wait()
	if err := cmd2.Wait(); err != nil {
		executionError("Error running second program", err)
	}
}

func invalidArg(message string) {
	fmt.Fprintln(os.Stdout, message)
	os.Exit(1)
}

func executionError(message string, err error) {
	fmt.Fprintln(os.Stderr, message, err)
	os.Exit(1)
}

func bufferMethod(command1, command2 []string) {
	cmd1 := exec.Command(command1[0], command1[1:]...)
	var outputBuf bytes.Buffer
	cmd1.Stdout = &outputBuf
	if err := cmd1.Run(); err != nil {
		executionError("Error running first program", err)
	}

	cmd2 := exec.Command(command2[0], command2[1:]...)

	cmd2.Stdin = &outputBuf
	cmd2.Stdout = os.Stdout

	if err := cmd2.Run(); err != nil {
		executionError("Error running second program", err)
	}
}

func pipeMethod(command1, command2 []string) {
	cmd1 := exec.Command(command1[0], command1[1:]...)
	cmd2 := exec.Command(command2[0], command2[1:]...)

	r, w := io.Pipe()
	cmd1.Stdout = w
	cmd2.Stdin = r

	cmd2.Stdout = os.Stdout

	if err := cmd1.Start(); err != nil {
		executionError("Error starting first program", err)
	}

	if err := cmd2.Start(); err != nil {
		executionError("Error starting second program", err)
	}

	go func() {
		defer w.Close()
		cmd1.Wait()
	}()

	if err := cmd2.Wait(); err != nil {
		fmt.Println("Error running second program", err)
		return
	}
}
