## 微信公众号

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/public.md#%E9%99%84%E5%BD%95)

- 微信公众号：[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

---

### 具体使用请参考 `public/public_test.go`

- ### NewSDK

```go
import (
    "github.com/go-pay/wechat-sdk/public"
    "github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 初始化微信公众号 SDK
//	Appid：Appid
//	Secret：appSecret
//	autoManageToken：是否自动获取并自动维护刷新 AccessToken
publicSDK, err := public.New(Appid, Secret, true)
if err != nil {
    xlog.Error(err)
    return
}

// 打开Debug开关，输出日志
publicSDK.DebugSwitch = wechat.DebugOn
```

#### AccessToken

```go
import (
    "github.com/go-pay/wechat-sdk/public"
    "github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 若 autoManageToken 为 false，需要手动设置 Token
// publicSDK.SetPublicAccessToken("access_token")

// 获取AccessToken请通过此方法获取，之后请通过下面的回调方法获取
at := publicSDK.GetPublicAccessToken()

// 每次刷新 accessToken 后，此方法回调返回 accessToken 和 有效时间（秒）
publicSDK.SetPublicAccessTokenCallback(func(accessToken string, expireIn int, err error) {
	if err != nil {
		xlog.Errorf("refresh access token error(%+v)", err)
		return
	}
	xlog.Infof("accessToken: %s", accessToken)
	xlog.Infof("expireIn: %d", expireIn)
})
```


#### 创建二维码

```go
body := make(bm.BodyMap)
// 临时二维码
body.Set("expire_seconds", 604800).
	Set("action_name", "QR_SCENE").
	SetBodyMap("action_info", func(b bm.BodyMap) {
		b.SetBodyMap("scene", func(b bm.BodyMap) {
			b.Set("scene_id", 123)
		})
	})

rsp, err := publicSDK.QRCodeCreate(ctx, body)
if err != nil {
	xlog.Error(err)
	return
}
xlog.Infof("rsp:%+v", rsp)
```

## 附录：

### 微信公众号 服务端API

* <font color='#07C160' size='4'>用户管理</font>
	* 用户标签创建：`sdk.UserTagCreate()`
	* 获取已创建的用户标签列表：`sdk.UserTagList()`
	* 用户标签编辑更新：`sdk.UserTagUpdate()`
	* 用户标签删除：`sdk.UserTagDelete()`
	* 获取标签下粉丝列表：`sdk.UserTagFansList()`
	* 批量为用户打标签：`sdk.UserTagBatchTagging()`
	* 批量为用户取消标签：`sdk.UserTagBatchUnTagging()`
	* 获取用户身上的标签列表：`sdk.UserTagIdList()`
* <font color='#07C160' size='4'>账号管理</font>
	* 生成带参数的二维码：`sdk.QRCodeCreate()`
	* 生成短key托管：`sdk.ShortKeyGen()`
	* 获取托管的短key：`sdk.ShortKeyFetch()`
