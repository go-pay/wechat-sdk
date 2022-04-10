package open

import "github.com/go-pay/wechat-sdk/model"

type SDK struct {
	Conf        *model.Config
	DebugSwitch int8
	AccessChan  chan string
}
