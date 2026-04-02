# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

这是一个微信 Golang SDK，提供微信小程序、公众号、开放平台三个模块的服务端 API 封装。

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

### 公共配置

- `common.go` - 定义全局常量、版本号、Host 映射
- 支持多个 API Host：默认、上海、深圳、香港
- 统一错误处理：errcode != 0 时返回 error

### 依赖库

核心依赖来自 `go-pay` 生态：
- `go-pay/xhttp` - HTTP 客户端
- `go-pay/xlog` - 日志库
- `go-pay/bm` - BodyMap 请求体构建
- `go-pay/util` - 工具函数（加密、随机字符串等）
- `go-pay/smap` - 线程安全 Map（开放平台用于管理多用户 token）

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

### AccessToken 管理

- 小程序和公众号：使用稳定版接口 `getStableAccessToken()`，自动刷新
- 开放平台：支持多用户 token 管理，使用 `openidAccessTokenMap` 存储
- 回调机制：通过 `SetMiniAccessTokenCallback` 等方法设置 token 刷新回调

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
- 微信公众号开发文档：https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html
- 微信开放平台文档：https://developers.weixin.qq.com/doc/oplatform/
- 项目文档：`doc/mini.md`、`doc/public.md`、`doc/open.md`
