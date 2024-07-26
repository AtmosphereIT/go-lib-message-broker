# Go - Message Broker

This library provides a simple message broker that can be used to send messages between different parts of a system.

### Installation

To include this library in your project, you can use the `go get` command:

```bash
go get github.com/AtmosphereIT/go-lib-message-broker
```

## Concepts

### Broker

A `Broker` is a central hub that can be used to send messages between different parts of a system. It is responsible for
managing the channels and routing messages between them.

### Channel

A `Channel` is a named pipe that can be used to send messages between different parts of a system.

### Message

When an event occurs, a `Message` is created. A `Message` is a structure that captures the data associated with the
event.
It contains the following fields:

- `type` - type of event that occurred,
- `data` - data associated with the event,
- `meta` - metadata associated with the event,
    - `timestamp` - time when the event occurred,
    - `source` - the source of the event.

Messages are immutable and cannot be modified once created. Messages should be as small as possible to reduce the
overhead of passing them around.

