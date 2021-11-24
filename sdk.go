package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-pay/wechat-sdk/mini"
	"github.com/go-pay/wechat-sdk/model"
	"github.com/go-pay/wechat-sdk/open"
	"github.com/go-pay/wechat-sdk/pkg/util"
	"github.com/go-pay/wechat-sdk/pkg/xhttp"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

type SDK struct {
	ctx             context.Context
	rwMu            sync.RWMutex
	Appid           string
	Secret          string
	accessToken     string
	atChanMap       map[string]chan string
	callback        func(accessToken string, expireIn int, err error)
	Host            string
	RefreshInternal time.Duration
	DebugSwitch     DebugSwitch
}

// NewSDK 初始化微信 SDK
//	Appid：Appid
//	Secret：appSecret
//	accessToken：AccessToken，若此参数为空，则自动获取并自动维护刷新
func NewSDK(appid, secret string, accessToken ...string) (sdk *SDK, err error) {
	sdk = &SDK{
		ctx:             context.Background(),
		Appid:           appid,
		Secret:          secret,
		atChanMap:       make(map[string]chan string),
		Host:            HostMap[HostDefault],
		RefreshInternal: time.Second * 20,
		DebugSwitch:     DebugOff,
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

// SetAccessToken 若 NewSDK() 时自传 AccessToken，则后续更新替换请调用此方法
func (s *SDK) SetAccessToken(accessToken string) {
	s.accessToken = accessToken
	if len(s.atChanMap) > 0 {
		for _, v := range s.atChanMap {
			v <- accessToken
		}
	}
}

// SetHost 设置微信请求Host
//	上海、深圳、香港 等
func (s *SDK) SetHost(host Host) (sdk *SDK) {
	if h, ok := HostMap[host]; ok {
		s.Host = h
	}
	return s
}

// NewOpen new 微信公众号
func (s *SDK) NewOpen() (o *open.SDK) {
	s.rwMu.Lock()
	defer s.rwMu.Unlock()
	s.atChanMap[model.Open] = make(chan string, 1)

	c := &model.Config{
		Appid:       s.Appid,
		Secret:      s.Secret,
		AccessToken: s.accessToken,
		Host:        s.Host,
	}

	return open.New(c, int8(s.DebugSwitch), s.atChanMap[model.Mini])
}

// NewMini new 微信小程序
func (s *SDK) NewMini() (m *mini.SDK) {
	s.rwMu.Lock()
	defer s.rwMu.Unlock()
	s.atChanMap[model.Mini] = make(chan string, 1)

	c := &model.Config{
		Appid:       s.Appid,
		Secret:      s.Secret,
		AccessToken: s.accessToken,
		Host:        s.Host,
	}

	return mini.New(c, int8(s.DebugSwitch), s.atChanMap[model.Mini])
}

func (s *SDK) DoRequestGet(c context.Context, path string, ptr interface{}) (err error) {
	uri := s.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_SDK_URI: %s", uri)
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
