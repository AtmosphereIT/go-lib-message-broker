package main

import (
	"github.com/AtmosphereIT/go-lib-message-broker/broker"
	channel2 "github.com/AtmosphereIT/go-lib-message-broker/channel"
)

func main() {
	// Create a broker
	//
	b := broker.MakeBroker()

	// Create a channel
	//
	c := channel2.MakeChannel()

	// Add the channel to the broker
	//
	b.AddChannel("test", c)

	//
	//
}
