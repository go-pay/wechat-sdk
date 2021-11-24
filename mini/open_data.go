package mini

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	xaes "github.com/go-pay/wechat-sdk/pkg/aes"
	"github.com/go-pay/wechat-sdk/pkg/util"
)

// VerifyDecryptOpenData 数据签名校验
//	rowData、signature：通过调用接口（如 wx.getUserInfo）获取数据时，接口会同时返回 rawData、signature
//	sessionKey：会话密钥，通过 sdk.Code2Session() 方法获取到
//	文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func (s *SDK) VerifyDecryptOpenData(rowData, signature, sessionKey string) (ok bool) {
	signData := rowData + sessionKey
	hash := sha1.New()
	hash.Write([]byte(signData))
	sign := hex.EncodeToString(hash.Sum(nil))
	return sign == signature
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
		return fmt.Errorf("json.Unmarshal(%m, %+v),error(%w)", string(plainText), ptr, err)
	}
	return
}
