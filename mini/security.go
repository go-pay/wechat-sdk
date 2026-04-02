package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// MsgSecCheck 检查一段文本是否含有违法违规内容
// 注意：errcode = 0 为成功
// content：需检测的文本内容，文本字数的上限为2500字，需使用UTF-8编码
// version：接口版本号，2.0版本为固定值2
// scene：场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
// openid：用户的openid（用户需在近两小时访问过小程序）
// title：文本标题，需使用UTF-8编码
// nickname：用户昵称，需使用UTF-8编码
// signature：个性签名，该参数仅在资料类场景有效(scene=1)，需使用UTF-8编码
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/msgSecCheck.html
func (s *SDK) MsgSecCheck(c context.Context, content string, version int, scene int, openid, title, nickname, signature string) (result *MsgSecCheckRsp, err error) {
	path := "/wxa/msg_sec_check?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("content", content).
		Set("version", version).
		Set("scene", scene).
		Set("openid", openid)

	if title != "" {
		body.Set("title", title)
	}
	if nickname != "" {
		body.Set("nickname", nickname)
	}
	if signature != "" {
		body.Set("signature", signature)
	}

	result = &MsgSecCheckRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ImgSecCheck 校验一张图片是否含有违法违规内容
// 注意：errcode = 0 为成功
// 注意：图片文件大小应小于 5MB
// mediaUrl：要检测的图片url，支持图片格式包括 jpg, jepg, png, bmp, gif（取首帧）
// mediaType：1:url；2:base64
// version：接口版本号，2.0版本为固定值2
// scene：场景枚举值（1.资料；2.评论；3.论坛；4.社交日志）
// openid：用户的openid（用户需在近两小时访问过小程序）
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/imgSecCheck.html
func (s *SDK) ImgSecCheck(c context.Context, mediaUrl string, mediaType, version, scene int, openid string) (result *ImgSecCheckRsp, err error) {
	path := "/wxa/img_sec_check?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("media_url", mediaUrl).
		Set("media_type", mediaType).
		Set("version", version).
		Set("scene", scene).
		Set("openid", openid)

	result = &ImgSecCheckRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
// 注意：errcode = 0 为成功
// 注意：音频文件大小不超过10MB，时长不超过5分钟；图片文件大小不超过5MB
// mediaUrl：要检测的多媒体url
// mediaType：1:音频;2:图片
// version：接口版本号，2.0版本为固定值2
// scene：场景枚举值（1.资料；2.评论；3.论坛；4.社交日志）
// openid：用户的openid（用户需在近两小时访问过小程序）
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/mediaCheckAsync.html
func (s *SDK) MediaCheckAsync(c context.Context, mediaUrl string, mediaType, version, scene int, openid string) (result *MediaCheckAsyncRsp, err error) {
	path := "/wxa/media_check_async?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("media_url", mediaUrl).
		Set("media_type", mediaType).
		Set("version", version).
		Set("scene", scene).
		Set("openid", openid)

	result = &MediaCheckAsyncRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
