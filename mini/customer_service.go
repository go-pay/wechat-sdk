package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/wechat-sdk/pkg/bmap"
	"github.com/go-pay/wechat-sdk/pkg/util"
)

// CSMessageGetTempMedia 获取客服消息内的临时素材
// mediaId：媒体文件 ID
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/kf-mgnt/kf-message/getTempMedia.html
func (s *SDK) CSMessageGetTempMedia(c context.Context, mediaId string) (media []byte, err error) {
	path := "/cgi-bin/media/get?access_token=" + s.accessToken + "&media_id=" + mediaId
	media, err = s.doRequestGetByte(c, path)
	if err != nil {
		return nil, err
	}
	return
}

// CSMessageSend 发送客服消息
// 注意：errcode = 0 为成功
// toUser：小程序用户的 OpenID
// msgType：消息类型，枚举值：mini.MsgTypeText、mini.MsgTypeImage、mini.MsgTypeLink、mini.MsgTypeMiniPage
// msgValue：对应 msgType 的value值，BodyMap key-value 格式传入
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/kf-mgnt/kf-message/sendCustomMessage.html
func (s *SDK) CSMessageSend(c context.Context, toUser string, msgType MsgType, msgValue bmap.BodyMap) (err error) {
	path := "/cgi-bin/message/custom/send?access_token=" + s.accessToken
	body := make(bmap.BodyMap)
	body.Set("touser", toUser)
	switch msgType {
	case MsgTypeText:
		body.Set("msgtype", "text").
			Set("text", msgValue)
	case MsgTypeImage:
		body.Set("msgtype", "image").
			Set("text", msgValue)
	case MsgTypeLink:
		body.Set("msgtype", "link").
			Set("text", msgValue)
	case MsgTypeMiniPage:
		body.Set("msgtype", "miniprogrampage").
			Set("text", msgValue)
	}
	ec := &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// CSMessageSetTyping 下发客服当前输入状态给用户
// 注意：errcode = 0 为成功
// toUser：小程序用户的 OpenID
// typingStatus：枚举值：mini.TypingTyping、mini.TypingCancel
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/kf-mgnt/kf-message/setTyping.html
func (s *SDK) CSMessageSetTyping(c context.Context, toUser string, typingStatus TypingStatus) (err error) {
	path := "/cgi-bin/message/custom/typing?access_token=" + s.accessToken
	body := make(bmap.BodyMap)
	body.Set("touser", toUser)
	switch typingStatus {
	case TypingTyping:
		body.Set("command", "Typing")
	case TypingCancel:
		body.Set("command", "CancelTyping")
	}
	ec := &ErrorCode{}
	if err = s.doRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// CSMessageUploadTempMedia 新增图片素材
// 注意：errcode = 0 为成功
// 注意：目前仅支持图片，用于发送客服消息或被动回复用户消息。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/kf-mgnt/kf-message/uploadTempMedia.html
func (s *SDK) CSMessageUploadTempMedia(c context.Context, img *util.File) (media *UploadTempMedia, err error) {
	path := "/cgi-bin/media/upload?access_token=" + s.accessToken
	body := make(bmap.BodyMap)
	body.Set("type", "image").
		SetFormFile("media", img)
	media = &UploadTempMedia{}
	if err = s.doRequestPostFile(c, path, body, media); err != nil {
		return nil, err
	}
	if media.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", media.Errcode, media.Errmsg)
	}
	return media, nil
}
