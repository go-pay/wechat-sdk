package mini

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/go-pay/wechat-sdk/model"
	"github.com/go-pay/wechat-sdk/pkg/bm"
)

// GetPaidUnionid 用户支付完成后，获取该用户的 UnionId，无需用户授权
//	openid：支付用户唯一标识
//	transactionId：微信支付订单号
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (s *SDK) GetPaidUnionid(c context.Context, openid, transactionId string) (unionid *model.PaidUnionId, err error) {
	path := "/wxa/getpaidunionid?access_token=" + s.Conf.AccessToken + "&openid=" + openid + "&transaction_id=" + transactionId
	unionid = &model.PaidUnionId{}
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
func (s *SDK) GetPaidUnionidByTradeNo(c context.Context, openid, mchid, tradeNo string) (unionid *model.PaidUnionId, err error) {
	path := "/wxa/getpaidunionid?access_token=" + s.Conf.AccessToken + "&openid=" + openid + "&mch_id=" + mchid + "&out_trade_no=" + tradeNo
	unionid = &model.PaidUnionId{}
	if err = s.doRequestGet(c, path, unionid); err != nil {
		return nil, err
	}
	return
}

// CheckEncryptedData 检查加密信息是否由微信生成
//	encryptedData：加密数据，无需sha256操作
//	注意：（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.checkEncryptedData.html
func (s *SDK) CheckEncryptedData(c context.Context, encryptedData string) (result *model.CheckEncryptedResult, err error) {
	path := "/wxa/business/checkencryptedmsg?access_token=" + s.Conf.AccessToken
	h := sha256.New()
	h.Write([]byte(encryptedData))
	body := make(bm.BodyMap)
	body.Set("encrypted_msg_hash", hex.EncodeToString(h.Sum(nil)))
	result = &model.CheckEncryptedResult{}
	if err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	return
}
