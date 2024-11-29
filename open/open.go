package open

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/smap"
	"github.com/go-pay/util"
	"github.com/go-pay/util/js"
	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/xhttp"
	"github.com/go-pay/xlog"
)

type SDK struct {
	ctx                      context.Context
	DebugSwitch              wechat.DebugSwitch
	Appid                    string
	Secret                   string
	Host                     string
	autoManageToken          bool                           // 是否自动维护刷新 AccessToken
	autoRefreshTokenInternal time.Duration                  // 自动刷新 token 的间隔时间
	openidAccessTokenMap     smap.Map[string, *AccessToken] // key: openid
	hc                       *xhttp.Client
	logger                   xlog.XLogger

	callback func(at *AT, err error)
}

// New 初始化微信开放平台 SDK
// Appid：Appid
// Secret：appSecret
// autoManageToken：是否自动维护刷新 AccessToken（用户量较少时推荐使用，默认10分钟轮询检测一次，发现有效期小于1.5倍轮询时间时，自动刷新）
func New(appid, secret string, autoManageToken bool) (o *SDK) {
	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)
	o = &SDK{
		ctx:             context.Background(),
		DebugSwitch:     wechat.DebugOff,
		Host:            HostDefault,
		Appid:           appid,
		Secret:          secret,
		autoManageToken: autoManageToken,
		hc:              xhttp.NewClient(),
		logger:          logger,
	}
	if autoManageToken {
		o.autoRefreshTokenInternal = time.Minute * 10
		go o.goAutoRefreshAccessTokenJob()
	}
	return
}

// SetHttpClient 设置自定义的xhttp.Client
func (s *SDK) SetHttpClient(client *xhttp.Client) {
	if client != nil {
		s.hc = client
	}
}

func (s *SDK) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		s.logger = logger
	}
}

func (s *SDK) DoRequestGet(c context.Context, path string, ptr any) (res *http.Response, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Open_SDK_URI: %s", uri)
	}
	req := s.hc.Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Get(uri).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(GET, %s), err:%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Open_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return res, fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
