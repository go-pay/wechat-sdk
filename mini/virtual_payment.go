package mini

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// XpayQueryUserBalance 查询代币余额
// 注意：errcode = 0 为成功
// signature：用户态签名
// paySig：支付签名
// body 参数说明：
//
//	openid：用户openid
//	env：0-正式环境，1-沙箱环境
//	user_ip：用户IP
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_user_balance.html
func (s *SDK) XpayQueryUserBalance(c context.Context, signature, paySig string, body bm.BodyMap) (result *XpayQueryUserBalanceRsp, err error) {
	path := "/xpay/query_user_balance?access_token=" + s.accessToken + "&signature=" + signature + "&pay_sig=" + paySig
	result = &XpayQueryUserBalanceRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayCurrencyPay 扣减代币
// 注意：errcode = 0 为成功
// signature：用户态签名
// paySig：支付签名
// body 参数说明：
//
//	openid：用户openid
//	env：0-正式环境，1-沙箱环境
//	user_ip：用户IP
//	amount：扣减代币数量
//	order_id：订单号
//	payitem：道具信息，记录在账户流水中
//	remark：备注
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_currency_pay.html
func (s *SDK) XpayCurrencyPay(c context.Context, signature, paySig string, body bm.BodyMap) (result *XpayCurrencyPayRsp, err error) {
	path := "/xpay/currency_pay?access_token=" + s.accessToken + "&signature=" + signature + "&pay_sig=" + paySig
	result = &XpayCurrencyPayRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayQueryOrder 查询创建的订单
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	openid：用户openid
//	env：0-正式环境，1-沙箱环境
//	order_id：创建的订单号（与wx_order_id二选一）
//	wx_order_id：微信内部订单号（与order_id二选一）
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_order.html
func (s *SDK) XpayQueryOrder(c context.Context, paySig string, body bm.BodyMap) (result *XpayQueryOrderRsp, err error) {
	path := "/xpay/query_order?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayQueryOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayCancelCurrencyPay 代币支付退款
// 注意：errcode = 0 为成功
// signature：用户态签名
// paySig：支付签名
// body 参数说明：
//
//	openid：用户openid
//	env：0-正式环境，1-沙箱环境
//	user_ip：用户IP
//	pay_order_id：原currency_pay调用的订单号
//	order_id：退款订单号
//	amount：退款金额
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_cancel_currency_pay.html
func (s *SDK) XpayCancelCurrencyPay(c context.Context, signature, paySig string, body bm.BodyMap) (result *XpayCancelCurrencyPayRsp, err error) {
	path := "/xpay/cancel_currency_pay?access_token=" + s.accessToken + "&signature=" + signature + "&pay_sig=" + paySig
	result = &XpayCancelCurrencyPayRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayNotifyProvideGoods 通知已发货完成（只能通知现金单）
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	order_id：商户订单号（与wx_order_id二选一）
//	wx_order_id：微信内部订单号（与order_id二选一）
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_notify_provide_goods.html
func (s *SDK) XpayNotifyProvideGoods(c context.Context, body bm.BodyMap) (err error) {
	path := "/xpay/notify_provide_goods?access_token=" + s.accessToken
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpayPresentCurrency 代币赠送
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	openid：用户openid
//	env：0-正式环境，1-沙箱环境
//	order_id：赠送订单号
//	amount：赠送数量
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_present_currency.html
func (s *SDK) XpayPresentCurrency(c context.Context, body bm.BodyMap) (result *XpayPresentCurrencyRsp, err error) {
	path := "/xpay/present_currency?access_token=" + s.accessToken
	result = &XpayPresentCurrencyRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayDownloadBill 下载小程序账单
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	begin_ds：开始日期（格式：20230801）
//	end_ds：结束日期（格式：20230810）
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_download_bill.html
func (s *SDK) XpayDownloadBill(c context.Context, paySig string, body bm.BodyMap) (result *XpayDownloadBillRsp, err error) {
	path := "/xpay/download_bill?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayDownloadBillRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayRefundOrder 启动订单退款任务
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	openid：下单时的用户openid
//	order_id：商户订单号（与wx_order_id二选一）
//	wx_order_id：微信侧订单号（与order_id二选一）
//	refund_order_id：退款订单号，长度[8,32]，仅支持字母、数字、'_'、'-'
//	left_fee：剩余可退金额（分）
//	refund_fee：本次退款金额（分），取值(0, left_fee]
//	biz_meta：商户自定义数据，长度[0,1024]
//	refund_reason：退款原因代码：0-无描述，1-商品问题，2-售后，3-用户要求，4-价格，5-其他
//	req_from：来源代码：1-客服，2-用户发起，3-其他
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_refund_order.html
func (s *SDK) XpayRefundOrder(c context.Context, paySig string, body bm.BodyMap) (result *XpayRefundOrderRsp, err error) {
	path := "/xpay/refund_order?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayRefundOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayCreateWithdrawOrder 创建提现单
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	withdraw_no：提现单号，长度[8,32]，仅支持字母、数字、'_'、'-'
//	withdraw_amount：提现金额（单位元，如0.01表示1分钱），允许不传表示全额提现
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_create_withdraw_order.html
func (s *SDK) XpayCreateWithdrawOrder(c context.Context, paySig string, body bm.BodyMap) (result *XpayCreateWithdrawOrderRsp, err error) {
	path := "/xpay/create_withdraw_order?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayCreateWithdrawOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayQueryWithdrawOrder 查询提现单
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	withdraw_no：提现单号
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_withdraw_order.html
func (s *SDK) XpayQueryWithdrawOrder(c context.Context, paySig string, body bm.BodyMap) (result *XpayQueryWithdrawOrderRsp, err error) {
	path := "/xpay/query_withdraw_order?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayQueryWithdrawOrderRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayStartUploadGoods 批量上传道具
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	upload_item：道具列表，每个元素包含 id、name、price（分）、remark、item_url
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_start_upload_goods.html
func (s *SDK) XpayStartUploadGoods(c context.Context, paySig string, body bm.BodyMap) (err error) {
	path := "/xpay/start_upload_goods?access_token=" + s.accessToken + "&pay_sig=" + paySig
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpayQueryUploadGoods 查询批量上传道具任务
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_upload_goods.html
func (s *SDK) XpayQueryUploadGoods(c context.Context, paySig string, body bm.BodyMap) (result *XpayQueryUploadGoodsRsp, err error) {
	path := "/xpay/query_upload_goods?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayQueryUploadGoodsRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayStartPublishGoods 启动批量发布道具任务
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	publish_item：发布道具列表，每个元素包含 id
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_start_publish_goods.html
func (s *SDK) XpayStartPublishGoods(c context.Context, paySig string, body bm.BodyMap) (err error) {
	path := "/xpay/start_publish_goods?access_token=" + s.accessToken + "&pay_sig=" + paySig
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpayQueryPublishGoods 查询批量发布道具任务
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_publish_goods.html
func (s *SDK) XpayQueryPublishGoods(c context.Context, paySig string, body bm.BodyMap) (result *XpayQueryPublishGoodsRsp, err error) {
	path := "/xpay/query_publish_goods?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayQueryPublishGoodsRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayQueryBizBalance 查询商家账户可提现余额
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_biz_balance.html
func (s *SDK) XpayQueryBizBalance(c context.Context, paySig string, body bm.BodyMap) (result *XpayQueryBizBalanceRsp, err error) {
	path := "/xpay/query_biz_balance?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayQueryBizBalanceRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayQueryTransferAccount 查询广告金充值账户
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_transfer_account.html
func (s *SDK) XpayQueryTransferAccount(c context.Context, body bm.BodyMap) (result *XpayQueryTransferAccountRsp, err error) {
	path := "/xpay/query_transfer_account?access_token=" + s.accessToken
	result = &XpayQueryTransferAccountRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayQueryAdverFunds 查询广告金发放记录
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	page：查询页码，最小为1
//	page_size：每页条数
//	filter：查询过滤条件（settle_begin、settle_end、fund_type）
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_adver_funds.html
func (s *SDK) XpayQueryAdverFunds(c context.Context, body bm.BodyMap) (result *XpayQueryAdverFundsRsp, err error) {
	path := "/xpay/query_adver_funds?access_token=" + s.accessToken
	result = &XpayQueryAdverFundsRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayCreateFundsBill 充值广告金
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	transfer_amount：充值金额（分）
//	transfer_account_uid：充值账户UID
//	transfer_account_name：充值账户名称
//	transfer_account_agency_id：商户服务商账户ID
//	request_id：唯一请求标识（最长1024字符）
//	settle_begin：结算周期开始时间（unix时间戳秒）
//	settle_end：结算周期结束时间（unix时间戳秒）
//	env：0-正式环境，1-沙箱环境
//	authorize_advertise：广告数据授权：0-否，1-是
//	fund_type：广告金原因：0-通用赠送，1-广告激励，2-定向激励
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_create_funds_bill.html
func (s *SDK) XpayCreateFundsBill(c context.Context, body bm.BodyMap) (result *XpayCreateFundsBillRsp, err error) {
	path := "/xpay/create_funds_bill?access_token=" + s.accessToken
	result = &XpayCreateFundsBillRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayBindTransferAccount 绑定广告金充值账户
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	transfer_account_uid：充值账户UID
//	transfer_account_org_name：充值账户组织名称
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_bind_transfer_accout.html
func (s *SDK) XpayBindTransferAccount(c context.Context, body bm.BodyMap) (err error) {
	path := "/xpay/bind_transfer_accout?access_token=" + s.accessToken
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpayQueryFundsBill 查询广告金充值记录
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	page：查询页码，最小为1
//	page_size：每页条数
//	filter：查询过滤条件（oper_time_begin、oper_time_end、bill_id、request_id）
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_funds_bill.html
func (s *SDK) XpayQueryFundsBill(c context.Context, body bm.BodyMap) (result *XpayQueryFundsBillRsp, err error) {
	path := "/xpay/query_funds_bill?access_token=" + s.accessToken
	result = &XpayQueryFundsBillRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayQueryRecoverBill 查询广告金回收记录
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	page：查询页码，最小为1
//	page_size：每页条数
//	filter：查询过滤条件（recover_time_begin、recover_time_end、bill_id）
//	env：0-正式环境，1-沙箱环境
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_recover_bill.html
func (s *SDK) XpayQueryRecoverBill(c context.Context, body bm.BodyMap) (result *XpayQueryRecoverBillRsp, err error) {
	path := "/xpay/query_recover_bill?access_token=" + s.accessToken
	result = &XpayQueryRecoverBillRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayQuerySubscribeContract 查询签约关系
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	openid：用户openid
//	product_id：道具ID（需为订阅类型道具）
//	out_contract_code：签约时传入的协议号
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_query_subscribe_contract.html
func (s *SDK) XpayQuerySubscribeContract(c context.Context, body bm.BodyMap) (result *XpayQuerySubscribeContractRsp, err error) {
	path := "/xpay/query_subscribe_contract?access_token=" + s.accessToken
	result = &XpayQuerySubscribeContractRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpaySendSubscribePrePayment 预通知扣款
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	openid：用户openid
//	deduct_price：扣款金额（分），取值[100, 道具价格]
//	product_id：道具ID（需为订阅类型道具）
//	out_contract_code：签约时传入的协议号
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_send_subscribe_pre_payment.html
func (s *SDK) XpaySendSubscribePrePayment(c context.Context, body bm.BodyMap) (err error) {
	path := "/xpay/send_subscribe_pre_payment?access_token=" + s.accessToken
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpaySubmitSubscribePayOrder 发起订阅扣款
// 注意：errcode = 0 为成功，受理成功不代表扣款成功，需通过回调确认
// body 参数说明：
//
//	openid：用户openid
//	offer_id：在平台注册的应用ID
//	buy_quantity：购买数量（填1）
//	env：环境配置（仅支持0-正式环境）
//	currency_type：币种（填CNY）
//	product_id：订阅道具ID
//	deduct_price：扣款金额（分）
//	order_id：业务订单号，长度[8,32]，仅支持字母、数字、'_'、'-'
//	attach：透传数据，用于回调通知
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_submit_subscribe_pay_order.html
func (s *SDK) XpaySubmitSubscribePayOrder(c context.Context, body bm.BodyMap) (err error) {
	path := "/xpay/submit_subscribe_pay_order?access_token=" + s.accessToken
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpayCancelSubscribeContract 商家解约
// 注意：errcode = 0 为成功
// body 参数说明：
//
//	openid：用户openid
//	termination_reason：解约原因
//	product_id：道具ID（需为订阅类型道具）
//	out_contract_code：签约时传入的协议号
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_cancel_subscribe_contract.html
func (s *SDK) XpayCancelSubscribeContract(c context.Context, body bm.BodyMap) (err error) {
	path := "/xpay/cancel_subscribe_contract?access_token=" + s.accessToken
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpayGetComplaintList 获取投诉列表
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//	begin_date：开始日期（格式：yyyy-mm-dd）
//	end_date：结束日期（格式：yyyy-mm-dd）
//	offset：分页偏移量（从0开始）
//	limit：最大返回条数
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_get_complaint_list.html
func (s *SDK) XpayGetComplaintList(c context.Context, paySig string, body bm.BodyMap) (result *XpayGetComplaintListRsp, err error) {
	path := "/xpay/get_complaint_list?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayGetComplaintListRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayGetComplaintDetail 获取投诉详情
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//	complaint_id：投诉ID
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_get_complaint_detail.html
func (s *SDK) XpayGetComplaintDetail(c context.Context, paySig string, body bm.BodyMap) (result *XpayGetComplaintDetailRsp, err error) {
	path := "/xpay/get_complaint_detail?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayGetComplaintDetailRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayGetNegotiationHistory 获取协商历史
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//	complaint_id：投诉ID
//	offset：分页偏移量（从0开始）
//	limit：最大返回条数
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_get_negotiation_history.html
func (s *SDK) XpayGetNegotiationHistory(c context.Context, paySig string, body bm.BodyMap) (result *XpayGetNegotiationHistoryRsp, err error) {
	path := "/xpay/get_negotiation_history?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayGetNegotiationHistoryRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayResponseComplaint 回复用户投诉
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//	complaint_id：投诉ID
//	response_content：回复内容
//	response_images：图片文件ID数组（通过 XpayUploadVpFile 上传获得）
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_response_complaint.html
func (s *SDK) XpayResponseComplaint(c context.Context, paySig string, body bm.BodyMap) (err error) {
	path := "/xpay/response_complaint?access_token=" + s.accessToken + "&pay_sig=" + paySig
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpayCompleteComplaint 完成投诉处理
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//	complaint_id：投诉ID
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_complete_complaint.html
func (s *SDK) XpayCompleteComplaint(c context.Context, paySig string, body bm.BodyMap) (err error) {
	path := "/xpay/complete_complaint?access_token=" + s.accessToken + "&pay_sig=" + paySig
	ec := &ErrorCode{}
	if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
		return err
	}
	if ec.Errcode != Success {
		return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return nil
}

// XpayUploadVpFile 上传媒体文件
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//	base64_img：Base64编码的图片内容（最大1MB）
//	img_url：图片URL（最大2MB，优先级高于base64_img）
//	file_name：图片文件名
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_upload_vp_file.html
func (s *SDK) XpayUploadVpFile(c context.Context, paySig string, body bm.BodyMap) (result *XpayUploadVpFileRsp, err error) {
	path := "/xpay/upload_vp_file?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayUploadVpFileRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// XpayGetUploadFileSign 获取微信支付反馈投诉图片的签名头部
// 注意：errcode = 0 为成功
// paySig：支付签名
// body 参数说明：
//
//	env：0-正式环境，1-沙箱环境
//	wxpay_url：微信支付图片URL（格式：https://api.mch.weixin.qq.com/v3/merchant-service/images/{xxx}）
//	convert_cos：是否转换为COS存储
//	complaint_id：关联的投诉ID
//
// 文档：https://developers.weixin.qq.com/miniprogram/dev/server/API/VirtualPayment/api_get_upload_file_sign.html
func (s *SDK) XpayGetUploadFileSign(c context.Context, paySig string, body bm.BodyMap) (result *XpayGetUploadFileSignRsp, err error) {
	path := "/xpay/get_upload_file_sign?access_token=" + s.accessToken + "&pay_sig=" + paySig
	result = &XpayGetUploadFileSignRsp{}
	if _, err = s.DoRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
