// A QOTD (RFC 865) server.  Only implements TCP, not UDP.
package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

var quotes = []string{
	`In film you will find four basic story lines. Man versus man, man
versus nature, nature versus nature, and dog versus vampire.
    - Steven Spielberg
`,
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func listen(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("%s: %s", l.Addr(), err)
			continue
		}

		q := quotes[rand.Intn(len(quotes))]
		io.WriteString(c, q)
		c.Close()
	}
}

func main() {
	var port string
	flag.StringVar(&port, "port", "1717", "TCP port to listen on (default 1717)")
	flag.Parse()

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("tcp: %s", err)
	}

	log.Printf("QOTD listening on port %s", port)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("%s: %s", l.Addr(), err)
			continue
		}

		q := quotes[rand.Intn(len(quotes))]
		io.WriteString(c, q)
		c.Close()
	}
}
