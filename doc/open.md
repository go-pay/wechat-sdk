## 微信开放平台

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/public.md#%E9%99%84%E5%BD%95)

- 微信开放平台：[官方文档](https://developers.weixin.qq.com/doc/oplatform/Mobile_App/Resource_Center_Homepage.html)

---

### 具体使用请参考 `doc/open.md`

#### 通过 code 获取 access_token

```go

```

## 附录：

### 微信开发平台 服务端API

* <font color='#07C160' size='4'>微信登录功能</font>
	* 通过 code 获取 access_token：``sdk.()`
	* 刷新或续期 access_token 使用：`sdk.()`
	* 检验授权凭证 access_token 是否有效：`sdk.()`
	* 获取用户个人信息（UnionID 机制）：`sdk.UserTagDelete()`