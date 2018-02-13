package main

import (
	"github.com/consenbus/core/node"
	"github.com/consenbus/core/store"
)

func main() {
	store.Init(store.LiveConfig)

	node.SendKeepAlive(node.PeerList[0])
	node.ListenForUdp()
}
