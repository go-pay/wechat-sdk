package mini

import (
	"context"

	"github.com/go-pay/wechat-sdk/pkg/bm"
	"github.com/go-pay/wechat-sdk/pkg/util"
)

// CSMessageGetTempMedia 获取客服消息内的临时素材
//	mediaId：媒体文件 ID
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.getTempMedia.html
func (s *SDK) CSMessageGetTempMedia(c context.Context, mediaId string) (media []byte, err error) {
	path := "/cgi-bin/media/get?access_token=" + s.accessToken + "&media_id=" + mediaId
	media, err = s.doRequestGetByte(c, path)
	if err != nil {
		return nil, err
	}
	return
}

// CSMessageSend 发送客服消息给用户
//	toUser：小程序用户的 OpenID
//	msgType：消息类型，枚举值：mini.MsgTypeText、mini.MsgTypeImage、mini.MsgTypeLink、mini.MsgTypeMiniPage
//	msgValue：对应 msgType 的value值，BodyMap key-value 格式传入
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
func (s *SDK) CSMessageSend(c context.Context, toUser string, msgType MsgType, msgValue bm.BodyMap) (ec *ErrorCode, err error) {
	path := "/cgi-bin/message/custom/send?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("touser", toUser)
	switch msgType {
	case MsgTypeText:
		body.Set("msgtype", "text").
			SetBodyMap("text", func(bm bm.BodyMap) {
				bm.Set("content", msgValue.GetString("content"))
			})
	case MsgTypeImage:
		body.Set("msgtype", "image").
			SetBodyMap("text", func(bm bm.BodyMap) {
				bm.Set("media_id", msgValue.GetString("media_id"))
			})
	case MsgTypeLink:
		body.Set("msgtype", "link").
			SetBodyMap("text", func(bm bm.BodyMap) {
				bm.Set("title", msgValue.GetString("title")).
					Set("description", msgValue.GetString("description")).
					Set("url", msgValue.GetString("url")).
					Set("thumb_url", msgValue.GetString("thumb_url"))
			})
	case MsgTypeMiniPage:
		body.Set("msgtype", "miniprogrampage").
			SetBodyMap("text", func(bm bm.BodyMap) {
				bm.Set("title", msgValue.GetString("title")).
					Set("pagepath", msgValue.GetString("pagepath")).
					Set("thumb_media_id", msgValue.GetString("thumb_media_id"))
			})
	}
	ec = &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return nil, err
	}
	return
}

// CSMessageSetTyping 下发客服当前输入状态给用户
//	toUser：小程序用户的 OpenID
//	typingStatus：枚举值：mini.TypingTyping、mini.TypingCancel
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html
func (s *SDK) CSMessageSetTyping(c context.Context, toUser string, typingStatus TypingStatus) (ec *ErrorCode, err error) {
	path := "/cgi-bin/message/custom/typing?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("touser", toUser)
	switch typingStatus {
	case TypingTyping:
		body.Set("command", "Typing")
	case TypingCancel:
		body.Set("command", "CancelTyping")
	}
	ec = &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return nil, err
	}
	return
}

// CSMessageUploadTempMedia 把媒体文件上传到微信服务器
//	注意：目前仅支持图片，用于发送客服消息或被动回复用户消息。
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func (s *SDK) CSMessageUploadTempMedia(c context.Context, img *util.File) (media *UploadTempMedia, err error) {
	path := "/cgi-bin/media/upload?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("type", "image").
		SetFormFile("media", img)
	media = &UploadTempMedia{}
	if err = s.doRequestPostFile(c, path, body, media); err != nil {
		return nil, err
	}
	return
}
