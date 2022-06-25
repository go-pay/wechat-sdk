## 微信开放平台

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/public.md#%E9%99%84%E5%BD%95)

- 微信开放平台：[官方文档](https://developers.weixin.qq.com/doc/oplatform/Mobile_App/Resource_Center_Homepage.html)

---

### 具体使用请参考 `open/open_test.go`

- ### NewSDK

```go
import (
    "github.com/go-pay/wechat-sdk/open"
    "github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 初始化微信开放平台 SDK
//	Appid：Appid
//	Secret：appSecret
//	autoManageToken：是否自动获取并自动维护刷新 AccessToken
openSDK, err := open.New(Appid, Secret, true)
if err != nil {
    xlog.Error(err)
    return
}

// 打开Debug开关，输出日志
openSDK.DebugSwitch = wechat.DebugOn
```


#### 通过 code 获取 access_token

```go
// 如果 自行维护 AccessToken，请需要手动设置 Token
// openSDK.SetOpenAccessToken("access_token")

// 注意：必须优先换取 开放平台 AccessToken，否则会导致部分接口调用失败
at, err := openSDK.Code2AccessToken(ctx, "code")
if err != nil {
	xlog.Error(err)
	return
}
xlog.Infof("at: %s", at)

// 每次刷新 accessToken 后，此方法回调返回 accessToken 和 有效时间（秒）
openSDK.SetOpenAccessTokenCallback(func(at *AccessToken, err error) {
	if err != nil {
		xlog.Errorf("refresh access token error(%+v)", err)
		return
	}
	xlog.Infof("AccessToken: %+v", at)
})
```

#### 获取用户信息

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