package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var host string
var timeout int
var port int

// 1. connect to host by tcp
// 2. write the input from stdin
// 3. read data from socket and print it on stdout

func init() {
	flag.StringVar(&host, "host", "localhost", "host name")
	flag.IntVar(&timeout, "timeout", 10, "timeout to establish connection in seconds")
	flag.IntVar(&port, "port", 8080, "port number")
}

func read(ctx context.Context, conn net.Conn, sigCh chan os.Signal) {
	scanner := bufio.NewScanner(conn)
ReadLoop:
	for {

		select {
		case <-ctx.Done():
			break ReadLoop
		default:
			if !scanner.Scan() {
				break ReadLoop
			}
			text := scanner.Text()
			fmt.Println("From server: ", text)
		}
	}

	fmt.Println("Finished read")
	sigCh <- syscall.SIGTERM

}

func write(ctx context.Context, conn net.Conn, sigCh chan os.Signal) {
	scanner := bufio.NewScanner(os.Stdin)
WriteLoop:
	for {
		select {
		case <-ctx.Done():
			break WriteLoop
		default:
			if !scanner.Scan() {
				break WriteLoop
			}
			str := scanner.Text()
			fmt.Println("To server: ", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}
	}
	fmt.Println("Finished write")
	sigCh <- syscall.SIGTERM

}

func main() {
	address := host + ":" + strconv.Itoa(port)
	timeout := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// Create a dialer with timeout
	dialer := &net.Dialer{
		Timeout:   timeout,
		KeepAlive: 0,
	}

	conn, err := dialer.DialContext(ctx, "tcp", address)

	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	defer conn.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("Sending cancel signal to all go-routines")
		cancel()
	}()

	go write(ctx, conn, sigCh)
	go read(ctx, conn, sigCh)

	<-ctx.Done()
}
