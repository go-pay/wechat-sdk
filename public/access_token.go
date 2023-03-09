package public

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 获取公众号全局唯一后台接口调用凭据（access_token）
// 公众号文档：https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func (s *SDK) getAccessToken() (err error) {
	defer func() {
		if err != nil {
			// reset default refresh internal
			s.RefreshInternal = time.Second * 20
			if s.callback != nil {
				go s.callback("", "", 0, err)
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
		go s.callback(s.Appid, at.AccessToken, at.ExpiresIn, nil)
	}
	return nil
}

func (s *SDK) goAutoRefreshAccessToken() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			xlog.Errorf("public_goAutoRefreshAccessToken: panic recovered: %s\n%s", r, buf)
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

// SetPublicAccessTokenCallback set public access token callback listener
func (s *SDK) SetPublicAccessTokenCallback(fn func(appid, accessToken string, expireIn int, err error)) {
	s.callback = fn
}

// GetPublicAccessToken get public access token string
func (s *SDK) GetPublicAccessToken() (at string) {
	return s.accessToken
}

// SetPublicAccessToken set public access token string
func (s *SDK) SetPublicAccessToken(accessToken string) {
	s.accessToken = accessToken
}
