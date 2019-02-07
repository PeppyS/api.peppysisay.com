package background

type Message struct {
	Type string
}

type Queue struct {
	Messages chan Message
}

func NewQueue(size int) *Queue {
	return &Queue{make(chan Message, size)}
}

func (q *Queue) QueueMessage(messageType string) {
	q.Messages <- Message{Type: messageType}
}
