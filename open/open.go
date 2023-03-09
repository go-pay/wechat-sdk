package open

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/wechat-sdk/pkg/util"
	"github.com/go-pay/wechat-sdk/pkg/xhttp"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

type SDK struct {
	ctx                      context.Context
	DebugSwitch              wechat.DebugSwitch
	mu                       sync.RWMutex
	Appid                    string
	Secret                   string
	Host                     string
	autoManageToken          bool                    // 是否自动维护刷新 AccessToken
	autoRefreshTokenInternal time.Duration           // 自动刷新 token 的间隔时间
	minRefreshTokenDuration  time.Duration           // 最小刷新 token 的阈值时长
	openidAccessTokenMap     map[string]*AccessToken // key: openid

	callback func(at *AT, err error)
}

// New 初始化微信开放平台 SDK
// Appid：Appid
// Secret：appSecret
// autoManageToken：是否自动维护刷新 AccessToken（用户量较少时推荐使用，默认10分钟轮询检测一次，发现有效期小于1.5倍轮询时间时，自动刷新）
func New(appid, secret string, autoManageToken bool) (o *SDK) {
	o = &SDK{
		ctx:             context.Background(),
		DebugSwitch:     wechat.DebugOff,
		Host:            HostDefault,
		Appid:           appid,
		Secret:          secret,
		autoManageToken: autoManageToken,
	}
	if autoManageToken {
		o.autoRefreshTokenInternal = time.Minute * 10
		o.openidAccessTokenMap = make(map[string]*AccessToken)
		go o.goAutoRefreshAccessTokenJob()
	}
	return
}

func (s *SDK) DoRequestGet(c context.Context, path string, ptr interface{}) (err error) {
	uri := s.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == wechat.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_URI: %s", uri)
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
