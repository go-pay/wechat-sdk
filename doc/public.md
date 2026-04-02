## 微信公众号

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/public.md#%E9%99%84%E5%BD%95)

- 微信公众号：[官方文档](https://developers.weixin.qq.com/doc/subscription/guide/)

---

### 具体使用请参考 `public/public_test.go`

- ### NewSDK

```go
import (
    "github.com/go-pay/wechat-sdk/public"
    "github.com/go-pay/wechat-sdk/pkg/xlog"
)

// 初始化微信公众号 SDK
// Appid：Appid
// Secret：appSecret
// autoManageToken：是否自动获取并自动维护刷新 AccessToken
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
publicSDK.SetPublicAccessTokenCallback(func(appid, accessToken string, expireIn int, err error) {
    if err != nil {
        xlog.Errorf("refresh access token error(%+v)", err)
        return
    }
    xlog.Infof("appid:%s, accessToken: %s",appid, accessToken)
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
	* 设置用户备注名：`sdk.UserInfoUpdateRemark()`
	* 获取用户基本信息：`sdk.UserInfoGet()`
	* 批量获取用户基本信息：`sdk.UserInfoBatchGet()`
	* 获取用户列表：`sdk.UserGet()`
	* 获取黑名单列表：`sdk.BlackListGetList()`
	* 拉黑用户：`sdk.BlackListBatchBlackList()`
	* 取消拉黑用户：`sdk.BlackListBatchUnBlackList()`
* <font color='#07C160' size='4'>账号管理</font>
	* 生成带参数的二维码：`sdk.QRCodeCreate()`
	* 生成短key托管：`sdk.ShortKeyGen()`
	* 获取托管的短key：`sdk.ShortKeyFetch()`
* <font color='#07C160' size='4'>自定义菜单</font>
	* 创建自定义菜单：`sdk.MenuCreate()`
	* 查询自定义菜单：`sdk.MenuGet()`
	* 删除自定义菜单：`sdk.MenuDelete()`
	* 创建个性化菜单：`sdk.MenuAddConditional()`
	* 删除个性化菜单：`sdk.MenuDelConditional()`
	* 测试个性化菜单匹配结果：`sdk.MenuTryMatch()`
	* 获取自定义菜单配置：`sdk.GetCurrentSelfMenuInfo()`
* <font color='#07C160' size='4'>素材管理</font>
	* 新增临时素材：`sdk.MediaUpload()`
	* 获取临时素材：`sdk.MediaGet()`
	* 新增永久图文素材：`sdk.MaterialAddNews()`
	* 上传图文消息内的图片：`sdk.MaterialUploadImg()`
	* 新增其他类型永久素材：`sdk.MaterialAddMaterial()`
	* 获取永久素材：`sdk.MaterialGetMaterial()`
	* 删除永久素材：`sdk.MaterialDelMaterial()`
	* 修改永久图文素材：`sdk.MaterialUpdateNews()`
	* 获取素材总数：`sdk.MaterialGetMaterialCount()`
	* 获取素材列表：`sdk.MaterialBatchGetMaterial()`
* <font color='#07C160' size='4'>客服消息</font>
	* 添加客服账号：`sdk.KfAccountAdd()`
	* 修改客服账号：`sdk.KfAccountUpdate()`
	* 删除客服账号：`sdk.KfAccountDel()`
	* 设置客服账号的头像：`sdk.KfAccountUploadHeadImg()`
	* 获取所有客服账号：`sdk.KfAccountGetList()`
	* 邀请绑定客服账号：`sdk.KfAccountInviteWorker()`
	* 创建会话：`sdk.KfSessionCreate()`
	* 关闭会话：`sdk.KfSessionClose()`
	* 获取客户会话状态：`sdk.KfSessionGetSession()`
	* 获取客服会话列表：`sdk.KfSessionGetSessionList()`
	* 获取未接入会话列表：`sdk.KfSessionGetWaitCase()`
	* 发送客服消息：`sdk.CustomSend()`
	* 客服输入状态：`sdk.CustomTyping()`
	* 获取聊天记录：`sdk.MsgRecordList()`
* <font color='#07C160' size='4'>微信网页开发</font>
	* 获取 jsapi_ticket：`sdk.GetJsApiTicket()`
	* 获取卡券 api_ticket：`sdk.GetApiTicket()`
* <font color='#07C160' size='4'>基础接口</font>
	* 获取微信服务器IP地址：`sdk.GetApiDomainIp()`
	* 获取微信callback IP地址：`sdk.GetCallbackIp()`
	* 清空API调用quota：`sdk.ClearQuota()`
	* 查询API调用额度：`sdk.GetApiQuota()`
	* 查询rid信息：`sdk.GetRid()`
* <font color='#07C160' size='4'>数据统计</font>
	* 获取用户增减数据：`sdk.GetUserSummary()`
	* 获取累计用户数据：`sdk.GetUserCumulate()`
	* 获取图文群发每日数据：`sdk.GetArticleSummary()`
	* 获取图文群发总数据：`sdk.GetArticleTotal()`
	* 获取图文统计数据：`sdk.GetUserRead()`
	* 获取图文统计分时数据：`sdk.GetUserReadHour()`
	* 获取图文分享转发数据：`sdk.GetUserShare()`
	* 获取图文分享转发分时数据：`sdk.GetUserShareHour()`
	* 获取消息发送概况数据：`sdk.GetUpstreamMsg()`
	* 获取消息发送分时数据：`sdk.GetUpstreamMsgHour()`
	* 获取消息发送周数据：`sdk.GetUpstreamMsgWeek()`
	* 获取消息发送月数据：`sdk.GetUpstreamMsgMonth()`
	* 获取消息发送分布数据：`sdk.GetUpstreamMsgDist()`
	* 获取消息发送分布周数据：`sdk.GetUpstreamMsgDistWeek()`
	* 获取消息发送分布月数据：`sdk.GetUpstreamMsgDistMonth()`
	* 获取接口分析数据：`sdk.GetInterfaceSummary()`
	* 获取接口分析分时数据：`sdk.GetInterfaceSummaryHour()`
* <font color='#07C160' size='4'>草稿箱</font>
	* 新建草稿：`sdk.DraftAddDraft()`
	* 获取草稿：`sdk.DraftGetDraft()`
	* 删除草稿：`sdk.DraftDelDraft()`
	* 修改草稿：`sdk.DraftUpdateDraft()`
* <font color='#07C160' size='4'>发布能力</font>
	* 发布接口：`sdk.FreepublishSubmit()`
	* 获取发布详情：`sdk.FreepublishGet()`
	* 删除发布：`sdk.FreepublishDelete()`
	* 通过article_id获取已发布文章：`sdk.FreepublishGetArticle()`
	* 获取成功发布列表：`sdk.FreepublishBatchGet()`
* <font color='#07C160' size='4'>评论管理</font>
	* 打开已群发文章评论：`sdk.CommentOpen()`
	* 关闭已群发文章评论：`sdk.CommentClose()`
	* 查看指定文章的评论数据：`sdk.CommentList()`
	* 将评论标记精选：`sdk.CommentMarkElect()`
	* 将评论取消精选：`sdk.CommentUnmarkElect()`
	* 删除评论：`sdk.CommentDelete()`
	* 回复评论：`sdk.CommentReplyAdd()`
* <font color='#07C160' size='4'>AI开放能力</font>
	* 语音识别：`sdk.VoiceTranslate()`
	* 二维码/条码识别：`sdk.QrcodeImgScan()`
	* 图片智能裁剪：`sdk.AiCropImg()`
	* 图片高清化：`sdk.SuperResolutionImg()`
	* 身份证OCR识别：`sdk.OcrIdCardImg()`
	* 银行卡OCR识别：`sdk.OcrBankCardImg()`
	* 行驶证OCR识别：`sdk.OcrDrivingImg()`
	* 驾驶证OCR识别：`sdk.OcrDrivingLicenseImg()`
	* 营业执照OCR识别：`sdk.OcrBizLicenseImg()`
	* 通用印刷体OCR识别：`sdk.OcrCommonImg()`

### 微信公众号 公共API

* `public.JsSDKUsePermissionSign()` => 获取JS-SDK使用权限签名
