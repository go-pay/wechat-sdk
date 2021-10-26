package mini

import (
	"context"
)

// Code2Session 登录凭证校验
//	wxCode:小程序调用 wx.login 获取的code
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (s *SDK) Code2Session(c context.Context, wxCode string) (session *Code2Session, err error) {
	path := "/sns/jscode2session?appid=" + s.appid + "&secret=" + s.secret + "&js_code=" + wxCode + "&grant_type=authorization_code"
	session = new(Code2Session)
	if err = s.doRequestGet(c, path, session); err != nil {
		return nil, err
	}
	return
}
