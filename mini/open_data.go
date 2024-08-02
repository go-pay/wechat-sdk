package mini

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-pay/crypto/aes"
	"github.com/go-pay/util"
	"github.com/go-pay/util/js"
)

// VerifyDecryptOpenData 数据签名校验
// rowData、signature：通过调用接口（如 wx.getUserInfo）获取数据时，接口会同时返回 rawData、signature
// sessionKey：会话密钥，通过 sdk.Code2Session() 方法获取到
// 文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func (s *SDK) VerifyDecryptOpenData(rowData, signature, sessionKey string) (ok bool) {
	signData := rowData + sessionKey
	hash := sha1.New()
	hash.Write([]byte(signData))
	sign := hex.EncodeToString(hash.Sum(nil))
	return sign == signature
}

// DecryptOpenData 解密开放数据到结构体
// encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
// iv：加密算法的初始向量，小程序获取
// sessionKey：会话密钥，通过 sdk.Code2Session() 方法获取到
// ptr：需要解析到的结构体指针，例：mini.UserPhone、mini.UserInfo
// 文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func (s *SDK) DecryptOpenData(encryptedData, iv, sessionKey string, ptr any) (err error) {
	if encryptedData == util.NULL || iv == util.NULL || sessionKey == util.NULL {
		return errors.New("input params can not null")
	}
	var (
		cipherText, aesKey, ivKey, plainText []byte
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
	plainText, err = aes.CBCDecrypt(cipherText, aesKey, ivKey)
	if err != nil {
		return fmt.Errorf("aes.CBCDecrypt(),err(%w)", err)
	}
	if err = js.UnmarshalBytes(plainText, ptr); err != nil {
		return fmt.Errorf("js.UnmarshalBytes(%s, %+v),error(%w)", string(plainText), ptr, err)
	}
	return
}
