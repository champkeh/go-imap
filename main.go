package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/champkeh/go-imap/internal/tag"
)

var (
	// imap server address
	// qq mail server: imap.qq.com:993(ssl)
	//				   imap.qq.com:143(no ssl)
	addr string

	// tag generator
	// generate tag from A0001, A0002, A0003,...
	tagGenerator func() string
)

func init() {
	tagGenerator = tag.NewTagGenerator()
}

func main() {
	flag.StringVar(&addr, "addr", "", "imap server address")
	flag.Parse()

	if len(addr) == 0 {
		log.Println("Please input imap server address with -addr flag")
		os.Exit(1)
	}

	fmt.Printf("connecting to %s\n", addr)
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	defer conn.Close()

	// 开启read-goroutine
	go func() {
		for {
			buf := make([]byte, 512)
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\rS[%d]:%s", n, buf[:n])
			fmt.Printf("imap>")
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = tagGenerator() + " " + strings.Trim(text, "\r\n") + "\r\n"
		send(conn, text)
	}
}

func send(conn *tls.Conn, cmd string) {
	fmt.Printf("C[%d]:%s", len(cmd), cmd)
	conn.Write([]byte(cmd))
}
