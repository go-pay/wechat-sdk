## 微信开放平台

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/public.md#%E9%99%84%E5%BD%95)

- 微信开放平台：[官方文档](https://developers.weixin.qq.com/doc/oplatform/Mobile_App/Resource_Center_Homepage.html)

---

### 具体使用请参考 `sdk_open_test.go`

#### 通过 code 获取 access_token

```go
// 注意：必须换取 开放平台 自己的AccessToken，与小程序和公众号不通用
openAT, err := openSDK.Code2AccessToken(ctx, "xxx")
if err != nil {
    xlog.Error(err)
    return
}
xlog.Infof("open at: %+v", openAT)
// 每次刷新 accessToken 后，此方法回调返回 accessToken 和 有效时间（秒）
openSDK.SetOpenATCallback(func(at *open.AccessToken, err error) {
    if err != nil {
        xlog.Errorf("refresh access token error(%+v)", err)
        return
    }
    xlog.Infof("AccessToken: %+v", at)
})
```

#### 通过 code 获取 access_token

```go
rsp, err := openSDK.UserInfo(ctx, "openid", "zh_CN")
if err != nil {
    xlog.Error(err)
    return
}
xlog.Infof("rsp:%+v", rsp)
```

## 附录：

### 微信开发平台 服务端API

* <font color='#07C160' size='4'>微信登录功能</font>
	* 通过 code 获取 access_token：``sdk.Code2AccessToken()`
	* 检验授权凭证 access_token 是否有效：`sdk.CheckAccessToken()`
	* 获取用户个人信息（UnionID 机制）：`sdk.UserInfo()`