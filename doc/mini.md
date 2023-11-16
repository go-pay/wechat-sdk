## 微信小程序

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/mini.md#%E9%99%84%E5%BD%95)

- 微信小程序服务端：[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)

- 开放数据校验与解密：[开放能力文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)

---

### 具体使用请参考 `mini/mini_test.go`

- ### NewSDK

```go
import (
    "github.com/go-pay/wechat-sdk/mini"
    "github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 初始化微信小程序 SDK
// Appid：Appid
// Secret：appSecret
// autoManageToken：是否自动获取并自动维护刷新 AccessToken，默认使用稳定版接口且force_refresh=false
miniSDK, err := mini.New(Appid, Secret, true)
if err != nil {
    xlog.Error(err)
    return
}

// 打开Debug开关，输出日志
miniSDK.DebugSwitch = wechat.DebugOn
```

#### AccessToken

```go
import (
    "github.com/go-pay/wechat-sdk/mini"
    "github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 若 autoManageToken 为 false，需要手动设置 Token
// miniSDK.SetMiniAccessToken("access_token")

// 获取AccessToken请通过此方法获取，之后请通过下面的回调方法获取
at := miniSDK.GetMiniAccessToken()

// 每次刷新 accessToken 后，此方法回调返回 accessToken 和 有效时间（秒）
miniSDK.SetMiniAccessTokenCallback(func(appid, accessToken string, expireIn int, err error) {
    if err != nil {
        xlog.Errorf("refresh access token error(%+v)", err)
        return
    }
    xlog.Infof("appid:%s, accessToken: %s",appid, accessToken)
    xlog.Infof("expireIn: %d", expireIn)
})
```

#### Code2Session

```go
session, err := miniSDK.Code2Session(ctx, "wxCode")
if err != nil {
    xlog.Error(err)
    return
}
xlog.Debugf("at:%+v", session)
```

#### 开放数据

- 开放数据校验

```go
rwData := `{"nickName":"Band","gender":1,"language":"zh_CN","city":"Guangzhou","province":"Guangdong","country":"CN","avatarUrl":"http://wx.qlogo.cn/mmopen/vi_32/1vZvI39NWFQ9XM4LtQpFrQJ1xlgZxx3w7bQxKARol6503Iuswjjn6nIGBiaycAjAtpujxyzYsrztuuICqIM5ibXQ/0"}`
sign := "75e81ceda165f4ffa64f4068af58c64b8f54b88c"
sessionKey := "HyVFkGl5F5OQWJZZaNzBBg=="
ok := miniSDK.VerifyDecryptOpenData(rwData, sign, sessionKey)
xlog.Debugf("verify result: %t", ok)
```

- 开放数据解密

```go
data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
iv := "Cds8j3VYoGvnTp1BrjXdJg=="
session := "lyY4HPQbaOYzZdG+JcYK9w=="

// 微信小程序 手机号
phone := new(model.UserPhone)
err := miniSDK.DecryptOpenData(data, iv, session, phone)
if err != nil {
    xlog.Error(err)
    return
}

sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
iv2 := "r7BXXKkLb8qrSNn05n0qiA=="

// 微信小程序 用户信息
userInfo := new(model.UserInfo)
err := miniSDK.DecryptOpenData(encryptedData, iv2, sessionKey, userInfo)
if err != nil {
    xlog.Error(err)
    return
}
```

## 附录：

### 微信小程序 服务端API

* <font color='#07C160' size='4'>登录</font>
	* Code2Session：`sdk.Code2Session()`
* <font color='#07C160' size='4'>用户信息</font>
	* 检查加密信息：`sdk.CheckEncryptedData()`
	* 获取支付后的UnionId：`sdk.GetPaidUnionid()`
	* 获取支付后的UnionId：`sdk.GetPaidUnionidByTradeNo()`
	* 获取插件用户OpenPid：`sdk.GetPluginOpenPid()`
	* 获取用户EncryptKey：`sdk.GetUserEncryptKey()`
	* 获取手机号：`sdk.GetPhoneNumber()`
* <font color='#07C160' size='4'>开放数据</font>
	* 开放数据校验：`sdk.VerifyDecryptOpenData()`
	* 开放数据解密：`sdk.DecryptOpenData()`
* <font color='#07C160' size='4'>客服消息</font>
	* 获取客服消息内的临时素材：`sdk.CSMessageGetTempMedia()`
	* 发送客服消息：`sdk.CSMessageSend()`
	* 下发客服当前输入状态给用户：`sdk.CSMessageSetTyping()`
	* 新增图片素材：`sdk.CSMessageUploadTempMedia()`
* <font color='#07C160' size='4'>统一服务消息</font>
	* 发送统一服务消息：`sdk.UniformMessageSend()` 未完成

### 微信小程序 公共API

* `mini.GetAccessToken()` => 获取接口调用凭据
* `mini.GetStableAccessToken()` => 获取稳定版接口调用凭据