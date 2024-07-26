package channel

type Channel struct {
	// Name - The name of the channel.
	Name string

	// Topics - A map of topics.
	Topics *TopicMap
}

type TopicMap struct {
	// Topics - A map of topics.
	Topics map[string]*Topic
}

type Topic struct {
	// Name - The name of the topic.
	Name string

	// Description - A description of the topic.
	Description string
}
