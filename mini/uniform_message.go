package mini

import (
	"context"

	"github.com/go-pay/wechat-sdk/pkg/bm"
)

// UniformMessageSend 发送统一服务消息
//	注意：小程序模板消息已下线，不用传 weapp_template_msg 此节点
//	toUser：用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
//	mpMsg：对应 mp_template_msg 的value值，BodyMap key-value 格式传入
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (s *SDK) UniformMessageSend(c context.Context, toUser string, mpMsg bm.BodyMap) (ec *ErrorCode, err error) {
	path := "/cgi-bin/message/wxopen/template/uniform_send?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("touser", toUser)

	ec = &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return nil, err
	}
	return
}
