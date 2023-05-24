package corecontroller

type Controller interface {
	Execute([]byte) (interface{}, error)
}
