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

// New 初始化微信开放平台 SDK
// Appid：Appid
// Secret：appSecret
// autoManageToken：是否自动维护刷新 AccessToken（用户量较少时推荐使用，默认10分钟轮询检测一次，发现有效期小于1.5倍轮询时间时，自动刷新）
openSDK, err := open.New(Appid, Secret, true)
if err != nil {
    xlog.Error(err)
    return
}

// 打开Debug开关，输出日志
openSDK.DebugSwitch = wechat.DebugOn

// 可自行设置 AccessToken 刷新间隔
//openSDK.SetAccessTokenRefreshInternal(5 * time.Minute)

// 此方法回调返回 AccessToken
openSDK.SetAccessTokenCallback(func(at *AT, err error) {
    if err != nil {
        xlog.Errorf("call back access token err:%+v", err)
        return
    }
    xlog.Infof("call back access token: %v", at)
})
```

#### 通过 code 获取用户 access_token

```go
at, err := openSDK.Code2AccessToken(ctx, "code")
if err != nil {
    xlog.Error(err)
    return
}
xlog.Infof("at: %s", at)
```

#### 刷新或续期 access_token 使用

```go
at, err := openSDK.RefreshAccessToken(ctx, "refreshToken")
if err != nil {
    xlog.Error(err)
    return
}
xlog.Infof("at: %s", at)
```

#### 检验授权凭证（access_token）是否有效

```go
err := openSDK.CheckAccessToken(ctx, "accessToken", "openid")
if err != nil {
    xlog.Errorf("CheckAccessToken,err:%v", err)
    return
}
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

### 微信开放平台 服务端API

* <font color='#07C160' size='4'>微信登录功能</font>
	* 通过 code 获取 access_token：`sdk.Code2AccessToken()`
	* 刷新或续期 access_token：`sdk.RefreshAccessToken()`
	* 检验授权凭证 access_token 是否有效：`sdk.CheckAccessToken()`
	* 获取用户个人信息（UnionID 机制）：`sdk.UserInfo()`