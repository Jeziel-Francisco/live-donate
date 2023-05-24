package usecaseinterface

type MessengerRepository interface {
	SendTopic(message interface{}, name string) error
	SendQueue(message interface{}, name string) error
}
