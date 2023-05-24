package corecontroller

func NewPingController() Controller {
	return &PingController{}
}

type PingController struct {
}

func (c *PingController) Execute([]byte) (interface{}, error) {
	result := outPutPingDto{Message: "pong"}

	return result, nil
}

type outPutPingDto struct {
	Message string `json:"message"`
}
