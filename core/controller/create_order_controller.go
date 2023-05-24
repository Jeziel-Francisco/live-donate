package corecontroller

func NewCreateOrderController() Controller {
	return &createOrderController{}
}

type createOrderController struct {
}

func (c *createOrderController) Execute([]byte) (interface{}, error) {
	result := outPutPingDto{Message: "pong"}

	return result, nil
}
