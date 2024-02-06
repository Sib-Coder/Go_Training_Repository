package main

import (
	"io"
	"log/slog"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "148.251.26.213:80")
	if err != nil {
		slog.Error("Unable to connect to our unreachable host")
	}
	defer dst.Close()

	//от получателя к источнику и назад
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
		}
		slog.Error("error is :", err)
	}()
	if _, err := io.Copy(src, dst); err != nil {
		slog.Error("error is :", err)
	}

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
