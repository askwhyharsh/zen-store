package main

import (
	"fmt"
	"log"

	"github.com/askwhyharsh/zen-store/p2p"
)

func main() {
	fmt.Println("let's go")

	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	
	select {}
}