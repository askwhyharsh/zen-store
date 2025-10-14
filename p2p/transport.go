package p2p

// Peer is an interface that represents the remote node 
type Peer interface {
}


// Transport handles the comminication 
// between the nodes in the network.
type Transport interface {

	// listen and accept is needed
	ListenAndAccept() error 
}