package public

import (
	"context"
	"fmt"

	"github.com/go-pay/bm"
)

// GetUserSummary 获取用户增减数据
// beginDate：获取数据的起始日期，begin_date和end_date的差值需小于"最大时间跨度"（比如最大时间跨度为1时，begin_date和end_date的差值只能为0，才能小于1），否则会报错
// endDate：获取数据的结束日期，end_date允许设置的最大值为昨日
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html
func (s *SDK) GetUserSummary(c context.Context, beginDate, endDate string) (result *UserSummaryRsp, err error) {
	path := "/datacube/getusersummary?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UserSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUserCumulate 获取累计用户数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html
func (s *SDK) GetUserCumulate(c context.Context, beginDate, endDate string) (result *UserCumulateRsp, err error) {
	path := "/datacube/getusercumulate?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UserCumulateRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetArticleSummary 获取图文群发每日数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (s *SDK) GetArticleSummary(c context.Context, beginDate, endDate string) (result *ArticleSummaryRsp, err error) {
	path := "/datacube/getarticlesummary?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &ArticleSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetArticleTotal 获取图文群发总数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (s *SDK) GetArticleTotal(c context.Context, beginDate, endDate string) (result *ArticleSummaryRsp, err error) {
	path := "/datacube/getarticletotal?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &ArticleSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUserRead 获取图文统计数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (s *SDK) GetUserRead(c context.Context, beginDate, endDate string) (result *ArticleSummaryRsp, err error) {
	path := "/datacube/getuserread?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &ArticleSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUserReadHour 获取图文统计分时数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (s *SDK) GetUserReadHour(c context.Context, beginDate, endDate string) (result *ArticleSummaryRsp, err error) {
	path := "/datacube/getuserreadhour?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &ArticleSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUserShare 获取图文分享转发数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (s *SDK) GetUserShare(c context.Context, beginDate, endDate string) (result *ArticleSummaryRsp, err error) {
	path := "/datacube/getusershare?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &ArticleSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUserShareHour 获取图文分享转发分时数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (s *SDK) GetUserShareHour(c context.Context, beginDate, endDate string) (result *ArticleSummaryRsp, err error) {
	path := "/datacube/getusersharehour?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &ArticleSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUpstreamMsg 获取消息发送概况数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (s *SDK) GetUpstreamMsg(c context.Context, beginDate, endDate string) (result *UpstreamMsgRsp, err error) {
	path := "/datacube/getupstreammsg?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UpstreamMsgRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUpstreamMsgHour 获取消息发送分时数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (s *SDK) GetUpstreamMsgHour(c context.Context, beginDate, endDate string) (result *UpstreamMsgRsp, err error) {
	path := "/datacube/getupstreammsghour?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UpstreamMsgRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUpstreamMsgWeek 获取消息发送周数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (s *SDK) GetUpstreamMsgWeek(c context.Context, beginDate, endDate string) (result *UpstreamMsgRsp, err error) {
	path := "/datacube/getupstreammsgweek?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UpstreamMsgRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUpstreamMsgMonth 获取消息发送月数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (s *SDK) GetUpstreamMsgMonth(c context.Context, beginDate, endDate string) (result *UpstreamMsgRsp, err error) {
	path := "/datacube/getupstreammsgmonth?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UpstreamMsgRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUpstreamMsgDist 获取消息发送分布数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (s *SDK) GetUpstreamMsgDist(c context.Context, beginDate, endDate string) (result *UpstreamMsgRsp, err error) {
	path := "/datacube/getupstreammsgdist?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UpstreamMsgRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUpstreamMsgDistWeek 获取消息发送分布周数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (s *SDK) GetUpstreamMsgDistWeek(c context.Context, beginDate, endDate string) (result *UpstreamMsgRsp, err error) {
	path := "/datacube/getupstreammsgdistweek?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UpstreamMsgRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetUpstreamMsgDistMonth 获取消息发送分布月数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (s *SDK) GetUpstreamMsgDistMonth(c context.Context, beginDate, endDate string) (result *UpstreamMsgRsp, err error) {
	path := "/datacube/getupstreammsgdistmonth?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &UpstreamMsgRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetInterfaceSummary 获取接口分析数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html
func (s *SDK) GetInterfaceSummary(c context.Context, beginDate, endDate string) (result *InterfaceSummaryRsp, err error) {
	path := "/datacube/getinterfacesummary?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &InterfaceSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}

// GetInterfaceSummaryHour 获取接口分析分时数据
// beginDate：获取数据的起始日期
// endDate：获取数据的结束日期
// 文档：https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html
func (s *SDK) GetInterfaceSummaryHour(c context.Context, beginDate, endDate string) (result *InterfaceSummaryRsp, err error) {
	path := "/datacube/getinterfacesummaryhour?access_token=" + s.accessToken
	body := make(bm.BodyMap)
	body.Set("begin_date", beginDate)
	body.Set("end_date", endDate)

	result = &InterfaceSummaryRsp{}
	if _, err = s.doRequestPost(c, path, body, result); err != nil {
		return nil, err
	}
	if result.Errcode != Success {
		return nil, fmt.Errorf("errcode(%d), errmsg(%s)", result.Errcode, result.Errmsg)
	}
	return result, nil
}
