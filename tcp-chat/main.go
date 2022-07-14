package main

import (
	"flag"

	"github.com/alexparco/tcp-chat/server"
)

var (
	PORT int
)

func init() {
	flag.IntVar(&PORT, "p", 9979, "Set port(Default is 9979)")
	flag.Parse()
}

func main() {
	srv := server.New()

	srv.Run(PORT)
}
