package public

import (
	"testing"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

func TestJsSDKUsePermissionSign(t *testing.T) {
	var (
		ticket    = "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg"
		nonce     = "Wm3WZYTPz0wzccnW"
		timestamp = 1414587457
		url       = "http://mp.weixin.qq.com?params=value"
	)

	sign := JsSDKUsePermissionSign(ticket, nonce, url, timestamp)
	xlog.Infof("sign: %s", sign)
}
