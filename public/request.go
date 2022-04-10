package public

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-pay/wechat-sdk/model"
	"github.com/go-pay/wechat-sdk/pkg/bmap"
	"github.com/go-pay/wechat-sdk/pkg/util"
	"github.com/go-pay/wechat-sdk/pkg/xhttp"
	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

func (s *SDK) doRequestGet(c context.Context, path string, ptr interface{}) (err error) {
	uri := s.Conf.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_URI: %s", uri)
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}

func (s *SDK) doRequestGetByte(c context.Context, path string) (bs []byte, err error) {
	uri := s.Conf.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_URI: %s", uri)
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	ec := &model.ErrorCode{}
	// 如果解析成功，说明获取buffer文件失败
	if err = json.Unmarshal(bs, ec); err == nil {
		return nil, fmt.Errorf("errcode(%d)，errmsg(%s)", ec.Errcode, ec.Errmsg)
	}
	return
}

func (s *SDK) doRequestPost(c context.Context, path string, body bmap.BodyMap, ptr interface{}) (err error) {
	uri := s.Conf.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_URI: %s", uri)
		xlog.Debugf("Wechat_Open_SDK_RequestBody: %s", body.JsonBody())
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Post(uri).SendBodyMap(body).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(POST, %s)：%w", uri, err)
	}
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}

func (s *SDK) doRequestPostFile(ctx context.Context, path string, body bmap.BodyMap, ptr interface{}) (err error) {
	uri := s.Conf.Host + path
	httpClient := xhttp.NewClient()
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_URI: %s", uri)
	}
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Type(xhttp.TypeMultipartFormData).Post(uri).SendMultipartBodyMap(body).EndBytes(ctx)
	if err != nil {
		return fmt.Errorf("http.request(POST, %s)：%w", uri, err)
	}
	if s.DebugSwitch == model.DebugOn {
		xlog.Debugf("Wechat_Open_SDK_Response: [%d] -> %s", res.StatusCode, string(bs))
	}
	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}
