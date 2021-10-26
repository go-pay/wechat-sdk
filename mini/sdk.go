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
	appid           string
	secret          string
	accessToken     string
	callback        func(accessToken string, expireIn int, err error)
	Host            string
	RefreshInternal time.Duration
	DebugSwitch     wechat.DebugSwitch
}

// NewSDK 初始化微信小程序 SDK
//	appid：小程序 appid
//	secret：小程序 appSecret
//	accessToken：微信小程序AccessToken，若此参数为空，则自动获取并自动维护刷新
func NewSDK(appid, secret string, accessToken ...string) (sdk *SDK, err error) {
	sdk = &SDK{
		ctx:             context.Background(),
		appid:           appid,
		secret:          secret,
		Host:            wechat.HostMap[wechat.HostDefault],
		RefreshInternal: time.Second * 20,
		DebugSwitch:     wechat.DebugOn,
	}
	if len(accessToken) >= 1 {
		sdk.accessToken = accessToken[0]
		return
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
func (s *SDK) SetHost(host wechat.Host) (sdk *SDK) {
	if h, ok := wechat.HostMap[host]; ok {
		s.Host = h
	}
	return s
}

func (s *SDK) doRequestGet(c context.Context, path string, ptr interface{}) (err error) {
	uri := s.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == wechat.DebugOn {
		xlog.Debugf("Wechat_SDK_URI: %s", uri)
	}
	httpClient.SetTimeout(5 * time.Second)
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		xlog.Debugf("Wechat_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
