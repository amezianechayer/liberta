package node

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/amezianechayer/liberta/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version  string
	peerLock sync.RWMutex
	peers    map[proto.NodeClient]bool
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		peers:   make(map[proto.NodeClient]bool),
		version: "liberta-0.1",
	}

}

func (n *Node) addPeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	n.peers[c] = true
}

func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	delete(n.peers, c)
}

func (n *Node) Start(listenAddr string) error {

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	proto.RegisterNodeServer(grpcServer, n)

	fmt.Println("node running on port:", ":3000")
	return grpcServer.Serve(ln)
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}
	p, _ := peer.FromContext(ctx)
	//c, err := makeNodeClient()

	fmt.Printf("received version from %s: %+v\n", v, p.Addr)

	return ourVersion, nil

}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("received tx from:", peer)
	return &proto.Ack{}, nil
}

func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	c, err := grpc.Dial(listenAddr)
	if err != nil {
		return nil, err
	}
	return proto.NewNodeClient(c), nil
}
