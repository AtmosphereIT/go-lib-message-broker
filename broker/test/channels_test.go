package broker_test

import (
	"github.com/AtmosphereIT/go-lib-message-broker/broker"
	"github.com/AtmosphereIT/go-lib-message-broker/channel"
	"testing"
)

// makeTestChannel - Create a new channel instance for testing.
func makeTestChannel() *channel.Channel {
	return channel.MakeChannel("test", nil)
}

// TestAddChannel - Test the AddChannel method.
func TestAddChannel(t *testing.T) {
	brokerInstance := broker.MakeBroker()
	channel := makeTestChannel()

	brokerInstance.AddChannel("test", channel)

	if brokerInstance.Channels["test"] == nil {
		t.Errorf("Channel not added successfully.")
	}

	if brokerInstance.Channels["test"] != channel {
		t.Errorf("Added channel does not match the expected channel.")
		t.Failed()
	}

	// Channel was added successfully.
	t.Logf("Channel added successfully.")
}

// TestAddChannelReturnsErrorWhenNameAlreadyExists - Test the AddChannel method returns an error when the channel name already exists.
func TestAddChannelReturnsErrorWhenNameAlreadyExists(t *testing.T) {
	brokerInstance := broker.MakeBroker()
	channel := makeTestChannel()

	brokerInstance.AddChannel("test", channel)

	err := brokerInstance.AddChannel("test", channel)

	if err == nil {
		t.Errorf("No error returned when adding a channel with a name that already exists.")
		t.Failed()
	}

	if err.Error() != broker.ERROR_CHANNEL_NAME_NOT_UNIQUE {
		t.Errorf("Received an error but not the expected error. Got: %s", err.Error())
		t.Failed()
	}
}

// TestGetChannelReturnsNilWhenNoChannelExistsWithGivenName - Test the GetChannel method returns nil when no channel exists with the given name.
func TestGetChannelReturnsNilWhenNoChannelExistsWithGivenName(t *testing.T) {
	brokerInstance := broker.MakeBroker()

	channel := brokerInstance.GetChannel("test")

	if channel != nil {
		t.Errorf("Expected nil channel but got a channel.")
		t.Failed()
	}

	// Channel was returned successfully.
	t.Logf("Channel returned successfully.")
}

// TestGetChannelReturnsChannelWhenChannelExistsWithGivenName - Test the GetChannel method returns the channel when a channel exists with the given name.
func TestGetChannelReturnsChannelWhenChannelExistsWithGivenName(t *testing.T) {
	brokerInstance := broker.MakeBroker()
	channel := makeTestChannel()

	brokerInstance.AddChannel("test", channel)

	channel = brokerInstance.GetChannel("test")

	if channel == nil {
		t.Errorf("Expected channel but got nil.")
		t.Failed()
	}

	// Channel was returned successfully.
	t.Logf("Channel returned successfully.")
}
