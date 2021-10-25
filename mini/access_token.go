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
			s.refreshInternal = time.Second * 20
		}
	}()

	at := new(AccessToken)
	path := "/cgi-bin/token?grant_type=client_credential&appid=" + s.Appid + "&secret=" + s.Secret
	if err = s.doRequestGet(s.ctx, path, at); err != nil {
		return err
	}
	if at.Errcode != wechat.Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
	}
	s.accessToken = at.AccessToken
	s.refreshInternal = time.Second * time.Duration(at.ExpiresIn)
	return nil
}

func (s *SDK) autoRefreshAccessToken() {
	for {
		// every one hour, request new access token, default 10s
		time.Sleep(s.refreshInternal / 2)
		err := s.getAccessToken()
		if err != nil {
			xlog.Errorf("get access token error, after 10s retry: %+v", err)
			continue
		}
	}
}

func (s *SDK) GetAccessToken() (at AccessToken) {
	return AccessToken{
		AccessToken: s.accessToken,
		ExpiresIn:   int(s.refreshInternal),
	}
}
