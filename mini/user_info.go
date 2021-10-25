package mini

import (
	"context"
)

// GetPaidUnionid 用户支付完成后，获取该用户的 UnionId，无需用户授权
//	openid：支付用户唯一标识
//	transactionId：微信支付订单号
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (s *SDK) GetPaidUnionid(c context.Context, openid, transactionId string) (unionid *PaidUnionId, err error) {
	unionid = new(PaidUnionId)
	path := "/wxa/getpaidunionid?access_token=" + s.accessToken + "&openid=" + openid + "&transaction_id=" + transactionId
	if err = s.doRequestGet(c, path, unionid); err != nil {
		return nil, err
	}
	return
}

// GetPaidUnionidByTradeNo 用户支付完成后，获取该用户的 UnionId，无需用户授权
//	openid：支付用户唯一标识
//	mchid：微信支付商户号
//	tradeNo：微信支付商户订单号
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (s *SDK) GetPaidUnionidByTradeNo(c context.Context, openid, mchid, tradeNo string) (unionid *PaidUnionId, err error) {
	unionid = new(PaidUnionId)
	path := "/wxa/getpaidunionid?access_token=" + s.accessToken + "&openid=" + openid + "&mch_id=" + mchid + "&out_trade_no=" + tradeNo
	if err = s.doRequestGet(c, path, unionid); err != nil {
		return nil, err
	}
	return
}
