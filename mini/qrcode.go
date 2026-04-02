package mini

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/bm"
	"github.com/go-pay/wechat-sdk"
)

// GetWxaCode 获取小程序码（适用于需要的码数量较少的业务场景）
// 注意：通过该接口生成的小程序码，永久有效，有数量限制
// path：扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
// width：二维码的宽度，单位 px。最小 280px，最大 1280px，默认 430px
// autoColor：自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// lineColor：auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
// isHyaline：是否需要透明底色，为 true 时，生成透明底色的小程序码
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getQRCode.html
func (s *SDK) GetWxaCode(c context.Context, path string, width int, autoColor, isHyaline bool, lineColor *LineColor) (qrcode []byte, err error) {
	apiPath := "/wxa/getwxacode?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("path", path)
	if width > 0 {
		body.Set("width", width)
	}
	body.Set("auto_color", autoColor)
	if lineColor != nil {
		body.Set("line_color", lineColor)
	}
	body.Set("is_hyaline", isHyaline)

	qrcode, err = s.doRequestPostForImage(c, apiPath, body)
	if err != nil {
		return nil, err
	}
	return qrcode, nil
}

// GetWxaCodeUnlimit 获取小程序码（适用于需要的码数量极多的业务场景）
// 注意：通过该接口生成的小程序码，永久有效，数量暂无限制
// scene：最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
// page：必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
// checkPath：检查page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）；为 false 时允许小程序未发布或者 page 不存在， 但page 有数量上限（60000个）请勿滥用
// envVersion：要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版
// width：二维码的宽度，单位 px。最小 280px，最大 1280px，默认 430px
// autoColor：自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// lineColor：auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
// isHyaline：是否需要透明底色，为 true 时，生成透明底色的小程序码
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
func (s *SDK) GetWxaCodeUnlimit(c context.Context, scene, page string, checkPath bool, envVersion string, width int, autoColor, isHyaline bool, lineColor *LineColor) (qrcode []byte, err error) {
	apiPath := "/wxa/getwxacodeunlimit?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("scene", scene)
	if page != "" {
		body.Set("page", page)
	}
	body.Set("check_path", checkPath)
	if envVersion != "" {
		body.Set("env_version", envVersion)
	}
	if width > 0 {
		body.Set("width", width)
	}
	body.Set("auto_color", autoColor)
	if lineColor != nil {
		body.Set("line_color", lineColor)
	}
	body.Set("is_hyaline", isHyaline)

	qrcode, err = s.doRequestPostForImage(c, apiPath, body)
	if err != nil {
		return nil, err
	}
	return qrcode, nil
}

// CreateWxaQRCode 获取小程序二维码（适用于需要的码数量较少的业务场景）
// 注意：通过该接口生成的小程序二维码，永久有效，有数量限制
// path：扫码进入的小程序页面路径，最大长度 128 字节，不能为空。对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
// width：二维码的宽度，单位 px。最小 280px，最大 1280px，默认 430px
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/createQRCode.html
func (s *SDK) CreateWxaQRCode(c context.Context, path string, width int) (qrcode []byte, err error) {
	apiPath := "/cgi-bin/wxaapp/createwxaqrcode?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("path", path)
	if width > 0 {
		body.Set("width", width)
	}

	qrcode, err = s.doRequestPostForImage(c, apiPath, body)
	if err != nil {
		return nil, err
	}
	return qrcode, nil
}

// doRequestPostForImage 处理返回图片的 POST 请求
func (s *SDK) doRequestPostForImage(c context.Context, path string, body bm.BodyMap) (image []byte, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_URI: %s", uri)
		s.logger.Debugf("Wechat_Mini_SDK_RequestBody: %s", body.JsonBody())
	}

	res, bs, err := s.hc.Req().Post(uri).SendBodyMap(body).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(POST, %s), err:%w", uri, err)
	}

	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_Response: [%d] -> %d bytes", res.StatusCode, len(bs))
	}

	// 判断返回的是图片还是错误信息
	contentType := res.Header.Get("Content-Type")
	if contentType == "image/jpeg" || contentType == "image/png" {
		return bs, nil
	}

	// 如果不是图片，尝试解析错误信息
	ec := &ErrorCode{}
	if err = json.Unmarshal(bs, ec); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if ec.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}

	return bs, nil
}
