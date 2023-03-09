package mini

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/wechat-sdk/pkg/util"
	"github.com/go-pay/wechat-sdk/pkg/xhttp"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

type SDK struct {
	ctx             context.Context
	DebugSwitch     wechat.DebugSwitch
	Appid           string
	Secret          string
	Host            string
	accessToken     string
	RefreshInternal time.Duration

	callback func(appid, accessToken string, expireIn int, err error)
}

// New 初始化微信小程序 SDK
// Appid：Appid
// Secret：appSecret
// autoManageToken：是否自动获取并自动维护刷新 AccessToken
func New(appid, secret string, autoManageToken bool) (m *SDK, err error) {
	m = &SDK{
		ctx:         context.Background(),
		DebugSwitch: wechat.DebugOff,
		Appid:       appid,
		Secret:      secret,
		Host:        HostDefault,
	}
	if autoManageToken {
		if err = m.getAccessToken(); err != nil {
			return nil, err
		}
		go m.goAutoRefreshAccessToken()
	}
	return
}

func (s *SDK) DoRequestGet(c context.Context, path string, ptr interface{}) (err error) {
	uri := s.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == wechat.DebugOn {
		xlog.Debugf("Wechat_SDK_URI: %s", uri)
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
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
