package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// InvokeService 调用服务平台提供的服务
// 注意：errcode = 0 为成功
// service：服务ID
// api：接口名
// data：业务数据，JSON 字符串
// clientMsgId：随机字符串ID，用于唯一标识请求
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/service-market/invokeService.html
func (s *SDK) InvokeService(c context.Context, service, api, data, clientMsgId string) (result *InvokeServiceRsp, err error) {
	path := "/wxa/servicemarket?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("service", service).
		Set("api", api).
		Set("data", data).
		Set("client_msg_id", clientMsgId)

	result = &InvokeServiceRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
