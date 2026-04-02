package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// DeliveryGetAllAccount 拉取已绑定账号
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/immediate-delivery/delivery_get_all_account.html
func (s *SDK) DeliveryGetAllAccount(c context.Context) (result *DeliveryGetAllAccountRsp, err error) {
	path := "/cgi-bin/express/delivery/open_msg/get_all_account?access_token=" + s.accessToken
	body := make(bm.BodyMap)

	result = &DeliveryGetAllAccountRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DeliveryPreAddOrder 配送单预下单
// 注意：errcode = 0 为成功
// deliveryId：配送公司ID
// shopOrderId：商家订单号
// shopNo：商家门店编号
// deliverySign：配送签名
// orderInfo：订单信息，BodyMap 格式
// cargo：货物信息，BodyMap 格式
// sender：发件人信息，BodyMap 格式
// receiver：收件人信息，BodyMap 格式
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/immediate-delivery/delivery_pre_add_order.html
func (s *SDK) DeliveryPreAddOrder(c context.Context, deliveryId, shopOrderId, shopNo, deliverySign string, orderInfo, cargo, sender, receiver bm.BodyMap) (result *DeliveryPreAddOrderRsp, err error) {
	path := "/cgi-bin/express/delivery/open_msg/pre_add_order?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("delivery_id", deliveryId).
		Set("shoporder_id", shopOrderId).
		Set("shop_no", shopNo).
		Set("delivery_sign", deliverySign).
		Set("order_info", orderInfo).
		Set("cargo", cargo).
		Set("sender", sender).
		Set("receiver", receiver)

	result = &DeliveryPreAddOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DeliveryAddOrder 配送单下单
// 注意：errcode = 0 为成功
// deliveryId：配送公司ID
// shopOrderId：商家订单号
// shopNo：商家门店编号
// deliverySign：配送签名
// deliveryToken：预下单接口返回的token
// orderInfo：订单信息，BodyMap 格式
// cargo：货物信息，BodyMap 格式
// sender：发件人信息，BodyMap 格式
// receiver：收件人信息，BodyMap 格式
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/immediate-delivery/delivery_add_order.html
func (s *SDK) DeliveryAddOrder(c context.Context, deliveryId, shopOrderId, shopNo, deliverySign, deliveryToken string, orderInfo, cargo, sender, receiver bm.BodyMap) (result *DeliveryAddOrderRsp, err error) {
	path := "/cgi-bin/express/delivery/open_msg/add_order?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("delivery_id", deliveryId).
		Set("shoporder_id", shopOrderId).
		Set("shop_no", shopNo).
		Set("delivery_sign", deliverySign).
		Set("delivery_token", deliveryToken).
		Set("order_info", orderInfo).
		Set("cargo", cargo).
		Set("sender", sender).
		Set("receiver", receiver)

	result = &DeliveryAddOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DeliveryAddTips 配送单增加小费
// 注意：errcode = 0 为成功
// shopOrderId：商家订单号
// shopNo：商家门店编号
// deliverySign：配送签名
// waybillId：配送单号
// openid：下单用户的openid
// tips：小费金额(单位：元)
// remarkContent：备注内容
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/immediate-delivery/delivery_add_tips.html
func (s *SDK) DeliveryAddTips(c context.Context, shopOrderId, shopNo, deliverySign, waybillId, openid string, tips float64, remarkContent string) (result *DeliveryAddTipsRsp, err error) {
	path := "/cgi-bin/express/delivery/open_msg/add_tip?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("shoporder_id", shopOrderId).
		Set("shop_no", shopNo).
		Set("delivery_sign", deliverySign).
		Set("waybill_id", waybillId).
		Set("openid", openid).
		Set("tips", tips).
		Set("remark_content", remarkContent)

	result = &DeliveryAddTipsRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DeliveryCancelOrder 配送单取消
// 注意：errcode = 0 为成功
// shopOrderId：商家订单号
// shopNo：商家门店编号
// deliverySign：配送签名
// deliveryId：配送公司ID
// waybillId：配送单号
// cancelReasonId：取消原因ID
// cancelReason：取消原因
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/immediate-delivery/delivery_cancel_order.html
func (s *SDK) DeliveryCancelOrder(c context.Context, shopOrderId, shopNo, deliverySign, deliveryId, waybillId string, cancelReasonId int, cancelReason string) (result *DeliveryCancelOrderRsp, err error) {
	path := "/cgi-bin/express/delivery/open_msg/cancel_order?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("shoporder_id", shopOrderId).
		Set("shop_no", shopNo).
		Set("delivery_sign", deliverySign).
		Set("delivery_id", deliveryId).
		Set("waybill_id", waybillId).
		Set("cancel_reason_id", cancelReasonId).
		Set("cancel_reason", cancelReason)

	result = &DeliveryCancelOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DeliveryGetOrder 配送单查询
// 注意：errcode = 0 为成功
// shopOrderId：商家订单号
// shopNo：商家门店编号
// deliverySign：配送签名
// deliveryId：配送公司ID
// waybillId：配送单号
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/immediate-delivery/delivery_get_order.html
func (s *SDK) DeliveryGetOrder(c context.Context, shopOrderId, shopNo, deliverySign, deliveryId, waybillId string) (result *DeliveryGetOrderRsp, err error) {
	path := "/cgi-bin/express/delivery/open_msg/get_order?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("shoporder_id", shopOrderId).
		Set("shop_no", shopNo).
		Set("delivery_sign", deliverySign).
		Set("delivery_id", deliveryId).
		Set("waybill_id", waybillId)

	result = &DeliveryGetOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DeliveryMockUpdateOrder 模拟配送公司更新配送单状态（仅用于测试）
// 注意：errcode = 0 为成功
// shopOrderId：商家订单号
// shopNo：商家门店编号
// actionTime：状态变更时间
// orderStatus：配送状态
// actionMsg：附加信息
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/immediate-delivery/delivery_mock_update_order.html
func (s *SDK) DeliveryMockUpdateOrder(c context.Context, shopOrderId, shopNo string, actionTime int64, orderStatus int, actionMsg string) (result *DeliveryMockUpdateOrderRsp, err error) {
	path := "/cgi-bin/express/delivery/open_msg/mock_update_order?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("shoporder_id", shopOrderId).
		Set("shop_no", shopNo).
		Set("action_time", actionTime).
		Set("order_status", orderStatus).
		Set("action_msg", actionMsg)

	result = &DeliveryMockUpdateOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// DeliveryAbnormalConfirm 异常件退回商家商圈接口
// 注意：errcode = 0 为成功
// shopOrderId：商家订单号
// shopNo：商家门店编号
// deliverySign：配送签名
// deliveryId：配送公司ID
// waybillId：配送单号
// remark：备注
// 文档：https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/immediate-delivery/delivery_abnormal_confirm.html
func (s *SDK) DeliveryAbnormalConfirm(c context.Context, shopOrderId, shopNo, deliverySign, deliveryId, waybillId, remark string) (result *DeliveryAbnormalConfirmRsp, err error) {
	path := "/cgi-bin/express/delivery/open_msg/abnormal_confirm?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("shoporder_id", shopOrderId).
		Set("shop_no", shopNo).
		Set("delivery_sign", deliverySign).
		Set("delivery_id", deliveryId).
		Set("waybill_id", waybillId).
		Set("remark", remark)

	result = &DeliveryAbnormalConfirmRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
