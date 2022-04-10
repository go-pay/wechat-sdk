package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-pay/wechat-sdk/model"
	"github.com/go-pay/wechat-sdk/pkg/util"
	"github.com/go-pay/wechat-sdk/pkg/xhttp"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

type OpenSDK struct {
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

// NewOpenSDK 初始化微信开放平台 SDK
//	Appid：Appid
//	Secret：appSecret
//	accessToken：AccessToken，若此参数为空，则自动获取并自动维护刷新
func NewOpenSDK(appid, secret string, accessToken ...string) (sdk *OpenSDK, err error) {
	sdk = &OpenSDK{
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
	return
}

func (s *OpenSDK) DoRequestGet(c context.Context, path string, ptr interface{}) (err error) {
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
