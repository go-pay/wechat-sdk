package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// GenerateScheme 获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景
// 注意：通过该接口生成的小程序 scheme 码，永久有效，有数量限制
// jumpWxa：跳转到的目标小程序信息，BodyMap 格式：{"path": "pages/index/index", "query": "a=1&b=2", "env_version": "release"}
// isExpire：生成的 scheme 码类型，到期失效：true，永久有效：false，默认 false
// expireType：到期失效的 scheme 码的失效类型，失效时间：0，失效间隔天数：1
// expireTime：到期失效的 scheme 码的失效时间，为 Unix 时间戳。生成的到期失效 scheme 码在该时间前有效。最长有效期为1年。生成到期失效的 scheme 时必填。
// expireInterval：到期失效的 scheme 码的失效间隔天数。生成的到期失效 scheme 码在该间隔时间到达前有效。最长间隔天数为365天。is_expire 为 true 且 expire_type 为 1 时必填。
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/generateScheme.html
func (s *SDK) GenerateScheme(c context.Context, jumpWxa bm.BodyMap, isExpire bool, expireType, expireTime, expireInterval int) (result *GenerateSchemeRsp, err error) {
	path := "/wxa/generatescheme?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if jumpWxa != nil {
		body.Set("jump_wxa", jumpWxa)
	}
	body.Set("is_expire", isExpire)
	if isExpire {
		body.Set("expire_type", expireType)
		if expireType == 0 && expireTime > 0 {
			body.Set("expire_time", expireTime)
		}
		if expireType == 1 && expireInterval > 0 {
			body.Set("expire_interval", expireInterval)
		}
	}

	result = &GenerateSchemeRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// QueryScheme 查询小程序 scheme 码
// scheme：小程序 scheme 码
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/queryScheme.html
func (s *SDK) QueryScheme(c context.Context, scheme string) (result *QuerySchemeRsp, err error) {
	path := "/wxa/queryscheme?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("scheme", scheme)

	result = &QuerySchemeRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GenerateUrlLink 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景
// path：通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query。path 为空时会跳转小程序主页
// query：通过 URL Link 进入小程序时的query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：`!#$&'()*+,/:;=?@-._~%“
// isExpire：生成的 URL Link 类型，到期失效：true，永久有效：false，默认 false
// expireType：小程序 URL Link 失效类型，失效时间：0，失效间隔天数：1
// expireTime：到期失效的 URL Link 的失效时间，为 Unix 时间戳。生成的到期失效 URL Link 在该时间前有效。最长有效期为1年。生成到期失效的 URL Link 时必填
// expireInterval：到期失效的 URL Link 的失效间隔天数。生成的到期失效 URL Link 在该间隔时间到达前有效。最长间隔天数为365天。is_expire 为 true 且 expire_type 为 1 时必填
// envVersion：要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-link/generateUrlLink.html
func (s *SDK) GenerateUrlLink(c context.Context, path, query string, isExpire bool, expireType, expireTime, expireInterval int, envVersion string) (result *GenerateUrlLinkRsp, err error) {
	apiPath := "/wxa/generate_urllink?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	if path != "" {
		body.Set("path", path)
	}
	if query != "" {
		body.Set("query", query)
	}
	body.Set("is_expire", isExpire)
	if isExpire {
		body.Set("expire_type", expireType)
		if expireType == 0 && expireTime > 0 {
			body.Set("expire_time", expireTime)
		}
		if expireType == 1 && expireInterval > 0 {
			body.Set("expire_interval", expireInterval)
		}
	}
	if envVersion != "" {
		body.Set("env_version", envVersion)
	}

	result = &GenerateUrlLinkRsp{}
	if _, err = s.DoRequestPost(c, apiPath, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// QueryUrlLink 查询小程序 URL Link
// urlLink：小程序 URL Link
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-link/queryUrlLink.html
func (s *SDK) QueryUrlLink(c context.Context, urlLink string) (result *QueryUrlLinkRsp, err error) {
	path := "/wxa/query_urllink?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("url_link", urlLink)

	result = &QueryUrlLinkRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GenerateShortLink 获取小程序 Short Link，适用于微信内拉起小程序的业务场景
// pageUrl：通过 Short Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，可携带 query，最大1024个字符
// pageTitle：页面标题，不能包含违法信息，超过20字符会用... 截断代替
// isPermanent：生成的 Short Link 类型，短期有效：false，永久有效：true，默认 false
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/short-link/generateShortLink.html
func (s *SDK) GenerateShortLink(c context.Context, pageUrl, pageTitle string, isPermanent bool) (result *GenerateShortLinkRsp, err error) {
	path := "/wxa/genwxashortlink?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("page_url", pageUrl)
	if pageTitle != "" {
		body.Set("page_title", pageTitle)
	}
	body.Set("is_permanent", isPermanent)

	result = &GenerateShortLinkRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
