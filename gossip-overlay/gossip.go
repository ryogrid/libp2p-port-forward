package gossip_overlay

import (
	"github.com/ryogrid/gossip-overlay/gossip"
	"github.com/weaveworks/mesh"
	"io"
	"log"
)

type Node struct {
	Peer *gossip.Peer
}

func New() (*Node, error) {
	// TODO: not implemented yet

	return nil, nil
}

func (node *Node) OpenStreamToTargetPeer(peerId mesh.PeerName) io.ReadWriteCloser {
	log.Println("Opening a stream to", peerId)

	//passId := peer_.ID
	//stream, err := n.Host.NewStream(ctx, passId, constants.Protocol)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println("Opened a stream to", peer_.ID)

	//return stream
	return nil
}
