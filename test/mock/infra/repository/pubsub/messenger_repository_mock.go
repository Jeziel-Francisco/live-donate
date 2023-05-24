package testmockinfrarepositorypubsub

func NewMessageMock(sendTopicError error, sendQueueError error) *MessengerRepositoryMock {
	return &MessengerRepositoryMock{sendTopicError: sendTopicError, sendQueueError: sendQueueError}
}

type MessengerRepositoryMock struct {
	sendTopicError error
	sendQueueError error
}

func (pb MessengerRepositoryMock) SendTopic(message interface{}, name string) error {
	return pb.sendTopicError
}

func (pb MessengerRepositoryMock) SendQueue(message interface{}, name string) error {
	return pb.sendQueueError
}
