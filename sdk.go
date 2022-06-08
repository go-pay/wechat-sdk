package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-pay/wechat-sdk/mini"
	"github.com/go-pay/wechat-sdk/open"
	"github.com/go-pay/wechat-sdk/pkg/util"
	"github.com/go-pay/wechat-sdk/pkg/xhttp"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
	"github.com/go-pay/wechat-sdk/public"
)

type SDK struct {
	ctx             context.Context
	rwMu            sync.RWMutex
	Appid           string
	Secret          string
	Host            string
	RefreshInternal time.Duration
	DebugSwitch     DebugSwitch

	plat        Platform
	accessToken string
	atChanMap   map[string]chan string
	callback    func(accessToken string, expireIn int, err error)
}

// NewSDK 初始化微信 SDK
//	plat：wechat.PlatformMiniOrPublic 或 wechat.PlatformOpen
//	appid：Appid
//	secret：appSecret
func NewSDK(plat Platform, appid, secret string) (sdk *SDK, err error) {
	sdk = &SDK{
		ctx:             context.Background(),
		Appid:           appid,
		Secret:          secret,
		atChanMap:       make(map[string]chan string),
		Host:            HostMap[HostDefault],
		RefreshInternal: time.Second * 20,
		DebugSwitch:     DebugOff,
		plat:            plat,
	}
	switch plat {
	case PlatformMini:
		// 获取AccessToken
		err = sdk.getAccessToken()
		if err != nil {
			return nil, err
		}
		// auto refresh access token
		go sdk.goAutoRefreshAccessToken()
	case PlatformPublic:
		// 获取AccessToken
		err = sdk.getAccessToken()
		if err != nil {
			return nil, err
		}
		// auto refresh access token
		go sdk.goAutoRefreshAccessToken()
	case PlatformOpen:
		// 需要开放平台Client单独调用 s.Code2AccessToken()
	default:
		return nil, fmt.Errorf("unsupported platform: %s", plat)
	}
	return
}

// SetMiniOrPublicAT 若 NewSDK() 时自传 AccessToken，则后续更新替换请调用此方法
func (s *SDK) SetMiniOrPublicAT(accessToken string) {
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

// NewMini new 微信小程序
func (s *SDK) NewMini() (m *mini.SDK, err error) {
	if s.plat != PlatformMini {
		return nil, fmt.Errorf("invalid platform: %s", s.plat)
	}
	s.rwMu.Lock()
	s.atChanMap[Mini] = make(chan string, 1)
	s.rwMu.Unlock()

	c := &mini.Config{
		Appid:       s.Appid,
		Secret:      s.Secret,
		AccessToken: s.accessToken,
		Host:        s.Host,
	}
	return mini.New(c, int8(s.DebugSwitch), s.atChanMap[Mini]), nil
}

// NewPublic new 微信公众号
func (s *SDK) NewPublic() (p *public.SDK, err error) {
	if s.plat != PlatformPublic {
		return nil, fmt.Errorf("invalid platform: %s", s.plat)
	}
	s.rwMu.Lock()
	s.atChanMap[Public] = make(chan string, 1)
	s.rwMu.Unlock()

	c := &public.Config{
		Appid:       s.Appid,
		Secret:      s.Secret,
		AccessToken: s.accessToken,
		Host:        s.Host,
	}
	return public.New(c, int8(s.DebugSwitch), s.atChanMap[Public]), nil
}

// NewOpen new 微信开放平台
func (s *SDK) NewOpen() (o *open.SDK, err error) {
	if s.plat != PlatformOpen {
		return nil, fmt.Errorf("invalid platform: %s", s.plat)
	}
	c := &open.Config{
		Ctx:    s.ctx,
		Appid:  s.Appid,
		Secret: s.Secret,
		Host:   s.Host,
	}
	return open.New(c, int8(s.DebugSwitch)), nil
}

func (s *SDK) DoRequestGet(c context.Context, path string, ptr interface{}) (err error) {
	uri := s.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == DebugOn {
		xlog.Debugf("Wechat_SDK_URI: %s", uri)
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}
	if s.DebugSwitch == DebugOn {
		xlog.Debugf("Wechat_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
