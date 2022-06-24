package mini

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 获取小程序全局唯一后台接口调用凭据（access_token）
//	微信小程序文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
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

	path := "/cgi-bin/token?grant_type=client_credential&appid=" + s.Appid + "&secret=" + s.Secret
	at := &AccessToken{}
	if err = s.DoRequestGet(s.ctx, path, at); err != nil {
		return
	}
	if at.Errcode != Success {
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

func (s *SDK) goAutoRefreshAccessToken() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			xlog.Errorf("mini_goAutoRefreshAccessToken: panic recovered: %s\n%s", r, buf)
		}
	}()
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

// SetMiniAccessTokenCallback set mini access token callback listener
func (s *SDK) SetMiniAccessTokenCallback(fn func(accessToken string, expireIn int, err error)) {
	s.callback = fn
}

// GetMiniAccessToken get mini access token string
func (s *SDK) GetMiniAccessToken() (at string) {
	return s.accessToken
}

// SetMiniAccessToken set mini access token string
func (s *SDK) SetMiniAccessToken(accessToken string) {
	s.accessToken = accessToken
}
