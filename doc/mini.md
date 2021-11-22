## 微信小程序

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/mini.md#%E9%99%84%E5%BD%95)

- 微信小程序服务端：[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)

- 开放数据校验与解密：[开放能力文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)

---

### 初始化 SDK

```go
// NewSDK 初始化微信小程序 SDK
//	appid：小程序 appid
//	secret：小程序 appSecret
//	accessToken：微信小程序AccessToken，若此参数为空，则自动获取并自动维护刷新
wxsdk, err = NewSDK(Appid, Secret)
if err != nil {
    xlog.Error(err)
    return
}

// 如需替换host节点，通过如下方法
// wxsdk.SetHost(wechat.HostSH)
```

### AccessToken 问题处理

> wechat-sdk 已支持获取并自动刷新 AccessToken 的维护，开发者可通过 `sdk.SetAccessTokenCallback()` 方法，间接拿到 AccessToken 等信息。

- 刚初始化完 SDK，AccessToken 通过下面方法获取
```go
// New完SDK，首次获取AccessToken请通过此方法获取，之后请通过下面的回调方法获取
at := wxsdk.GetAccessToken()
xlog.Infof("at: %s", at)
```
- 后续通过注册AT回调方法，获取每次刷新后的 AccessToken
```go
// 每次刷新 AccessToken 后，此方法回调返回 AccessToken 和 有效时间（秒）
wxsdk.SetAccessTokenCallback(func(accessToken string, expireIn int, err error) {
    if err != nil {
        xlog.Errorf("refresh access token error(%+v)", err)
    }
    xlog.Infof("accessToken: %s", accessToken)
    xlog.Infof("expireIn: %d", expireIn)
})
```

## 附录：

### 微信小程序 服务端API

* <font color='#07C160' size='4'>登录</font>
    * Code2Session：`sdk.Code2Session()`
* <font color='#07C160' size='4'>用户信息</font>
    * 检查加密信息：`sdk.CheckEncryptedData()`
    * 获取支付后的UnionId：`sdk.GetPaidUnionid()`
    * 获取支付后的UnionId：`sdk.GetPaidUnionidByTradeNo()`
* <font color='#07C160' size='4'>开放数据</font>
    * 开放数据校验：`sdk.VerifyDecryptOpenData()`
    * 开放数据解密：`sdk.DecryptOpenData()`
* <font color='#07C160' size='4'>客服消息</font>
    * 获取客服消息内的临时素材：`sdk.CSMessageGetTempMedia()`
    * 发送客服消息给用户：`sdk.CSMessageSend()`
    * 下发客服当前输入状态给用户：`sdk.CSMessageSetTyping()`
    * 把媒体文件上传到微信服务器（目前仅支持图片）：`sdk.CSMessageUploadTempMedia()`
* <font color='#07C160' size='4'>统一服务消息</font>
    * 发送统一服务消息：`sdk.UniformMessageSend()` 未完成
