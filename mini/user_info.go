package mini

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	xaes "github.com/go-pay/wechat-sdk/pkg/aes"
	"github.com/go-pay/wechat-sdk/pkg/bm"
	"github.com/go-pay/wechat-sdk/pkg/util"
)

// GetPaidUnionid 用户支付完成后，获取该用户的 UnionId，无需用户授权
//	openid：支付用户唯一标识
//	transactionId：微信支付订单号
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (s *SDK) GetPaidUnionid(c context.Context, openid, transactionId string) (unionid *PaidUnionId, err error) {
	path := "/wxa/getpaidunionid?access_token=" + s.accessToken + "&openid=" + openid + "&transaction_id=" + transactionId
	unionid = new(PaidUnionId)
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
	path := "/wxa/getpaidunionid?access_token=" + s.accessToken + "&openid=" + openid + "&mch_id=" + mchid + "&out_trade_no=" + tradeNo
	unionid = new(PaidUnionId)
	if err = s.doRequestGet(c, path, unionid); err != nil {
		return nil, err
	}
	return
}

// CheckEncryptedData 检查加密信息是否由微信生成
//	encryptedData：加密数据，无需sha256操作
//	注意：（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
//	文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.checkEncryptedData.html
func (s *SDK) CheckEncryptedData(c context.Context, encryptedData string) (result *CheckEncryptedResult, err error) {
	path := "/wxa/business/checkencryptedmsg?access_token=" + s.accessToken
	h := sha256.New()
	h.Write([]byte(encryptedData))
	body := make(bm.BodyMap)
	body.Set("encrypted_msg_hash", hex.EncodeToString(h.Sum(nil)))
	result = new(CheckEncryptedResult)
	if err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	return
}

// DecryptOpenData 解密开放数据到结构体
//	encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//	iv：加密算法的初始向量，小程序获取
//	sessionKey：会话密钥，通过 sdk.Code2Session() 方法获取到
//	ptr：需要解析到的结构体指针，例：mini.UserPhone、mini.UserInfo
//	文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func (s *SDK) DecryptOpenData(encryptedData, iv, sessionKey string, ptr interface{}) (err error) {
	if encryptedData == util.NULL || iv == util.NULL || sessionKey == util.NULL {
		return errors.New("input params can not null")
	}
	var (
		cipherText, aesKey, ivKey, plainText []byte
		block                                cipher.Block
		blockMode                            cipher.BlockMode
	)
	beanValue := reflect.ValueOf(ptr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("ptr must be point type")
	}
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("ptr point must be struct type")
	}
	cipherText, _ = base64.StdEncoding.DecodeString(encryptedData)
	aesKey, _ = base64.StdEncoding.DecodeString(sessionKey)
	ivKey, _ = base64.StdEncoding.DecodeString(iv)
	if len(cipherText)%len(aesKey) != 0 {
		return errors.New("encryptedData error")
	}
	if block, err = aes.NewCipher(aesKey); err != nil {
		return fmt.Errorf("aes.NewCipher(),error(%w)", err)
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	plainText = make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	if len(plainText) > 0 {
		plainText = xaes.PKCS7UnPadding(plainText)
	}
	if err = json.Unmarshal(plainText, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v),error(%w)", string(plainText), ptr, err)
	}
	return
}
