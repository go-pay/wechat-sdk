## 微信公众号

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/open.md#%E9%99%84%E5%BD%95)

- 微信公众号：[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

---

## 附录：

### 微信公众号 服务端API

* <font color='#07C160' size='4'>用户管理</font>
	* 创建标签：`sdk.ss()`
	* 获取已创建的标签：`sdk.ss()`
	* 编辑标签：`sdk.ss()`
	* 删除标签：`sdk.ss()`
	* 获取标签下粉丝列表：`sdk.ss()`
	* 批量为用户打标签：`sdk.ss()`
	* 批量为用户取消标签：`sdk.ss()`
	* 获取用户身上的标签列表：`sdk.ss()`
* <font color='#07C160' size='4'>账号管理</font>
	* 生成带参数的二维码：`sdk.QRCodeCreate()`
	* 生成短key托管：`sdk.ShortKeyGen()`
	* 获取托管的短key：`sdk.ShortKeyFetch()`
