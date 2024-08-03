package mini

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// 获取小程序全局唯一后台接口调用凭据（access_token）
// 微信小程序文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
//func (s *SDK) getAccessToken() (err error) {
//	defer func() {
//		if err != nil {
//			// reset default refresh internal
//			s.RefreshInternal = time.Second * 20
//			if s.callback != nil {
//				go s.callback("", "", 0, err)
//			}
//		}
//	}()
//
//	path := "/cgi-bin/token?grant_type=client_credential&appid=" + s.Appid + "&secret=" + s.Secret
//	at := &AccessToken{}
//	if _, err = s.DoRequestGet(s.ctx, path, at); err != nil {
//		return
//	}
//	if at.Errcode != Success {
//		err = fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
//		return
//	}
//	s.accessToken = at.AccessToken
//	s.RefreshInternal = time.Second * time.Duration(at.ExpiresIn)
//	if s.callback != nil {
//		go s.callback(s.Appid, at.AccessToken, at.ExpiresIn, nil)
//	}
//	return nil
//}

//func (s *SDK) goAutoRefreshAccessToken() {
//	defer func() {
//		if r := recover(); r != nil {
//			buf := make([]byte, 64<<10)
//			buf = buf[:runtime.Stack(buf, false)]
//			s.logger.Errorf("mini_goAutoRefreshAccessToken: panic recovered: %s\n%s", r, buf)
//		}
//	}()
//	for {
//		// every one hour, request new access token, default 10s
//		time.Sleep(s.RefreshInternal / 2)
//		err := s.getAccessToken()
//		if err != nil {
//			s.logger.Errorf("get access token error, after 10s retry: %+v", err)
//			continue
//		}
//	}
//}

// 获取稳定版接口调用凭据
// 微信小程序文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getStableAccessToken.html
func (s *SDK) getStableAccessToken() (err error) {
	defer func() {
		if err != nil {
			// reset default refresh internal
			s.RefreshInternal = time.Second * 20
			if s.callback != nil {
				go s.callback("", "", 0, err)
			}
		}
	}()

	path := "/cgi-bin/stable_token?grant_type=client_credential&appid=" + s.Appid + "&secret=" + s.Secret + "&force_refresh=false"
	at := &AccessToken{}
	if _, err = s.DoRequestGet(s.ctx, path, at); err != nil {
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

func (s *SDK) goAutoRefreshStableAccessToken() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			s.logger.Errorf("mini_goAutoRefreshAccessToken: panic recovered: %s\n%s", r, buf)
			time.Sleep(time.Second * 3)
			if err := s.getStableAccessToken(); err != nil {
				// 失败就不再自动刷新了
				return
			}
			s.goAutoRefreshStableAccessToken()
		}
	}()
	for {
		// every one hour, request new access token, default 10s
		time.Sleep(s.RefreshInternal / 2)
		err := s.getStableAccessToken()
		if err != nil {
			s.logger.Errorf("get access token error, after 10s retry: %+v", err)
			continue
		}
	}
}

// SetMiniAccessTokenCallback set mini access token callback listener
func (s *SDK) SetMiniAccessTokenCallback(fn func(appid, accessToken string, expireIn int, err error)) {
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

// =====================================================================================================================

// 获取接口调用凭据
// 微信小程序文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
func GetAccessToken(c context.Context, appid, secret string) (at *AccessToken, err error) {
	uri := HostDefault + "/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret
	at = &AccessToken{}
	if err = doRequestGet(c, uri, at); err != nil {
		return nil, err
	}
	if at.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
	}
	return at, nil
}

// 获取稳定版接口调用凭据
// 微信小程序文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getStableAccessToken.html
func GetStableAccessToken(c context.Context, appid, secret string, forceRefresh bool) (at *AccessToken, err error) {
	uri := HostDefault + "/cgi-bin/stable_token?grant_type=client_credential&appid=" + appid + "&secret=" + secret + "&force_refresh=" + strconv.FormatBool(forceRefresh)
	at = &AccessToken{}
	if err = doRequestGet(c, uri, at); err != nil {
		return nil, err
	}
	if at.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
	}
	return at, nil
}
