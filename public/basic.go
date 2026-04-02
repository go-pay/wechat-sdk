package public

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// GetApiDomainIp 获取微信服务器IP地址
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (s *SDK) GetApiDomainIp(c context.Context) (result *ApiDomainIpRsp, err error) {
	path := "/cgi-bin/get_api_domain_ip?access_token=" + s.accessToken

	result = &ApiDomainIpRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetCallbackIp 获取微信callback IP地址
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (s *SDK) GetCallbackIp(c context.Context) (result *CallbackIpRsp, err error) {
	path := "/cgi-bin/getcallbackip?access_token=" + s.accessToken

	result = &CallbackIpRsp{}
	if _, err = s.doRequestGet(c, path, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ClearQuota 清空API调用quota
// appid：公众号的appid
// 文档：https://developers.weixin.qq.com/doc/offiaccount/openApi/clear_quota.html
func (s *SDK) ClearQuota(c context.Context, appid string) (err error) {
	path := "/cgi-bin/clear_quota?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("appid", appid)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}

// GetApiQuota 查询API调用额度
// cgiPath：api的请求地址，例如"/cgi-bin/message/custom/send"
// 文档：https://developers.weixin.qq.com/doc/offiaccount/openApi/get_api_quota.html
func (s *SDK) GetApiQuota(c context.Context, cgiPath string) (result *ApiQuotaRsp, err error) {
	path := "/cgi-bin/openapi/quota/get?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("cgi_path", cgiPath)

	result = &ApiQuotaRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetRid 查询rid信息
// rid：调用接口报错返回的rid
// 文档：https://developers.weixin.qq.com/doc/offiaccount/openApi/get_rid_info.html
func (s *SDK) GetRid(c context.Context, rid string) (result *RidInfoRsp, err error) {
	path := "/cgi-bin/openapi/rid/get?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("rid", rid)

	result = &RidInfoRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ClearQuotaByAppSecret 使用AppSecret重置API调用次数
// appid：公众号的appid
// appsecret：公众号的appsecret
// 文档：https://developers.weixin.qq.com/doc/offiaccount/openApi/clear_quota_by_appsecret.html
func (s *SDK) ClearQuotaByAppSecret(c context.Context, appid, appsecret string) (err error) {
	path := "/cgi-bin/clear_quota/v2?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("appid", appid)
	body.Set("appsecret", appsecret)

	result := &ErrorCode{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return err
	}
	if result.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return nil
}
