package public

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/bm"
	"github.com/go-pay/util"
	"github.com/go-pay/util/js"
	"github.com/go-pay/wechat-sdk"
)

func (s *SDK) doRequestGet(c context.Context, path string, ptr any) (res *http.Response, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Public_SDK_URI: %s", uri)
	}
	req := s.hc.Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Get(uri).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(GET, %s), err:%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Public_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return res, fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return res, nil
}

func (s *SDK) doRequestPost(c context.Context, path string, body bm.BodyMap, ptr any) (res *http.Response, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Public_SDK_URI: %s", uri)
		s.logger.Debugf("Wechat_Public_SDK_RequestBody: %s", body.JsonBody())
	}
	req := s.hc.Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Post(uri).SendBodyMap(body).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(POST, %s), err:%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Public_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return res, fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return res, nil
}
