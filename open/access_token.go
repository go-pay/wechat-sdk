package open

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 刷新或续期 access_token 使用
// 微信开放平台文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html
func (s *SDK) refreshAccessToken() (err error) {
	defer func() {
		if err != nil {
			// reset default refresh internal
			s.RefreshInternal = time.Second * 20
			if s.callback != nil {
				go s.callback(nil, err)
			}
		}
	}()
	path := "/sns/oauth2/refresh_token?grant_type=refresh_token&appid=" + s.Appid + "&refresh_token=" + s.refreshToken
	at := &AccessToken{}
	if err = s.DoRequestGet(s.ctx, path, at); err != nil {
		return
	}
	if at.Errcode != Success {
		err = fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
		return
	}
	s.accessToken = at.AccessToken
	s.refreshToken = at.RefreshToken
	s.RefreshInternal = time.Second * time.Duration(at.ExpiresIn)
	if s.callback != nil {
		go s.callback(at, nil)
	}
	return nil
}

func (s *SDK) goAutoRefreshAccessToken() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			xlog.Errorf("open_goAutoRefreshAccessToken: panic recovered: %s\n%s", r, buf)
		}
	}()
	for {
		// every one hour, request new access token, default 10s
		time.Sleep(s.RefreshInternal / 2)
		err := s.refreshAccessToken()
		if err != nil {
			xlog.Errorf("get access token error, after 10s retry: %+v", err)
			continue
		}
	}
}

// SetOpenAccessTokenCallback open access token callback listener
func (s *SDK) SetOpenAccessTokenCallback(fn func(at *AccessToken, err error)) {
	s.callback = fn
}

// GetOpenAccessToken get open access_token string
func (s *SDK) GetOpenAccessToken() (at string) {
	return s.accessToken
}

// SetOpenAccessToken set open access token string
func (s *SDK) SetOpenAccessToken(accessToken string) {
	s.accessToken = accessToken
}

// Code2AccessToken 获取开放平台全局唯一后台接口调用凭据（access_token）
// 注意：必须换取 开放平台 自己的AccessToken，与小程序和公众号不通用
// 微信开放平台文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html
func (s *SDK) Code2AccessToken(c context.Context, code string) (at *AccessToken, err error) {
	path := "/sns/oauth2/access_token?grant_type=authorization_code&appid=" + s.Appid + "&secret=" + s.Secret + "&code=" + code
	at = &AccessToken{}
	if err = s.DoRequestGet(c, path, at); err != nil {
		return
	}
	if at.Errcode != Success {
		err = fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
		return
	}
	s.accessToken = at.AccessToken
	s.refreshToken = at.RefreshToken
	s.RefreshInternal = time.Second * time.Duration(at.ExpiresIn)
	if s.callback != nil {
		go s.callback(at, nil)
	}
	if s.autoManageToken {
		// 自动刷新 AccessToken
		go s.goAutoRefreshAccessToken()
	}
	return at, nil
}

// CheckAccessToken check access_token is ok
// 文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (s *SDK) CheckAccessToken(c context.Context, openid string) (err error) {
	path := "/sns/auth?access_token=" + s.accessToken + "&openid=" + openid
	ec := &ErrorCode{}
	if err = s.DoRequestGet(c, path, ec); err != nil {
		return
	}
	if ec.Errcode != Success {
		err = fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
		return
	}
	return nil
}
