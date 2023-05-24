package public

import (
	"context"
	"fmt"
)

// GetJsApiTicket 获取 jsapi_ticket
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html#62
func (s *SDK) GetJsApiTicket(c context.Context) (jt *TicketRsp, err error) {
	path := "/cgi-bin/ticket/getticket?access_token=" + s.accessToken + "&type=jsapi"
	jt = &TicketRsp{}
	if err = s.doRequestGet(c, path, jt); err != nil {
		return nil, err
	}
	if jt.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", jt.Errcode, jt.Errmsg)
	}
	return jt, nil
}

// GetApiTicket 获取卡券 api_ticket
// 注意：errcode = 0 为成功
// 文档：https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html#54
func (s *SDK) GetApiTicket(c context.Context) (at *TicketRsp, err error) {
	// /cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=wx_card
	path := "/cgi-bin/ticket/getticket?access_token=" + s.accessToken + "&type=wx_card"
	at = &TicketRsp{}
	if err = s.doRequestGet(c, path, at); err != nil {
		return nil, err
	}
	if at.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", at.Errcode, at.Errmsg)
	}
	return at, nil
}
