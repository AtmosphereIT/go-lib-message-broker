package broker_test

import (
	"github.com/AtmosphereIT/go-lib-message-broker/broker"
	"testing"
)

// MakeBrokerTest - Create a new broker instance for testing.
func TestMakeBroker(t *testing.T) {
	instance := broker.MakeBroker()

	func(instance interface{}) {
		switch instance.(type) {
		case *broker.Broker:
			t.Logf("Broker instance created successfully.")
		case broker.Broker:
			t.Errorf("Broker instance created successfully, but not as a pointer.")
		default:
			t.Errorf("Broker instance not created successfully. Got type: %T", instance)
		}
	}(instance)
}
