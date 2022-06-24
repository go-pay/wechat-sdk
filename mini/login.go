package mini

import (
	"context"
	"fmt"
)

// Code2Session 登录凭证校验
//	注意：errcode = 0 为成功
//	wxCode:小程序调用 wx.login 获取的code
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (s *SDK) Code2Session(c context.Context, wxCode string) (session *Code2Session, err error) {
	path := "/sns/jscode2session?appid=" + s.Appid + "&secret=" + s.Secret + "&js_code=" + wxCode + "&grant_type=authorization_code"
	session = &Code2Session{}
	if err = s.doRequestGet(c, path, session); err != nil {
		return nil, err
	}
	if session.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", session.Errcode, session.Errmsg)
	}
	return session, nil
}
