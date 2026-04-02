package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// ExpressAddOrder 生成运单
// 注意：errcode = 0 为成功
// addSource：订单来源，0为小程序订单，2为App或H5订单
// wxAppid：App或H5的appid，add_source=2时必填
// orderId：订单ID，需保证全局唯一
// openid：用户openid，当add_source=2时无需填写
// deliveryId：快递公司ID
// bizId：快递公司客户编码
// customRemark：快递备注信息
// sender：发件人信息，BodyMap 格式
// receiver：收件人信息，BodyMap 格式
// cargo：包裹信息，BodyMap 格式
// shop：商品信息，BodyMap 格式
// insured：保价信息，BodyMap 格式
// service：服务类型，BodyMap 格式
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_add_order.html
func (s *SDK) ExpressAddOrder(c context.Context, addSource int, wxAppid, orderId, openid, deliveryId, bizId, customRemark string, sender, receiver, cargo, shop, insured, service bm.BodyMap) (result *ExpressAddOrderRsp, err error) {
	path := "/cgi-bin/express/business/order/add?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("add_source", addSource).
		Set("order_id", orderId).
		Set("delivery_id", deliveryId).
		Set("biz_id", bizId).
		Set("sender", sender).
		Set("receiver", receiver).
		Set("cargo", cargo).
		Set("shop", shop)

	if wxAppid != "" {
		body.Set("wx_appid", wxAppid)
	}
	if openid != "" {
		body.Set("openid", openid)
	}
	if customRemark != "" {
		body.Set("custom_remark", customRemark)
	}
	if insured != nil {
		body.Set("insured", insured)
	}
	if service != nil {
		body.Set("service", service)
	}

	result = &ExpressAddOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ExpressCancelOrder 取消运单
// 注意：errcode = 0 为成功
// orderId：订单ID
// openid：用户openid
// deliveryId：快递公司ID
// waybillId：运单号
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_cancel_order.html
func (s *SDK) ExpressCancelOrder(c context.Context, orderId, openid, deliveryId, waybillId string) (result *ExpressCancelOrderRsp, err error) {
	path := "/cgi-bin/express/business/order/cancel?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("order_id", orderId).
		Set("openid", openid).
		Set("delivery_id", deliveryId).
		Set("waybill_id", waybillId)

	result = &ExpressCancelOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ExpressGetAllAccount 获取所有绑定的物流账号
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_get_all_account.html
func (s *SDK) ExpressGetAllAccount(c context.Context) (result *ExpressGetAllAccountRsp, err error) {
	path := "/cgi-bin/express/business/account/getall?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &ExpressGetAllAccountRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ExpressGetQuota 获取电子面单余额
// 注意：errcode = 0 为成功
// deliveryId：快递公司ID
// bizId：快递公司客户编码
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_get_quota.html
func (s *SDK) ExpressGetQuota(c context.Context, deliveryId, bizId string) (result *ExpressGetQuotaRsp, err error) {
	path := "/cgi-bin/express/business/quota/get?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("delivery_id", deliveryId).
		Set("biz_id", bizId)

	result = &ExpressGetQuotaRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ExpressGetPath 查询运单轨迹
// 注意：errcode = 0 为成功
// orderId：订单ID
// openid：用户openid
// deliveryId：快递公司ID
// waybillId：运单号
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_get_path.html
func (s *SDK) ExpressGetPath(c context.Context, orderId, openid, deliveryId, waybillId string) (result *ExpressGetPathRsp, err error) {
	path := "/cgi-bin/express/business/path/get?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("order_id", orderId).
		Set("openid", openid).
		Set("delivery_id", deliveryId).
		Set("waybill_id", waybillId)

	result = &ExpressGetPathRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ExpressGetAllDelivery 获取支持的快递公司列表
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_get_all_delivery.html
func (s *SDK) ExpressGetAllDelivery(c context.Context) (result *ExpressGetAllDeliveryRsp, err error) {
	path := "/cgi-bin/express/business/delivery/getall?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &ExpressGetAllDeliveryRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ExpressGetPrinter 获取打印员
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_get_printer.html
func (s *SDK) ExpressGetPrinter(c context.Context) (result *ExpressGetPrinterRsp, err error) {
	path := "/cgi-bin/express/business/printer/getall?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &ExpressGetPrinterRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ExpressUpdatePrinter 配置面单打印员
// 注意：errcode = 0 为成功
// updateType：更新类型，bind为绑定，unbind为解绑
// openid：打印员openid，数组格式
// tagidList：用户标签ID列表
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_update_printer.html
func (s *SDK) ExpressUpdatePrinter(c context.Context, updateType string, openid []string, tagidList string) (result *ExpressUpdatePrinterRsp, err error) {
	path := "/cgi-bin/express/business/printer/update?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("update_type", updateType)

	if len(openid) > 0 {
		body.Set("openid", openid)
	}
	if tagidList != "" {
		body.Set("tagid_list", tagidList)
	}

	result = &ExpressUpdatePrinterRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// ExpressGetContact 获取面单联系人信息
// 注意：errcode = 0 为成功
// waybillId：运单号
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_get_contact.html
func (s *SDK) ExpressGetContact(c context.Context, waybillId string) (result *ExpressGetContactRsp, err error) {
	path := "/cgi-bin/express/business/contact/get?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("waybill_id", waybillId)

	result = &ExpressGetContactRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
