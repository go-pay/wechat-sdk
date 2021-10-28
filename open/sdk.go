package open

import (
	"context"
	"time"

	"github.com/go-pay/wechat-sdk"
)

type SDK struct {
	ctx             context.Context
	appid           string
	secret          string
	accessToken     string
	callback        func(accessToken string, expireIn int, err error)
	Host            string
	RefreshInternal time.Duration
	DebugSwitch     wechat.DebugSwitch
}
