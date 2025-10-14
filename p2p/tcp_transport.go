package p2p

import (
	"fmt"
	"log/slog"
	"net"
	"sync"
)

// tcp peer represents the remote node over4 a TCP established connection
type TCPPeer struct {
	conn net.Conn
	// if we dial and connect (outbound) == true
	// but if we accept a connection => then it's outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn: 		conn,
		outbound: 	outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener net.Listener

	mu sync.Mutex // mutex for the Peers 
	peers map[net.Addr]Peer
}


func NewTCPTransport(listenrAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenrAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err !=  nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

// this is private method as starts with a (just a note for myself)
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			slog.Error("TCP accept error: %s\n", err)
		}
		// connection handler
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("new incoming connection %+v\n", peer)
}