# Wechat-SDK

### 微信 Golang 版本SDK

[![Github](https://img.shields.io/github/followers/iGoogle-ink?label=Follow&style=social)](https://github.com/iGoogle-ink)
[![Github](https://img.shields.io/github/forks/go-pay/wechat-sdk?label=Fork&style=social)](https://github.com/go-pay/wechat-sdk/fork)

[![Golang](https://img.shields.io/badge/golang-1.21-brightgreen.svg)](https://golang.google.cn)
[![GoDoc](https://img.shields.io/badge/doc-pkg.go.dev-informational.svg)](https://pkg.go.dev/github.com/go-pay/wechat-sdk)
[![Go](https://github.com/go-pay/wechat-sdk/actions/workflows/go.yml/badge.svg)](https://github.com/go-pay/wechat-sdk/actions/workflows/go.yml)
[![GitHub Release](https://img.shields.io/github/v/release/go-pay/wechat-sdk)](https://github.com/go-pay/wechat-sdk/releases)
[![License](https://img.shields.io/github/license/go-pay/wechat-sdk)](https://www.apache.org/licenses/LICENSE-2.0)

<br>

# 一、安装

```bash
go get -u github.com/go-pay/wechat-sdk
```

#### 查看 Wechat-SDK 版本

[版本更新记录](https://github.com/go-pay/wechat-sdk/blob/main/release_note.txt)

```go
import (
    "github.com/go-pay/wechat-sdk"
    "github.com/go-pay/xlog"
)

func main() {
    xlog.Warn("Wechat-SDK Version: ", wechat.Version)
}
```

---

<br>

# 二、文档说明

- ### 点击分别查看小程序、公众号使用文档

  * #### [微信小程序](https://github.com/go-pay/wechat-sdk/blob/main/doc/mini.md)
  * #### [微信公众号](https://github.com/go-pay/wechat-sdk/blob/main/doc/public.md)
  * #### [微信开放平台](https://github.com/go-pay/wechat-sdk/blob/main/doc/open.md)

---

<br>

# 三、其他说明

* 请仔细查看参考各个平台下 `*_test.go` 使用方式
* 有问题请加QQ群（加群验证答案：gopay），或加微信好友拉群。在此，非常感谢提出宝贵意见和反馈问题的同志们！

QQ群：
<img width="280" height="280" src=".github/qq_gopay.png"/>
加微信拉群：
<img width="280" height="280" src=".github/wechat_jerry.png"/>

---

<br>

## 赞赏多少是您的心意，感谢支持！

微信赞赏码： <img width="240" height="240" src=".github/zanshang.png"/>
支付宝赞助码： <img width="240" height="240" src=".github/zanshang_zfb.png"/>

---

<br>

## 鸣谢

> [GoLand](https://www.jetbrains.com/go/?from=gopay) A Go IDE with extended support for JavaScript, TypeScript, and databases。
> 
特别感谢 [JetBrains](https://www.jetbrains.com/?from=gopay) 为开源项目提供免费的 [GoLand](https://www.jetbrains.com/go/?from=gopay) 等 IDE 的授权  
[<img src=".github/jetbrains-variant-3.png" width="200"/>](https://www.jetbrains.com/?from=gopay)
