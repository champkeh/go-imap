package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"bytes"

	"github.com/champkeh/go-imap/internal/tag"
	"github.com/fatih/color"
)

var (
	// imap server address
	// qq mail server: imap.qq.com:993(ssl)
	//				   imap.qq.com:143(no ssl)
	addr string

	// tag generator
	// generate tag from A0001, A0002, A0003,...
	tagGenerator func() string

	sColor *color.Color
	cColor *color.Color
)

func init() {
	tagGenerator = tag.NewTagGenerator()

	cColor = color.New(color.FgMagenta)
	sColor = color.New(color.FgBlue)
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
		scanner := bufio.NewScanner(conn)
		scanner.Split(ScanCRLF)
		isGreeting := true
		for scanner.Scan() {
			resp := scanner.Text()
			sColor.Printf("S[%d]:", len(resp))
			fmt.Printf("%s\n", resp)
			if resp[0] != '*' {
				fmt.Printf("imap>")
			}
			if isGreeting {
				fmt.Printf("imap>")
				isGreeting = false
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
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
	cColor.Printf("C[%d]:", len(cmd))
	fmt.Printf("%s", cmd)
	conn.Write([]byte(cmd))
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}
