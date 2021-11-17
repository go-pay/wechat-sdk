## 微信小程序

- 已实现API列表附录：[API 列表附录](https://github.com/go-pay/wechat-sdk/blob/main/doc/mini.md#%E9%99%84%E5%BD%95)

- 微信小程序服务端：[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)

- 开放数据校验与解密：[开放能力文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)

---

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
