package main

import (
	"flag"
	"log"

	"github.com/marcusbello/paydistro/internal/aedc/server"
)

var (
	aedcAddr = flag.String("addr", "127.0.0.1:6675", "The address to run on.")
)

func main() {
	flag.Parse()
	s, err := server.New(*aedcAddr)
	if err != nil {
		panic(err)
	}
	done := make(chan error, 1)

	log.Println("Starting server at: ", *aedcAddr)
	go func() {
		defer close(done)
		done <- s.Start()
	}()

	err = <-done
	log.Println("Server exited with error: ", err)
}
