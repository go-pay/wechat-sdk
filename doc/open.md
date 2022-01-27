## 微信公众号

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/open.md#%E9%99%84%E5%BD%95)

- 微信公众号：[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

---

### 具体使用请参考 `doc/open.md`

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

rsp, err := openSDK.QRCodeCreate(ctx, body)
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
