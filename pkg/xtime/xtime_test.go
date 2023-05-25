package xtime

import (
	"encoding/json"
	"testing"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

type TimeParser struct {
	T1 Duration `json:"t1"`
	T2 Duration `json:"t2"`
	T3 Duration `json:"t3"`
	T4 Duration `json:"t4"`
}

func TestParseTime(t *testing.T) {
	parseText := `
{
    "t1":"10s",
    "t2":"1m20s",
    "t3":"5m",
    "t4":"1h10m10s"
}`
	tp := new(TimeParser)
	err := json.Unmarshal([]byte(parseText), tp)
	if err != nil {
		xlog.Error(err)
		return
	}

	xlog.Infof("%+v", tp)

	xlog.Debugf("t1: %s", tp.T1.UnitTime())
	xlog.Debugf("t2: %s", tp.T2.UnitTime())
	xlog.Debugf("t3: %s", tp.T3.UnitTime())
	xlog.Debugf("t4: %s", tp.T4.UnitTime())
}
