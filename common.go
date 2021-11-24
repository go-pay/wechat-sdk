package wechat

const (
	Success  = 0
	DebugOff = 0
	DebugOn  = 1
	Version  = "1.0.0"
)

const (
	HostDefault Host = iota
	HostDefault2
	HostSH
	HostSZ
	HostHK
)

type DebugSwitch int8

type Host int

var (
	HostMap = map[Host]string{
		HostDefault:  "https://api.weixin.qq.com",
		HostDefault2: "https://api2.weixin.qq.com",
		HostSH:       "https://sh.api.weixin.qq.com",
		HostSZ:       "https://sz.api.weixin.qq.com",
		HostHK:       "https://hk.api.weixin.qq.com",
	}
)
