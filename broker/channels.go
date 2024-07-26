package broker

import (
	"errors"
	"github.com/AtmosphereIT/go-lib-message-broker/channel"
)

// GetChannel - Get a channel from the broker by name.
func (b *Broker) GetChannel(name string) *channel.Channel {
	if b.Channels[name] == nil {
		return nil
	}
	return b.Channels[name]
}

// AddChannel - Add a channel to the broker.
func (b *Broker) AddChannel(name string, channel *channel.Channel) error {
	if b.GetChannel(name) != nil {
		return errors.New(ERROR_CHANNEL_NAME_NOT_UNIQUE)
	}

	b.Channels[name] = channel
	return nil
}
