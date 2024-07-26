package channel

import "errors"

// IsNameAvailable - Check if a channel name is available.
func (t *TopicMap) IsNameAvailable(name string) bool {
	if t.Topics[name] == nil {
		return true
	}
	return false
}

// AddTopic - Add a new topic into the topic map.
func (t *TopicMap) AddTopic(name string, description string) error {
	if !t.IsNameAvailable(name) {
		return errors.New(ERROR_TOPIC_NAME_NOT_UNIQUE)
	}

	t.Topics[name] = &Topic{
		Name: name,
	}

	return nil
}
