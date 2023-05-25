package wechat

type DebugSwitch int8
type Host int
type Platform string

const (
	Success = 0

	DebugOff = 0
	DebugOn  = 1

	Version = "1.1.5"
)

const (
	HostDefault Host = iota + 1
	HostDefault2
	HostSH // 上海
	HostSZ // 深圳
	HostHK // 香港
)

var (
	HostMap = map[Host]string{
		HostDefault:  "https://api.weixin.qq.com",
		HostDefault2: "https://api2.weixin.qq.com",
		HostSH:       "https://sh.api.weixin.qq.com",
		HostSZ:       "https://sz.api.weixin.qq.com",
		HostHK:       "https://hk.api.weixin.qq.com",
	}
)
