package main

import (
	"fmt"
	"log"

	"github.com/askwhyharsh/zen-store/p2p"
)

func main() {
	fmt.Println("let's go")
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HandShakeFunc: p2p.NOPHandShakeFunc,
		Decoder: p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}