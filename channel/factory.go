package channel

// MakeChannel - Create a new channel instance.
func MakeChannel(name string, topics *TopicMap) *Channel {

	// When there were no topics passed in we initialize a new topics map.
	//
	if topics == nil {
		topics = &TopicMap{
			Topics: make(map[string]*Topic),
		}
	}

	return &Channel{
		Name:   name,
		Topics: topics,
	}
}
