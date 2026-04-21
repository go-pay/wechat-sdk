# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

这是一个微信 Golang SDK（当前版本 v1.1.10），提供微信小程序、公众号、开放平台三个模块的服务端 API 封装。
Go module: `github.com/go-pay/wechat-sdk`，Go 版本要求 1.24.0。

## 常用命令

### 测试
```bash
# 运行所有测试
go test -v ./...

# 测试特定模块
go test -v ./mini
go test -v ./public
go test -v ./open

# 运行单个测试文件
go test -v ./mini/mini_test.go
```

### 构建
```bash
# 安装依赖
go mod tidy

# 验证模块
go mod verify
```

## 代码架构

### 项目结构总览

```
wechat-sdk/
├── common.go          # 全局常量（Version、DebugSwitch、Host 映射、HeaderRequestID）
├── go.mod / go.sum    # 依赖管理
├── mini/              # 微信小程序 SDK（26 个文件）
├── public/            # 微信公众号 SDK（17 个文件）
├── open/              # 微信开放平台 SDK（5 个文件）
├── doc/               # API 文档（mini.md, public.md, open.md, *_api_coverage.md）
├── release_note.md    # 版本记录
└── README.md
```

### 三大模块结构

项目按微信平台类型划分为三个独立模块，每个模块都遵循相同的设计模式：

- **mini/** - 微信小程序 SDK
- **public/** - 微信公众号 SDK
- **open/** - 微信开放平台 SDK

每个模块的核心结构：
- `SDK` 结构体：包含 Appid、Secret、accessToken、自动刷新机制
- 自动管理 AccessToken：通过 `autoManageToken` 参数控制，使用稳定版接口
- Debug 开关：通过 `DebugSwitch` 控制日志输出
- HTTP 客户端：基于 `xhttp.Client`，所有请求带 `Request-ID` header

### mini/ 模块文件说明

| 文件 | 用途 |
|------|------|
| `mini.go` | SDK 结构体定义、初始化、包级别请求函数（doRequestGet/doRequestPost） |
| `request.go` | SDK 实例级请求方法（DoRequestGet/DoRequestPost/DoRequestPostFile/DoRequestGetByte） |
| `model.go` | 所有请求/响应结构体、常量定义（Success=0, HostDefault 等） |
| `access_token.go` | AccessToken 管理（稳定版接口、自动刷新） |
| `login.go` | 用户登录/session 相关 |
| `user_info.go` | 用户信息获取（手机号等） |
| `qrcode.go` | 小程序码/二维码生成（含 doRequestPostForImage 图片响应处理） |
| `url_link.go` | URL Scheme/Link 生成 |
| `subscribe_message.go` | 订阅消息模板管理 |
| `uniform_message.go` | 统一服务消息 |
| `delivery.go` | 即时配送 |
| `express.go` | 物流管理 |
| `search.go` | 搜索页面提交 |
| `domain.go` | 域名管理 |
| `management.go` | 小程序账户管理 |
| `customer_service.go` | 客服消息 |
| `data_analysis.go` | 数据分析 |
| `image_process.go` | 图像处理（OCR、裁剪、扫码） |
| `security.go` | 内容安全检测 |
| `service_market.go` | 服务市场调用 |
| `open_data.go` | 开放数据 |
| `mini_drama_manage.go` | 短剧媒资管理 |
| `mini_drama_upload.go` | 短剧上传 |

### public/ 模块文件说明

| 文件 | 用途 |
|------|------|
| `public.go` | SDK 结构体定义、初始化、包级别请求函数 |
| `request.go` | SDK 实例级请求方法（doRequestGet/doRequestPost/doRequestUpload/doRequestUploadWithForm） |
| `model.go` | 所有请求/响应结构体 |
| `access_token.go` | AccessToken 管理（稳定版接口、自动刷新） |
| `account_manage.go` | 二维码/短链接管理 |
| `basic.go` | 基础 API（IP 列表、配额、RID 查询） |
| `menu.go` | 自定义菜单管理 |
| `material.go` | 素材管理（临时/永久素材、图文） |
| `user_manage.go` | 用户管理（标签、粉丝列表、黑名单） |
| `ticket.go` | Ticket 获取（jsapi_ticket、api_ticket） |
| `customer_service.go` | 客服管理（账号、会话、消息、历史） |
| `sign.go` | 签名工具（JS-SDK 签名） |
| `ai.go` | AI 能力（语音识别、图像处理、OCR） |
| `statistics.go` | 数据统计（用户、文章、消息、接口） |
| `publish.go` | 草稿/发布管理 |

### open/ 模块文件说明

| 文件 | 用途 |
|------|------|
| `open.go` | SDK 结构体（含 `openidAccessTokenMap smap.Map`）、DoRequestGet |
| `model.go` | AT、AccessToken、UserInfo、ErrorCode 结构体 |
| `access_token.go` | 多用户 token 管理（Code2AccessToken、RefreshAccessToken、自动刷新、回调） |
| `user_info.go` | 用户信息获取（UnionID 机制） |

### 公共配置（common.go）

- `Version = "v1.1.10"` - 当前版本
- `Success = 0` - 成功码
- `HeaderRequestID = "Request-ID"` - 请求追踪头
- Host 映射：HostDefault(`api.weixin.qq.com`)、HostDefault2、HostSH(上海)、HostSZ(深圳)、HostHK(香港)
- 类型定义：`DebugSwitch int8`、`Host int`、`Platform string`

### 依赖库

核心依赖来自 `go-pay` 生态（共 7 个）：
- `go-pay/xhttp v0.0.3` - HTTP 客户端
- `go-pay/xlog v0.0.3` - 日志库
- `go-pay/bm v0.0.5` - BodyMap 请求体构建
- `go-pay/util v0.0.4` - 工具函数（加密、随机字符串等）
- `go-pay/smap v0.0.2` - 线程安全 Map（开放平台用于管理多用户 token）
- `go-pay/crypto v0.0.1` - 加密操作
- `go-pay/xtime v0.0.2` - 时间工具

## 开发规范

### 新增接口实现

1. **接口定义**：在对应模块的 `model.go` 中定义请求和响应结构体
2. **常量定义**：如需要，在模块内定义 API 路径常量
3. **方法实现**：
   - 方法接收器为 `(s *SDK)`
   - 第一个参数为 `ctx context.Context`
   - 使用 `doRequestGet` 或 `doRequestPost` 发起请求
   - 检查响应的 `errcode`，非 0 时返回 error
4. **测试文件**：在 `*_test.go` 中添加测试用例
5. **文档更新**：在 `doc/` 目录对应的 `.md` 文件中更新 API 列表

### 标准 API 实现模板

```go
// MethodName 方法描述
// 注意：errcode = 0 为成功
// param1：参数说明
// 文档：https://developers.weixin.qq.com/...
func (s *SDK) MethodName(c context.Context, param1 string) (result *ResultRsp, err error) {
    path := "/api/endpoint?access_token=" + s.accessToken
    body := make(bm.BodyMap)
    body.Set("param1", param1)
    result = &ResultRsp{}
    if _, err = s.DoRequestPost(c, path, body, result); err != nil {
        return nil, err
    }
    if result.Errcode != Success {
        return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
    }
    return result, nil
}
```

### 响应结构体规范

```go
type ResultRsp struct {
    Errcode int    `json:"errcode"`           // 必须包含
    Errmsg  string `json:"errmsg"`            // 必须包含
    Data    *Info  `json:"data,omitempty"`     // 业务字段用 omitempty
}
```

### 只返回错误的接口（无业务数据）

```go
func (s *SDK) MethodName(c context.Context, param1 string) (err error) {
    path := "/api/endpoint?access_token=" + s.accessToken
    body := make(bm.BodyMap)
    body.Set("param1", param1)
    ec := &ErrorCode{}
    if _, err = s.DoRequestPost(c, path, body, ec); err != nil {
        return err
    }
    if ec.Errcode != Success {
        return fmt.Errorf("errcode(%d), errmsg(%s)", ec.Errcode, ec.Errmsg)
    }
    return nil
}
```

### mini 模块请求方法说明

- **包级别函数**（mini.go）：`doRequestGet(ctx, uri, ptr)` / `doRequestPost(ctx, url, body, ptr)` — 用于无需 SDK 实例的静态调用（如获取 AccessToken）
- **SDK 实例方法**（request.go）：
   - `DoRequestGet(ctx, path, ptr)` — 自动拼接 Host + path，带 Debug 日志
   - `DoRequestPost(ctx, path, body, ptr)` — POST JSON 请求
   - `DoRequestPostFile(ctx, path, body, ptr)` — multipart 文件上传
   - `DoRequestGetByte(ctx, path)` — 获取二进制响应（图片等）
   - `doRequestPostForImage(ctx, path, body)` — POST 后判断 Content-Type 返回图片或错误（qrcode.go）

### public 模块请求方法说明

- **包级别函数**（public.go）：同 mini，`doRequestGet` / `doRequestPost`
- **SDK 实例方法**（request.go）：
   - `doRequestGet` / `doRequestPost` — 小写开头（包内使用）
   - `doRequestUpload(ctx, path, fieldName, file, ptr)` — 文件上传
   - `doRequestUploadWithForm(ctx, path, fieldName, file, formField, formBody, ptr)` — 文件+表单混合上传

### open 模块请求方法说明

- 只有 `DoRequestGet(ctx, path, ptr)` — 开放平台 API 主要是 GET 请求
- Token 通过 URL 参数传递，不存储在 SDK 实例的单一字段中，而是 per-openid 管理

### AccessToken 管理

- **小程序和公众号**：使用稳定版接口 `getStableAccessToken()`，自动刷新（刷新间隔 = ExpiresIn/2）
- **开放平台**：支持多用户 token 管理，使用 `openidAccessTokenMap` 存储，每 10 分钟轮询刷新
- 回调机制：通过 `SetMiniAccessTokenCallback` / `SetPublicAccessTokenCallback` / `SetAccessTokenCallback` 设置 token 刷新回调

### 请求规范

所有 HTTP 请求必须：
- 添加 `Request-ID` header（格式：`{21位随机字符串}-{时间戳}`）
- 使用 `js.UnmarshalBytes` 解析响应
- Debug 模式下记录请求 URI 和响应内容

### 版本发布

更新版本时需要同步修改：
1. `common.go` 中的 `Version` 常量
2. `release_note.md` 添加版本记录
3. Git tag 格式：`v1.1.x`

## 参考文档

- 微信小程序服务端：https://developers.weixin.qq.com/miniprogram/dev/api-backend/
- 微信公众号开发文档：https://developers.weixin.qq.com/doc/subscription/api/
- 微信开放平台文档：https://developers.weixin.qq.com/doc/oplatform/
- 项目文档：`doc/mini.md`、`doc/public.md`、`doc/open.md`
- API 覆盖率文档：`doc/mini_api_coverage.md`、`doc/public_api_coverage.md`
