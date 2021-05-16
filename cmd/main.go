package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/Habibullo-1999/http/pkg/server"

)

const header = "HTTP/1.1 200 OK\r\n" +
	"Content-Length: %s\r\n" +
	"Content-Type: %s\r\n" +
	"Connection: close\r\n" +
	"\r\n"

func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host, port string) (err error) {
	srv := server.NewServer(net.JoinHostPort(host, port))

	srv.Register("/", func(conn net.Conn) {
		body := "Welcome to our website"
		ctl := strconv.Itoa(len(body))
		log.Println(ctl)
		_, err = conn.Write([]byte(fmt.Sprintf(header, ctl, "text/html") + body))
		if err != nil {
			log.Println(err)
		}
	})

	srv.Register("/about", func(conn net.Conn) {
		body := "About Goland Academy"
		ctl := strconv.Itoa(len(body))
		log.Println(ctl)
		_, err = conn.Write([]byte(fmt.Sprintf(header, ctl, "text/html") + body))
		if err != nil {
			log.Println(err)
		}
	})

	return srv.Start()
}
