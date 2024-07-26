package broker

import "github.com/AtmosphereIT/go-lib-message-broker/channel"

type Broker struct {
	// Channels - A map of channels.
	Channels map[string]*channel.Channel
}
