package mini

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/wechat-sdk/pkg/xhttp"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

type SDK struct {
	sync.RWMutex
	ctx             context.Context
	Appid           string
	Secret          string
	Host            WechatHost
	accessToken     string
	refreshInternal time.Duration
	DebugSwitch     wechat.DebugSwitch
}

// NewSDK 初始化微信小程序 SDK
//	appid：小程序 appid
//	secret：小程序 appSecret
func NewSDK(appid, secret string) (sdk *SDK, err error) {
	sdk = &SDK{
		ctx:             context.Background(),
		Appid:           appid,
		Secret:          secret,
		Host:            BaseHost,
		refreshInternal: time.Second * 20,
		DebugSwitch:     wechat.DebugOff,
	}
	// 获取AccessToken
	err = sdk.getAccessToken()
	if err != nil {
		return nil, err
	}
	// auto refresh access token
	go sdk.autoRefreshAccessToken()
	return
}

// SetHost 设置微信请求Host
//	上海、深圳、香港 等
func (s *SDK) SetHost(host WechatHost) (sdk *SDK) {
	if host != "" {
		s.Host = host
	}
	return s
}

func (s *SDK) doRequestGet(c context.Context, path string, ptr interface{}) (err error) {
	uri := string(BaseHost) + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == wechat.DebugOn {
		xlog.Debugf("Wechat_SDK_URI: %s", uri)
	}
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		xlog.Debugf("Wechat_SDK_Response: %d -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
