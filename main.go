package main

import (
	"time"

	"github.com/consenbus/core/node"
	"github.com/consenbus/core/store"
)

func main() {
	store.Init(store.LiveConfig)

	keepAliveSender := node.NewAlarm(node.AlarmFn(node.SendKeepAlives), []interface{}{node.PeerList}, 20*time.Second)
	node.ListenForUdp()

	keepAliveSender.Stop()
}
