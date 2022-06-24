package open

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
	autoManageToken bool
	accessToken     string
	refreshToken    string
	RefreshInternal time.Duration

	callback func(at *AccessToken, err error)
}

func New(appid, secret string, autoManageToken bool) (o *SDK) {
	o = &SDK{
		ctx:             context.Background(),
		DebugSwitch:     DebugOff,
		Host:            HostDefault,
		Appid:           appid,
		Secret:          secret,
		autoManageToken: autoManageToken,
	}
	return
}

func (s *SDK) DoRequestGet(c context.Context, path string, ptr interface{}) (err error) {
	uri := s.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == DebugOn {
		xlog.Debugf("Wechat_Open_SDK_URI: %s", uri)
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}
	if s.DebugSwitch == DebugOn {
		xlog.Debugf("Wechat_Open_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
