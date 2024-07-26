package broker

import "github.com/AtmosphereIT/go-lib-message-broker/channel"

// MakeBroker - Create a new broker instance.
func MakeBroker() *Broker {
	return &Broker{
		Channels: make(map[string]*channel.Channel),
	}
}
