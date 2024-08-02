package open

import (
	"context"
	"fmt"
)

// UserInfo 获取用户个人信息（UnionID 机制）
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html
func (s *SDK) UserInfo(c context.Context, accessToken, openid, lan string) (ui *UserInfo, err error) {
	switch lan {
	case "zh_CN", "zh_TW", "en":
	default:
		lan = "en"
	}
	path := "/sns/userinfo?access_token=" + accessToken + "&openid=" + openid + "&lang=" + lan
	ui = &UserInfo{}
	if _, err = s.DoRequestGet(c, path, ui); err != nil {
		return
	}
	if ui.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", ui.Errcode, ui.Errmsg)
	}
	return ui, nil
}
