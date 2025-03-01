## 版本号：Release 1.1.7

* 修改记录：
  * (1) 微信小程序：更新获取token的方法。
  * (2) 微信小程序：新增媒资相关接口。

## 版本号：Release 1.1.6

* 修改记录：
  * (1) big change

## 版本号：Release 1.1.5

* 修改记录：
  * (1) 微信公众号：新增 ticket获取，js-sdk使用校验签名接口
  * (2) 微信小程序：自动维护token改为获取稳定版接口调用凭据，并且 forceRefresh=false
  * (3) 微信小程序：新增公共API 获取接口调用凭据、获取稳定版接口调用凭据
  * (4) go mod 升级至 1.21

## 版本号：Release 1.1.4

* 修改记录：
  * (1) 修复开放平台问题，重新设计开放平台相关接口以及token维护机制
  * (2) 小程序、公众号平台，AccessToken回调增加 appid 返回

## 版本号：Release 1.1.1

* 修改记录：
  * (1) 修改sdk初始化流程入参，appid 和appkey 分别单独设置

## 版本号：Release 1.1.0

* 修改记录：
  * (1) 比较大的更新，区分了 mini（小程序）、open（开放平台）、public（公众号）三个平台
  * (2) 接口增加 errcode 是否成功的预判的处理，出错时直接返回 error

## 版本号：Release 1.0.0

* 修改记录：
  * (1) 首次 release
