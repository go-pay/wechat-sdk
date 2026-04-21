# 微信小程序服务端 API 实现情况

## ✅ 已实现模块

### 1. 登录
- ✅ Code2Session - 登录凭证校验

### 2. 用户信息
- ✅ GetPluginOpenPid - 获取插件用户OpenPid
- ✅ GetPaidUnionid - 获取支付后的UnionId（通过微信支付订单号）
- ✅ GetPaidUnionidByTradeNo - 获取支付后的UnionId（通过商户订单号）
- ✅ CheckEncryptedData - 检查加密信息
- ✅ GetUserEncryptKey - 获取用户EncryptKey
- ✅ GetPhoneNumber - 获取手机号

### 3. 开放数据
- ✅ VerifyDecryptOpenData - 开放数据校验
- ✅ DecryptOpenData - 开放数据解密

### 4. 客服消息
- ✅ CSMessageGetTempMedia - 获取客服消息内的临时素材
- ✅ CSMessageSend - 发送客服消息
- ✅ CSMessageSetTyping - 下发客服当前输入状态
- ✅ CSMessageUploadTempMedia - 新增图片素材

### 5. 统一服务消息
- ✅ UniformMessageSend - 发送统一服务消息

### 6. 订阅消息 ⭐ 新增
- ✅ SubscribeMessageSend - 发送订阅消息
- ✅ GetTemplateList - 获取模板列表
- ✅ GetPubTemplateTitleList - 获取模板标题列表
- ✅ AddTemplate - 添加模板
- ✅ DeleteTemplate - 删除模板
- ✅ GetCategory - 获取类目

### 7. 小程序码 ⭐ 新增
- ✅ GetWxaCode - 获取小程序码（数量较少）
- ✅ GetWxaCodeUnlimit - 获取小程序码（数量极多）
- ✅ CreateWxaQRCode - 获取小程序二维码

### 8. URL 跳转 ⭐ 新增
- ✅ GenerateScheme - 生成 URL Scheme
- ✅ QueryScheme - 查询 URL Scheme
- ✅ GenerateUrlLink - 生成 URL Link
- ✅ QueryUrlLink - 查询 URL Link
- ✅ GenerateShortLink - 生成 Short Link

### 9. 内容安全 ⭐ 新增
- ✅ MsgSecCheck - 文本内容安全检测
- ✅ ImgSecCheck - 图片内容安全检测
- ✅ MediaCheckAsync - 音视频异步检测

### 10. 数据分析 ⭐ 新增
- ✅ GetDailySummary - 获取用户访问数据概况
- ✅ GetDailyVisitTrend - 获取访问数据日趋势
- ✅ GetWeeklyVisitTrend - 获取访问数据周趋势
- ✅ GetMonthlyVisitTrend - 获取访问数据月趋势
- ✅ GetDailyRetain - 获取用户日留存
- ✅ GetWeeklyRetain - 获取用户周留存
- ✅ GetMonthlyRetain - 获取用户月留存
- ✅ GetVisitPage - 获取访问页面数据
- ✅ GetUserPortrait - 获取用户画像分布
- ✅ GetPerformanceData - 获取性能数据
- ✅ GetVisitDistribution - 获取访问分布数据

### 11. 媒资上传（短剧）
- ✅ MediaAssetSingleFileUpload - 单个文件上传
- ✅ MediaAssetPullUpload - 拉取上传
- ✅ MediaAssetGetTask - 查询任务
- ✅ MediaAssetApplyUpload - 申请分片上传
- ✅ MediaAssetUploadPart - 上传分片
- ✅ MediaAssetCommitUpload - 确认上传

### 12. 媒资管理（短剧）
- ✅ MediaAssetListMedia - 获取媒资列表
- ✅ MediaAssetGetMedia - 获取媒资详细信息
- ✅ MediaAssetGetMediaLink - 获取媒资播放链接
- ✅ MediaAssetDeleteMedia - 删除媒资

### 13. 接口调用凭证
- ✅ GetAccessToken - 获取接口调用凭据
- ✅ GetStableAccessToken - 获取稳定版接口调用凭据

### 14. 图像处理 ⭐ 新增
- ✅ AiCrop - 图片智能裁剪
- ✅ ScanQRCode - 条码/二维码识别
- ✅ SuperResolution - 图片高清化
- ✅ OcrIdCard - 身份证OCR识别
- ✅ OcrBankCard - 银行卡OCR识别
- ✅ OcrDriving - 驾驶证OCR识别
- ✅ OcrVehicleLicense - 行驶证OCR识别
- ✅ OcrBusinessLicense - 营业执照OCR识别
- ✅ OcrCommon - 通用印刷体OCR识别
- ✅ OcrPlateNumber - 车牌OCR识别

### 15. 即时配送 ⭐ 新增
- ✅ DeliveryGetAllAccount - 拉取已绑定账号
- ✅ DeliveryPreAddOrder - 配送单预下单
- ✅ DeliveryAddOrder - 配送单下单
- ✅ DeliveryAddTips - 配送单增加小费
- ✅ DeliveryCancelOrder - 配送单取消
- ✅ DeliveryGetOrder - 配送单查询
- ✅ DeliveryMockUpdateOrder - 模拟配送公司更新配送单状态
- ✅ DeliveryAbnormalConfirm - 异常件退回商家商圈

### 16. 物流助手 ⭐ 新增
- ✅ ExpressAddOrder - 生成运单
- ✅ ExpressCancelOrder - 取消运单
- ✅ ExpressGetAllAccount - 获取所有绑定的物流账号
- ✅ ExpressGetQuota - 获取电子面单余额
- ✅ ExpressGetPath - 查询运单轨迹
- ✅ ExpressGetAllDelivery - 获取支持的快递公司列表
- ✅ ExpressGetPrinter - 获取打印员
- ✅ ExpressUpdatePrinter - 配置面单打印员
- ✅ ExpressGetContact - 获取面单联系人信息

### 17. 搜索 ⭐ 新增
- ✅ SearchSubmitPages - 提交小程序页面url及参数信息
- ✅ SearchDeletePage - 删除已提交的小程序页面

### 18. 运维中心 ⭐ 新增
- ✅ GetDomainInfo - 获取域名配置
- ✅ ModifyDomain - 修改服务器域名
- ✅ SetWebviewDomain - 设置业务域名
- ✅ GetQrcodeJumppublish - 获取小程序码扫码打开的页面
- ✅ SetQrcodeJump - 设置小程序码扫码打开的页面

### 19. 小程序管理 ⭐ 新增
- ✅ GetAccountBasicInfo - 获取小程序基本信息
- ✅ GetPage - 获取已上传的代码的页面列表
- ✅ GetMiniCategory - 获取授权小程序帐号的可选类目
- ✅ GetExtConfig - 获取小程序的第三方提交代码的页面配置
- ✅ SetExtConfig - 设置小程序的第三方提交代码的页面配置

### 20. 服务市场 ⭐ 新增
- ✅ InvokeService - 调用服务平台提供的服务

### 21. 虚拟支付 ⭐ 新增
- ✅ XpayQueryUserBalance - 查询代币余额
- ✅ XpayCurrencyPay - 扣减代币
- ✅ XpayQueryOrder - 查询创建的订单
- ✅ XpayCancelCurrencyPay - 代币支付退款
- ✅ XpayNotifyProvideGoods - 通知已发货完成
- ✅ XpayPresentCurrency - 代币赠送
- ✅ XpayDownloadBill - 下载小程序账单
- ✅ XpayRefundOrder - 启动订单退款任务
- ✅ XpayCreateWithdrawOrder - 创建提现单
- ✅ XpayQueryWithdrawOrder - 查询提现单
- ✅ XpayStartUploadGoods - 批量上传道具
- ✅ XpayQueryUploadGoods - 查询批量上传道具任务
- ✅ XpayStartPublishGoods - 启动批量发布道具任务
- ✅ XpayQueryPublishGoods - 查询批量发布道具任务
- ✅ XpayQueryBizBalance - 查询商家账户可提现余额
- ✅ XpayQueryTransferAccount - 查询广告金充值账户
- ✅ XpayQueryAdverFunds - 查询广告金发放记录
- ✅ XpayCreateFundsBill - 充值广告金
- ✅ XpayBindTransferAccount - 绑定广告金充值账户
- ✅ XpayQueryFundsBill - 查询广告金充值记录
- ✅ XpayQueryRecoverBill - 查询广告金回收记录
- ✅ XpayQuerySubscribeContract - 查询签约关系
- ✅ XpaySendSubscribePrePayment - 预通知扣款
- ✅ XpaySubmitSubscribePayOrder - 发起订阅扣款
- ✅ XpayCancelSubscribeContract - 商家解约
- ✅ XpayGetComplaintList - 获取投诉列表
- ✅ XpayGetComplaintDetail - 获取投诉详情
- ✅ XpayGetNegotiationHistory - 获取协商历史
- ✅ XpayResponseComplaint - 回复用户投诉
- ✅ XpayCompleteComplaint - 完成投诉处理
- ✅ XpayUploadVpFile - 上传媒体文件
- ✅ XpayGetUploadFileSign - 获取上传文件签名

---

## ❌ 未实现且较少使用的模块

### 1. 云开发 (cloudbase)
- 数据库操作（增删改查、聚合、导入导出等）
- 云函数调用
- 云存储（上传、下载、删除文件）
- 短信发送
- 其他云开发能力

### 2. 城市服务 (cityservice)
- 基础能力（获取城市服务限定页面链接、消息通路等）
- 微信长辈就医
- 微信就医助手

### 3. B2B 支付
- 商户号进件
- 支付、退款、查询订单
- 分账相关
- 提现相关

### 4. 付费能力 (charge)
- 获取付费能力用量数据
- 查询购买资源包用量

### 6. 直播
- 创建直播间
- 获取直播间列表
- 获取直播间回放
- 添加管理员
- 删除管理员
- 获取管理员列表
- 添加主播
- 删除主播
- 获取主播列表
- 添加商品
- 删除商品
- 获取商品列表
- 推送商品
- 下架商品
- 获取商品状态
- 获取成员列表
- 添加成员
- 删除成员

### 7. 小程序直播
- 创建直播间
- 获取直播间列表和回放
- 导入商品
- 商品管理
- 成员管理
- 订阅管理

### 8. 广告
- 获取广告金数据
- 获取广告收入数据
