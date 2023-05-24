package public

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/go-pay/wechat-sdk/pkg/bmap"
)

// JsSDKUsePermissionSign 获取JS-SDK使用权限签名
// 文档介绍：https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html#62
func JsSDKUsePermissionSign(jsapiTicket, nonceStr, url string, timestamp int) (sign string) {
	bm := make(bmap.BodyMap)
	params := bm.Set("jsapi_ticket", jsapiTicket).
		Set("noncestr", nonceStr).
		Set("timestamp", timestamp).
		Set("url", url).EncodeSortedSignParams()
	hash := sha1.New()
	hash.Write([]byte(params))
	sign = hex.EncodeToString(hash.Sum(nil))
	return
}
