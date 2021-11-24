package mini

import (
	"time"

	"github.com/go-pay/wechat-sdk/model"
)

type SDK struct {
	Conf        *model.Config
	DebugSwitch int8
	AccessChan  chan string
}

func New(c *model.Config, ds int8, accessChan chan string) (m *SDK) {
	m = &SDK{
		Conf:        c,
		DebugSwitch: ds,
		AccessChan:  accessChan,
	}
	go m.accessTokenListener()
	return
}

func (s *SDK) accessTokenListener() {
	defer func() {
		time.Sleep(time.Second)
		close(s.AccessChan)
	}()
	for {
		at := <-s.AccessChan
		s.Conf.AccessToken = at
	}
}