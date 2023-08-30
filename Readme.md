## IPC using Go

- pipe
  - The output of one program is fed into another program through a pipe.
    - `io.Pipe()`
    - `bytes.Buffer`
    - `exec.Cmd.StdoutPipe(), exec.Cmd.StdinPipe()`
- named pipe (FIFO)
  - The process responsible for reading and the process responsible for writing to one pipe connects and communicates.
    - `syscall.Mkfifo()`
    - `type FileMode = fs.FileMode`
      - `ModeNamedPipe  = fs.ModeNamedPipe  // p: named pipe (FIFO)`
     