package mini

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pay/bm"
	"github.com/go-pay/util"
	"github.com/go-pay/util/js"
	"github.com/go-pay/wechat-sdk"
	"github.com/go-pay/xhttp"
)

func (s *SDK) DoRequestGet(c context.Context, path string, ptr any) (res *http.Response, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_URI: %s", uri)
	}
	req := s.hc.Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Get(uri).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(GET, %s), err:%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return res, fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return res, nil
}

func (s *SDK) DoRequestGetByte(c context.Context, path string) (bs []byte, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_URI: %s", uri)
	}
	req := s.hc.Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Get(uri).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(GET, %s), err:%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	ec := &ErrorCode{}
	// 如果解析成功，说明获取buffer文件失败
	if err = json.Unmarshal(bs, ec); err == nil {
		return nil, fmt.Errorf("errcode(%d)，errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return
}

func (s *SDK) DoRequestPost(c context.Context, path string, body bm.BodyMap, ptr any) (res *http.Response, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_URI: %s", uri)
		s.logger.Debugf("Wechat_Mini_SDK_RequestBody: %s", body.JsonBody())
	}
	req := s.hc.Req()
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Post(uri).SendBodyMap(body).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(POST, %s), err:%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return res, fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}

func (s *SDK) DoRequestPostFile(ctx context.Context, path string, body bm.BodyMap, ptr any) (res *http.Response, err error) {
	uri := s.Host + path
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_URI: %s", uri)
	}
	req := s.hc.Req(xhttp.TypeMultipartFormData)
	req.Header.Add(wechat.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := req.Post(uri).SendMultipartBodyMap(body).EndBytes(ctx)
	if err != nil {
		return nil, fmt.Errorf("http.request(POST, %s), err:%w", uri, err)
	}
	if s.DebugSwitch == wechat.DebugOn {
		s.logger.Debugf("Wechat_Mini_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = js.UnmarshalBytes(bs, ptr); err != nil {
		return res, fmt.Errorf("js.UnmarshalBytes(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
