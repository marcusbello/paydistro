package main

import (
	"flag"
	"log"

	"github.com/marcusbello/paydistro/internal/wallet/server"
)

var (
	walletAddr = flag.String("addr", "127.0.0.1:6665", "The address to run on.")
)

func main() {
	flag.Parse()
	s, err := server.New(*walletAddr)
	if err != nil {
		panic(err)
	}
	done := make(chan error, 1)

	log.Println("Starting server at: ", *walletAddr)
	go func() {
		defer close(done)
		done <- s.Start()
	}()

	err = <-done
	log.Println("Server exited with error: ", err)
}
