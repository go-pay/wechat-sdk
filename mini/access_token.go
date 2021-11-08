package mini

import (
	"fmt"
	"time"

	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 获取小程序全局唯一后台接口调用凭据（access_token）
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (s *SDK) getAccessToken() (err error) {
	defer func() {
		if err != nil {
			// reset default refresh internal
			s.RefreshInternal = time.Second * 20
			if s.callback != nil {
				go s.callback("", 0, err)
			}
		}
	}()

	path := "/cgi-bin/token?grant_type=client_credential&appid=" + s.appid + "&secret=" + s.secret
	at := &AccessToken{}
	if err = s.doRequestGet(s.ctx, path, at); err != nil {
		return
	}
	if at.Errcode != wechat.Success {
		err = fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
		return
	}
	s.accessToken = at.AccessToken
	s.RefreshInternal = time.Second * time.Duration(at.ExpiresIn)
	if s.callback != nil {
		go s.callback(at.AccessToken, at.ExpiresIn, nil)
	}
	return nil
}

func (s *SDK) autoRefreshAccessToken() {
	for {
		// every one hour, request new access token, default 10s
		time.Sleep(s.RefreshInternal / 2)
		err := s.getAccessToken()
		if err != nil {
			xlog.Errorf("get access token error, after 10s retry: %+v", err)
			continue
		}
	}
}

// SetAccessTokenCallback access token callback listener
func (s *SDK) SetAccessTokenCallback(fn func(accessToken string, expireIn int, err error)) {
	s.callback = fn
}

// GetAccessToken get access token string
func (s *SDK) GetAccessToken() (at string) {
	return s.accessToken
}
