package mini

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/go-pay/wechat-sdk/pkg/bmap"
)

// GetPluginOpenPid 获取插件用户openpid
// 注意：errcode = 0 为成功
// code：通过 wx.pluginLogin 获得的插件用户标志凭证 code，有效时间为5分钟，一个 code 只能获取一次 openpid
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/basic-info/getPluginOpenPId.html
func (s *SDK) GetPluginOpenPid(c context.Context, code string) (openpid *PluginOpenPid, err error) {
	path := "/wxa/getpluginopenpid?access_token=" + s.accessToken
	openpid = &PluginOpenPid{}
	body := make(bmap.BodyMap)
	body.Set("code", code)
	if err = s.doRequestPost(c, path, body, openpid); err != nil {
		return nil, err
	}
	if openpid.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", openpid.Errcode, openpid.Errmsg)
	}
	return openpid, nil
}

// GetPaidUnionid 用户支付完成后，获取该用户的 UnionId，无需用户授权
// 注意：errcode = 0 为成功
// openid：支付用户唯一标识
// transactionId：微信支付订单号
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/basic-info/getPaidUnionid.html
func (s *SDK) GetPaidUnionid(c context.Context, openid, transactionId string) (unionid *PaidUnionId, err error) {
	path := "/wxa/getpaidunionid?access_token=" + s.accessToken + "&openid=" + openid + "&transaction_id=" + transactionId
	unionid = &PaidUnionId{}
	if err = s.doRequestGet(c, path, unionid); err != nil {
		return nil, err
	}
	if unionid.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", unionid.Errcode, unionid.Errmsg)
	}
	return unionid, nil
}

// GetPaidUnionidByTradeNo 用户支付完成后，获取该用户的 UnionId，无需用户授权
// 注意：errcode = 0 为成功
// openid：支付用户唯一标识
// mchid：微信支付商户号
// tradeNo：微信支付商户订单号
// 文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (s *SDK) GetPaidUnionidByTradeNo(c context.Context, openid, mchid, tradeNo string) (unionid *PaidUnionId, err error) {
	path := "/wxa/getpaidunionid?access_token=" + s.accessToken + "&openid=" + openid + "&mch_id=" + mchid + "&out_trade_no=" + tradeNo
	unionid = &PaidUnionId{}
	if err = s.doRequestGet(c, path, unionid); err != nil {
		return nil, err
	}
	if unionid.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", unionid.Errcode, unionid.Errmsg)
	}
	return unionid, nil
}

// CheckEncryptedData 检查加密信息是否由微信生成
// 注意：errcode = 0 为成功
// encryptedData：加密数据，无需sha256操作
// 注意：（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
// 文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.checkEncryptedData.html
func (s *SDK) CheckEncryptedData(c context.Context, encryptedData string) (result *CheckEncryptedResult, err error) {
	path := "/wxa/business/checkencryptedmsg?access_token=" + s.accessToken
	h := sha256.New()
	h.Write([]byte(encryptedData))
	body := make(bmap.BodyMap)
	body.Set("encrypted_msg_hash", hex.EncodeToString(h.Sum(nil)))
	result = &CheckEncryptedResult{}
	if err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUserEncryptKey 获取用户encryptKey
// 注意：errcode = 0 为成功
// encryptedData：加密数据，无需sha256操作
// 注意：（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/internet/getUserEncryptKey.html
func (s *SDK) GetUserEncryptKey(c context.Context, openid, signature, sigMethod string) (uek *UserEncryptKey, err error) {
	path := "/wxa/business/getuserencryptkey?access_token=" + s.accessToken
	body := make(bmap.BodyMap)
	body.Set("openid", openid).
		Set("signature", signature).
		Set("sig_method", sigMethod)
	uek = &UserEncryptKey{}
	if err = s.doRequestPost(c, path, body, uek); err != nil {
		return nil, err
	}
	if uek.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", uek.Errcode, uek.Errmsg)
	}
	return uek, nil
}

// GetPhoneNumber 获取手机号
// 注意：errcode = 0 为成功
// code：手机号获取凭证
// 文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
func (s *SDK) GetPhoneNumber(c context.Context, code string) (pn *PhoneNumberRsp, err error) {
	path := "/wxa/business/getuserphonenumber?access_token=" + s.accessToken
	pn = &PhoneNumberRsp{}
	body := make(bmap.BodyMap)
	body.Set("code", code)
	if err = s.doRequestPost(c, path, body, pn); err != nil {
		return nil, err
	}
	if pn.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", pn.Errcode, pn.Errmsg)
	}
	return pn, nil
}
