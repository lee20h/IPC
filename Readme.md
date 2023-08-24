## IPC using Go

- pipe
  - The output of one program is fed into another program through a pipe.
    - `io.Pipe()`
    - `bytes.Buffer`
    - `exec.Cmd.StdoutPipe(), exec.Cmd.StdinPipe()`
