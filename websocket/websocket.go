package websocket

import (
	"github.com/TankerHQ/go-engine.io/transport"
)

var Creater = transport.Creater{
	Name:      "websocket",
	Upgrading: true,
	Server:    NewServer,
	Client:    NewClient,
}
