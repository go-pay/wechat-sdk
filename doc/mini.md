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
* <font color='#07C160' size='4'>媒资上传</font>
	* 单个文件上传：`sdk.MediaAssetSingleFileUpload()`
	* 拉取上传：`sdk.MediaAssetPullUpload()`
	* 查询任务：`sdk.MediaAssetGetTask()`
	* 申请分片上传：`sdk.MediaAssetApplyUpload()`
	* 上传分片：`sdk.MediaAssetUploadPart()`
	* 确认上传：`sdk.MediaAssetCommitUpload()`
* <font color='#07C160' size='4'>媒资管理</font>
	* 获取媒资列表：`sdk.MediaAssetListMedia()`
	* 获取媒资详细信息：`sdk.MediaAssetGetMedia()`
	* 获取媒资播放链接：`sdk.MediaAssetGetMediaLink()`
	* 删除媒资：`sdk.MediaAssetDeleteMedia()`
* <font color='#07C160' size='4'>订阅消息</font>
	* 发送订阅消息：`sdk.SubscribeMessageSend()`
	* 获取模板列表：`sdk.GetTemplateList()`
	* 获取模板标题列表：`sdk.GetPubTemplateTitleList()`
	* 添加模板：`sdk.AddTemplate()`
	* 删除模板：`sdk.DeleteTemplate()`
	* 获取类目：`sdk.GetCategory()`
* <font color='#07C160' size='4'>小程序码</font>
	* 获取小程序码（数量较少）：`sdk.GetWxaCode()`
	* 获取小程序码（数量极多）：`sdk.GetWxaCodeUnlimit()`
	* 获取小程序二维码：`sdk.CreateWxaQRCode()`
* <font color='#07C160' size='4'>URL 跳转</font>
	* 生成 URL Scheme：`sdk.GenerateScheme()`
	* 查询 URL Scheme：`sdk.QueryScheme()`
	* 生成 URL Link：`sdk.GenerateUrlLink()`
	* 查询 URL Link：`sdk.QueryUrlLink()`
	* 生成 Short Link：`sdk.GenerateShortLink()`
* <font color='#07C160' size='4'>内容安全</font>
	* 文本内容安全检测：`sdk.MsgSecCheck()`
	* 图片内容安全检测：`sdk.ImgSecCheck()`
	* 音视频异步检测：`sdk.MediaCheckAsync()`
* <font color='#07C160' size='4'>数据分析</font>
	* 获取用户访问数据概况：`sdk.GetDailySummary()`
	* 获取访问数据日趋势：`sdk.GetDailyVisitTrend()`
	* 获取访问数据周趋势：`sdk.GetWeeklyVisitTrend()`
	* 获取访问数据月趋势：`sdk.GetMonthlyVisitTrend()`
	* 获取用户日留存：`sdk.GetDailyRetain()`
	* 获取用户周留存：`sdk.GetWeeklyRetain()`
	* 获取用户月留存：`sdk.GetMonthlyRetain()`
	* 获取访问页面数据：`sdk.GetVisitPage()`
	* 获取用户画像分布：`sdk.GetUserPortrait()`
	* 获取性能数据：`sdk.GetPerformanceData()`
	* 获取访问分布数据：`sdk.GetVisitDistribution()`
* <font color='#07C160' size='4'>图像处理</font>
	* 图片智能裁剪：`sdk.AiCrop()`
	* 条码/二维码识别：`sdk.ScanQRCode()`
	* 图片高清化：`sdk.SuperResolution()`
	* 身份证OCR识别：`sdk.OcrIdCard()`
	* 银行卡OCR识别：`sdk.OcrBankCard()`
	* 驾驶证OCR识别：`sdk.OcrDriving()`
	* 行驶证OCR识别：`sdk.OcrVehicleLicense()`
	* 营业执照OCR识别：`sdk.OcrBusinessLicense()`
	* 通用印刷体OCR识别：`sdk.OcrCommon()`
	* 车牌OCR识别：`sdk.OcrPlateNumber()`
* <font color='#07C160' size='4'>即时配送</font>
	* 拉取已绑定账号：`sdk.DeliveryGetAllAccount()`
	* 配送单预下单：`sdk.DeliveryPreAddOrder()`
	* 配送单下单：`sdk.DeliveryAddOrder()`
	* 配送单增加小费：`sdk.DeliveryAddTips()`
	* 配送单取消：`sdk.DeliveryCancelOrder()`
	* 配送单查询：`sdk.DeliveryGetOrder()`
	* 模拟配送公司更新配送单状态：`sdk.DeliveryMockUpdateOrder()`
	* 异常件退回商家商圈：`sdk.DeliveryAbnormalConfirm()`
* <font color='#07C160' size='4'>物流助手</font>
	* 生成运单：`sdk.ExpressAddOrder()`
	* 取消运单：`sdk.ExpressCancelOrder()`
	* 获取所有绑定的物流账号：`sdk.ExpressGetAllAccount()`
	* 获取电子面单余额：`sdk.ExpressGetQuota()`
	* 查询运单轨迹：`sdk.ExpressGetPath()`
	* 获取支持的快递公司列表：`sdk.ExpressGetAllDelivery()`
	* 获取打印员：`sdk.ExpressGetPrinter()`
	* 配置面单打印员：`sdk.ExpressUpdatePrinter()`
	* 获取面单联系人信息：`sdk.ExpressGetContact()`
* <font color='#07C160' size='4'>搜索</font>
	* 提交小程序页面：`sdk.SearchSubmitPages()`
	* 删除已提交的小程序页面：`sdk.SearchDeletePage()`
* <font color='#07C160' size='4'>运维中心</font>
	* 获取域名配置：`sdk.GetDomainInfo()`
	* 修改服务器域名：`sdk.ModifyDomain()`
	* 设置业务域名：`sdk.SetWebviewDomain()`
	* 获取小程序码扫码打开的页面：`sdk.GetQrcodeJumppublish()`
	* 设置小程序码扫码打开的页面：`sdk.SetQrcodeJump()`
* <font color='#07C160' size='4'>小程序管理</font>
	* 获取小程序基本信息：`sdk.GetAccountBasicInfo()`
	* 获取已上传的代码的页面列表：`sdk.GetPage()`
	* 获取授权小程序帐号的可选类目：`sdk.GetMiniCategory()`
	* 获取小程序的第三方提交代码的页面配置：`sdk.GetExtConfig()`
	* 设置小程序的第三方提交代码的页面配置：`sdk.SetExtConfig()`
* <font color='#07C160' size='4'>服务市场</font>
	* 调用服务平台提供的服务：`sdk.InvokeService()`

### 微信小程序 公共API

* `mini.GetAccessToken()` => 获取接口调用凭据
* `mini.GetStableAccessToken()` => 获取稳定版接口调用凭据