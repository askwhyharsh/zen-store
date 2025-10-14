package p2p

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"sync"
)


type TCPTransportOpts struct {
	ListenAddr string
	HandShakeFunc HandShakeFunc
	Decoder Decoder
}
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
	TCPTransportOpts
	listener net.Listener

	mu sync.Mutex // mutex for the Peers 
	peers map[net.Addr]Peer
}


func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
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


	// first we do NOP handshake, if that fails then we don't proceed with decoding
	if err := t.HandShakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	//  close connection if lots of errors, by len of error string

	const maxConsecutiveErrors = 5
	consecutiveErrors := 0

	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			// close immediately on EOF
			if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
				_ = conn.Close()
				return
			}
			consecutiveErrors++
			fmt.Printf("TCP decode error (%d/%d): %s\n", consecutiveErrors, maxConsecutiveErrors, err)
			if consecutiveErrors >= maxConsecutiveErrors {
				_ = conn.Close()
				return
			}
			continue
		}
		msg.From = conn.RemoteAddr()
		// successful decode: reset error counter
		consecutiveErrors = 0
		fmt.Printf("mesaage: %+v\n", msg)
	}
}