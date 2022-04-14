package open

import (
	"fmt"
	"time"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

// Code2AccessToken 获取开放平台全局唯一后台接口调用凭据（access_token）
//	微信开放平台文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html
func (s *SDK) Code2AccessToken(code string) (err error) {
	path := "/sns/oauth2/access_token?grant_type=authorization_code&appid=" + s.Appid + "&secret=" + s.Secret + "&code=" + code
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
	if len(s.atChanMap) > 0 {
		for _, v := range s.atChanMap {
			v <- at.AccessToken
		}
	}
	return nil
}

// 刷新或续期 access_token 使用
//	微信开放平台文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html
func (s *SDK) refreshAccessToken() (err error) {
	defer func() {
		if err != nil {
			// reset default refresh internal
			s.RefreshInternal = time.Second * 20
			if s.callback != nil {
				go s.callback("", 0, err)
			}
		}
	}()
	// /sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN
	path := "/sns/oauth2/refresh_token?grant_type=refresh_token&appid=" + s.Appid + "&refresh_token=" + s.Secret
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
	if len(s.atChanMap) > 0 {
		for _, v := range s.atChanMap {
			v <- at.AccessToken
		}
	}
	return nil
}

func (s *SDK) autoRefreshAccessToken() {
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

// SetAccessTokenCallback access token callback listener
func (s *SDK) SetAccessTokenCallback(fn func(accessToken string, expireIn int, err error)) {
	s.callback = fn
}

// GetAccessToken get access token string
func (s *SDK) GetAccessToken() (at string) {
	return s.accessToken
}
