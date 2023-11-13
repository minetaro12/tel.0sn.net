package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	listen := fmt.Sprintf(":%s", getEnv("PORT", "8023"))
	tcpAddr, err := net.ResolveTCPAddr("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	count := loadCounter()
	log.Println("Listening:", listener.Addr().String())
	go acceptTCP(listener, &count)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down...")
	saveCounter(count)
}

func acceptTCP(listener *net.TCPListener, count *int) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		go echoHandler(conn, count)
	}
}

func echoHandler(conn *net.TCPConn, count *int) {
	log.Println("Connected:", conn.RemoteAddr().String())
	*count++
	text := `
--------------------------
tel.0sn.netへようこそ！
_       _   ___                         _   
| |_ ___| | / _ \ ___ _ __    _ __   ___| |_ 
| __/ _ \ || | | / __| '_ \  | '_ \ / _ \ __|
| ||  __/ || |_| \__ \ | | |_| | | |  __/ |_ 
 \__\___|_(_)___/|___/_| |_(_)_| |_|\___|\__|
--------------------------
あなたは %d 人目の訪問者です。
--------------------------
Web: https://0sn.net
--------------------------
`
	for _, v := range fmt.Sprintf(text, *count) {
		io.WriteString(conn, string(v))
		time.Sleep(50 * time.Millisecond)
	}

	conn.CloseWrite()
}

func loadCounter() int {
	if _, err := os.Stat("counter.txt"); os.IsNotExist(err) {
		os.WriteFile("counter.txt", []byte("0"), 0644)
	}

	data, err := os.ReadFile("counter.txt")
	if err != nil {
		log.Fatal(err)
	}

	c, err := strconv.Atoi(string(data))
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func saveCounter(counter int) {
	err := os.WriteFile("counter.txt", []byte(strconv.Itoa(counter)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if v, isFound := os.LookupEnv(key); isFound {
		return v
	} else {
		return fallback
	}
}
