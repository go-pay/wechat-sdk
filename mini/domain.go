package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// GetDomainInfo 获取域名配置
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/domain-management/getDomainInfo.html
func (s *SDK) GetDomainInfo(c context.Context) (result *GetDomainInfoRsp, err error) {
	path := "/wxa/getwxadevinfo?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &GetDomainInfoRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ModifyDomain 修改服务器域名
// 注意：errcode = 0 为成功
// action：操作类型，add添加，delete删除，set覆盖，get获取
// requestdomain：request合法域名，当action为get时不需要此字段
// wsrequestdomain：socket合法域名，当action为get时不需要此字段
// uploaddomain：uploadFile合法域名，当action为get时不需要此字段
// downloaddomain：downloadFile合法域名，当action为get时不需要此字段
// udpdomain：udp合法域名，当action为get时不需要此字段
// tcpdomain：tcp合法域名，当action为get时不需要此字段
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/domain-management/modifyDomain.html
func (s *SDK) ModifyDomain(c context.Context, action string, requestdomain, wsrequestdomain, uploaddomain, downloaddomain, udpdomain, tcpdomain []string) (result *ModifyDomainRsp, err error) {
	path := "/wxa/modify_domain?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("action", action)

	if action != "get" {
		if len(requestdomain) > 0 {
			body.Set("requestdomain", requestdomain)
		}
		if len(wsrequestdomain) > 0 {
			body.Set("wsrequestdomain", wsrequestdomain)
		}
		if len(uploaddomain) > 0 {
			body.Set("uploaddomain", uploaddomain)
		}
		if len(downloaddomain) > 0 {
			body.Set("downloaddomain", downloaddomain)
		}
		if len(udpdomain) > 0 {
			body.Set("udpdomain", udpdomain)
		}
		if len(tcpdomain) > 0 {
			body.Set("tcpdomain", tcpdomain)
		}
	}

	result = &ModifyDomainRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// SetWebviewDomain 设置业务域名
// 注意：errcode = 0 为成功
// action：操作类型，add添加，delete删除，set覆盖，get获取
// webviewdomain：业务域名，当action为get时不需要此字段
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/domain-management/setWebviewDomain.html
func (s *SDK) SetWebviewDomain(c context.Context, action string, webviewdomain []string) (result *SetWebviewDomainRsp, err error) {
	path := "/wxa/setwebviewdomain?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("action", action)

	if action != "get" && len(webviewdomain) > 0 {
		body.Set("webviewdomain", webviewdomain)
	}

	result = &SetWebviewDomainRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetQrcodeJumppublish 获取小程序码扫码打开的页面
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code-jump/getQRCodeJumpPublish.html
func (s *SDK) GetQrcodeJumppublish(c context.Context) (result *GetQrcodeJumppublishRsp, err error) {
	path := "/wxa/get_qrcode_jumppublish?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &GetQrcodeJumppublishRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// SetQrcodeJump 设置小程序码扫码打开的页面
// 注意：errcode = 0 为成功
// prefix：二维码规则
// permitSubRule：是否独占符合二维码前缀匹配规则的所有子规则，1为独占，0为不独占
// path：小程序功能页面
// openVersion：测试范围，1为开发版，2为体验版，3为正式版
// debugUrl：测试链接（选填）
// isEdit：编辑标志位，0表示新增，1表示修改
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code-jump/setQRCodeJump.html
func (s *SDK) SetQrcodeJump(c context.Context, prefix string, permitSubRule int, path string, openVersion int, debugUrl []string, isEdit int) (result *SetQrcodeJumpRsp, err error) {
	apiPath := "/wxa/qrcodejumpadd?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("prefix", prefix).
		Set("permit_sub_rule", permitSubRule).
		Set("path", path).
		Set("open_version", openVersion).
		Set("is_edit", isEdit)

	if len(debugUrl) > 0 {
		body.Set("debug_url", debugUrl)
	}

	result = &SetQrcodeJumpRsp{}
	if _, err = s.DoRequestPost(c, apiPath, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
