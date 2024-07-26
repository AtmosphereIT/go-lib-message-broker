package channel_test

import (
	"github.com/AtmosphereIT/go-lib-message-broker/channel"
	"testing"
)

func TestIsNameAvailable_ReturnsTrueWhenNameIsAvailable(t *testing.T) {
	topicMap := &channel.TopicMap{Topics: make(map[string]*channel.Topic)}
	if !topicMap.IsNameAvailable("new_topic") {
		t.Errorf("expected true, got false")
	}
}

func TestIsNameAvailable_ReturnsFalseWhenNameIsNotAvailable(t *testing.T) {
	topicMap := &channel.TopicMap{Topics: map[string]*channel.Topic{"existing_topic": {Name: "existing_topic"}}}
	if topicMap.IsNameAvailable("existing_topic") {
		t.Errorf("expected false, got true")
	}
}

func TestAddTopic_AddsNewTopicWhenNameIsAvailable(t *testing.T) {
	topicMap := &channel.TopicMap{Topics: make(map[string]*channel.Topic)}
	err := topicMap.AddTopic("new_topic", "description")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if topicMap.Topics["new_topic"] == nil {
		t.Errorf("expected topic to be added, but it was not")
	}
}

func TestAddTopic_ReturnsErrorWhenNameIsNotAvailable(t *testing.T) {
	topicMap := &channel.TopicMap{Topics: map[string]*channel.Topic{"existing_topic": {Name: "existing_topic"}}}
	err := topicMap.AddTopic("existing_topic", "description")
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if err.Error() != channel.ERROR_TOPIC_NAME_NOT_UNIQUE {
		t.Errorf("expected error %s, got %s", channel.ERROR_TOPIC_NAME_NOT_UNIQUE, err.Error())
	}

	if topicMap.Topics["existing_topic"] == nil {
		t.Errorf("expected topic to not be overwritten, but it was")
	}
}
