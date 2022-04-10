# Wechat-SDK

### 微信 Golang 版本SDK

[![Github](https://img.shields.io/github/followers/iGoogle-ink?label=Follow&style=social)](https://github.com/iGoogle-ink)
[![Github](https://img.shields.io/github/forks/go-pay/wechat-sdk?label=Fork&style=social)](https://github.com/go-pay/wechat-sdk/fork)

[![Golang](https://img.shields.io/badge/golang-1.16-brightgreen.svg)](https://golang.google.cn)
[![GoDoc](https://img.shields.io/badge/doc-pkg.go.dev-informational.svg)](https://pkg.go.dev/github.com/go-pay/wechat-sdk)
[![Drone CI](https://cloud.drone.io/api/badges/go-pay/wechat-sdk/status.svg)](https://cloud.drone.io/go-pay/wechat-sdk)
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
    "github.com/go-pay/wechat-sdk/pkg/xlog"
)

func main() {
    xlog.Info("Wechat-SDK Version: ", wechat.Version)
}
```

---

<br>

# 二、文档说明

- ### NewSDK

```go
import (
    "github.com/go-pay/wechat-sdk"
    "github.com/go-pay/wechat-sdk/pkg/xlog"
)

// NewSDK 初始化微信 SDK
//  Appid：Appid
//  Secret：appSecret
//  accessToken：AccessToken，若此参数为空，则自动获取并自动维护刷新
wxsdk, err := wechat.NewSDK(Appid, Secret)
if err != nil {
    xlog.Error(err)
    return
}

// 可替换host节点
//wxsdk.SetHost(wechat.HostSH)
// 打开Debug开关，输出日志
//wxsdk.DebugSwitch = wechat.DebugOn
```

- ### AccessToken 说明

```go
// NewSDK 后，首次获取AccessToken请通过此方法获取，之后请通过下面的回调方法获取
at := wxsdk.GetAccessToken()
xlog.Infof("at: %s", at)

// 每次刷新 accessToken 后，此方法回调返回 accessToken 和 有效时间（秒）
wxsdk.SetAccessTokenCallback(func(accessToken string, expireIn int, err error) {
    if err != nil {
        xlog.Errorf("refresh access token error(%+v)", err)
    }
    xlog.Infof("accessToken: %s", accessToken)
    xlog.Infof("expireIn: %d", expireIn)
})

// 若 NewSDK() 时自传 AccessToken，则后续更新替换请调用此方法
wxsdk.SetAccessToken()
```

- ### NewMiniSDK

```go
// New 微信小程序 SDK
miniSDK := wxsdk.NewMini()
```

- ### NewPublicSDK

```go
// New 微信公众号 SDK
publicSDK := wxsdk.NewPublic()
```

- ### 点击分别查看小程序、公众号使用文档

  * #### [微信小程序](https://github.com/go-pay/wechat-sdk/blob/main/doc/mini.md)
  * #### [微信公众号](https://github.com/go-pay/wechat-sdk/blob/main/doc/public.md)
  * #### [微信开放平台](https://github.com/go-pay/wechat-sdk/blob/main/doc/open.md)

---

<br>

# 三、其他说明

* 请仔细查看参考 `sdk_test.go` 使用方式
* 有问题请加QQ群（加群验证答案：gopay），或加微信好友拉群。在此，非常感谢提出宝贵意见和反馈问题的同志们！

QQ群：
<img width="280" height="280" src="https://raw.githubusercontent.com/go-pay/wechat-sdk/main/qq_gopay.png"/>
加微信拉群：
<img width="280" height="280" src="https://raw.githubusercontent.com/go-pay/wechat-sdk/main/wechat_jerry.png"/>

---

<br>

## 赞赏多少是您的心意，感谢支持！

微信赞赏码： <img width="240" height="240" src="https://raw.githubusercontent.com/go-pay/gopay/main/zanshang.png"/>
支付宝赞助码： <img width="240" height="240" src="https://raw.githubusercontent.com/go-pay/gopay/main/zanshang_zfb.png"/>

## License

```
Copyright 2021 Jerry

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```