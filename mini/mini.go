package mini

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/bm"
	"github.com/go-pay/util"
	"github.com/go-pay/util/js"
	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/xhttp"
	"github.com/go-pay/xlog"
)

type SDK struct {
	ctx             context.Context
	DebugSwitch     wechat.DebugSwitch
	Appid           string
	Secret          string
	Host            string
	accessToken     string
	RefreshInternal time.Duration
	hc              *xhttp.Client
	logger          xlog.XLogger

	callback func(appid, accessToken string, expireIn int, err error)
}

// New 初始化微信小程序 SDK
// Appid：Appid
// Secret：appSecret
// autoManageToken：是否自动获取并自动维护刷新 AccessToken，默认使用稳定版接口且force_refresh=false
func New(appid, secret string, autoManageToken bool) (m *SDK, err error) {
	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)
	m = &SDK{
		ctx:         context.Background(),
		DebugSwitch: wechat.DebugOff,
		Appid:       appid,
		Secret:      secret,
		Host:        HostDefault,
		hc:          xhttp.NewClient(),
		logger:      logger,
	}
	if autoManageToken {
		if err = m.getStableAccessToken(); err != nil {
			return nil, err
		}
		go m.goAutoRefreshStableAccessToken()
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
		s.logger.Debugf("Wechat_SDK_URI: %s", uri)
	}
	req := s.hc.Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Get(uri).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(GET, %s), err:%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return res, fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}

func doRequestGet(c context.Context, uri string, ptr any) (err error) {
	req := xhttp.NewClient().Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	_, bs, err := req.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s), err:%w", uri, err)
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}

func doRequestPost(c context.Context, url string, body bm.BodyMap, ptr any) (err error) {
	req := xhttp.NewClient().Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	_, bs, err := req.Post(url).SendBodyMap(body).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(POST, %s), err:%w", url, err)
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
