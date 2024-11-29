package open

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func (s *SDK) refreshAccessToken(openid, refreshToken string) {
	var (
		path = "/sns/oauth2/refresh_token?grant_type=refresh_token&appid=" + s.Appid + "&refresh_token=" + refreshToken
		at   = &AccessToken{}
		err  error
	)
	if _, err = s.DoRequestGet(s.ctx, path, at); err != nil {
		if s.callback != nil {
			go s.callback(nil, fmt.Errorf("openid[%s],refresh_token[%s] refresh access token failed: %w", openid, refreshToken, err))
		}
		return
	}
	if at.Errcode != Success {
		if s.callback != nil {
			go s.callback(nil, fmt.Errorf("openid[%s],refresh_token[%s] refresh access token failed: errcode(%d), errmsg(%s)", openid, refreshToken, at.Errcode, at.Errmsg))
		}
		return
	}

	s.openidAccessTokenMap.Store(at.Openid, at)
	if s.callback != nil {
		go s.callback(&AT{
			AccessToken:  at.AccessToken,
			ExpiresIn:    at.ExpiresIn,
			RefreshToken: at.RefreshToken,
			Openid:       at.Openid,
			Scope:        at.Scope,
			Unionid:      at.Unionid,
		}, nil)
	}
}

func (s *SDK) goAutoRefreshAccessTokenJob() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			s.logger.Errorf("open_goAutoRefreshAccessToken: panic recovered: %s\n%s", r, buf)
			time.Sleep(time.Second * 3)
			s.goAutoRefreshAccessTokenJob()
		}
	}()
	for {
		// request new access token, default internal 10min
		time.Sleep(s.autoRefreshTokenInternal)
		s.openidAccessTokenMap.Range(func(k string, v *AccessToken) bool {
			// 有效期小于1.5倍轮询时间时，自动刷新）
			if time.Duration(v.ExpiresIn)*time.Second < (s.autoRefreshTokenInternal*3)/2 {
				s.refreshAccessToken(k, v.RefreshToken)
			}
			return true
		})
	}
}

// SetAccessTokenCallback access token callback listener
func (s *SDK) SetAccessTokenCallback(fn func(at *AT, err error)) {
	s.callback = fn
}

// SetAccessTokenRefreshInternal 设置自动刷新 access_token 间隔时长，默认10分钟
func (s *SDK) SetAccessTokenRefreshInternal(internal time.Duration) {
	s.autoRefreshTokenInternal = internal
}

// GetAccessTokenMap 获取 access_token map，key 为 openid
func (s *SDK) GetAccessTokenMap() (openidATMap map[string]*AT) {
	openidATMap = make(map[string]*AT)
	s.openidAccessTokenMap.Range(func(k string, v *AccessToken) bool {
		openidATMap[k] = &AT{
			AccessToken:  v.AccessToken,
			ExpiresIn:    v.ExpiresIn,
			RefreshToken: v.RefreshToken,
			Openid:       v.Openid,
			Scope:        v.Scope,
			Unionid:      v.Unionid,
		}
		return true
	})
	return
}

// DelAccessToken 根据 openid 删除 map 中维护的 access_token
func (s *SDK) DelAccessToken(openid string) {
	s.openidAccessTokenMap.Delete(openid)
}

// Code2AccessToken 通过 code 获取用户 access_token
// 微信开放平台文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html
func (s *SDK) Code2AccessToken(c context.Context, code string) (at *AccessToken, err error) {
	path := "/sns/oauth2/access_token?grant_type=authorization_code&appid=" + s.Appid + "&secret=" + s.Secret + "&code=" + code
	at = &AccessToken{}
	if _, err = s.DoRequestGet(c, path, at); err != nil {
		return nil, err
	}
	if at.Errcode != Success {
		err = fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
		return nil, err
	}
	if s.callback != nil {
		go s.callback(&AT{
			AccessToken:  at.AccessToken,
			ExpiresIn:    at.ExpiresIn,
			RefreshToken: at.RefreshToken,
			Openid:       at.Openid,
			Scope:        at.Scope,
			Unionid:      at.Unionid,
		}, nil)
	}
	if s.autoManageToken {
		s.openidAccessTokenMap.Store(at.Openid, at)
	}
	return at, nil
}

// RefreshAccessToken 刷新或续期 access_token
// 文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (s *SDK) RefreshAccessToken(c context.Context, refreshToken string) (at *AccessToken, err error) {
	path := "/sns/oauth2/refresh_token?grant_type=refresh_token&appid=" + s.Appid + "&refresh_token=" + refreshToken
	at = &AccessToken{}
	if _, err = s.DoRequestGet(s.ctx, path, at); err != nil {
		return nil, err
	}
	if at.Errcode != Success {
		err = fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
		return nil, err
	}
	if s.callback != nil {
		go s.callback(&AT{
			AccessToken:  at.AccessToken,
			ExpiresIn:    at.ExpiresIn,
			RefreshToken: at.RefreshToken,
			Openid:       at.Openid,
			Scope:        at.Scope,
			Unionid:      at.Unionid,
		}, nil)
	}
	if s.autoManageToken {
		s.openidAccessTokenMap.Store(at.Openid, at)
	}
	return at, nil
}

// CheckAccessToken 检验授权凭证 access_token 是否有效
// 注意：当开启自动管理 access_token 时，会自动删除无效的 access_token
// 文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (s *SDK) CheckAccessToken(c context.Context, accessToken, openid string) (err error) {
	path := "/sns/auth?access_token=" + accessToken + "&openid=" + openid
	ec := &ErrorCode{}
	if _, err = s.DoRequestGet(c, path, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		if s.autoManageToken {
			s.DelAccessToken(openid)
		}
		err = fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
		return err
	}
	return nil
}
