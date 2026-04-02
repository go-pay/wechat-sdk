package public

import "github.com/go-pay/xtime"

const (
	Success = 0

	HostDefault = "https://api.weixin.qq.com"
)

type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in,omitempty"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	Errcode     int    `json:"errcode,omitempty"`      // 错误码
	Errmsg      string `json:"errmsg,omitempty"`       // 错误信息
}

type ErrorCode struct {
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type QRCodeRsp struct {
	Errcode       int    `json:"errcode,omitempty"`
	Errmsg        string `json:"errmsg,omitempty"`
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds,omitempty"`
	Url           string `json:"url"`
}

type ShortKeyGenRsp struct {
	Errcode  int    `json:"errcode,omitempty"`
	Errmsg   string `json:"errmsg,omitempty"`
	ShortKey string `json:"short_key"`
}

type ShortKeyFetchRsp struct {
	Errcode       int        `json:"errcode,omitempty"`
	Errmsg        string     `json:"errmsg,omitempty"`
	LongData      string     `json:"long_data"`
	CreateTime    xtime.Time `json:"create_time"`
	ExpireSeconds int        `json:"expire_seconds"`
}

type UserTagRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Tag     *Tag   `json:"tag,omitempty"`
}

type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"` // 此标签下粉丝数
}

type UserTagListRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Tags    []*Tag `json:"tags,omitempty"`
}
type UserTagFansListRsp struct {
	Errcode    int       `json:"errcode,omitempty"`
	Errmsg     string    `json:"errmsg,omitempty"`
	Count      int       `json:"count"`
	NextOpenid string    `json:"next_openid,omitempty"`
	Data       *FansData `json:"data"`
}

type FansData struct {
	Openid []string `json:"openid"`
}

type UserTagIdListRsp struct {
	Errcode   int    `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	TagidList []int  `json:"tagid_list,omitempty"`
}

type TicketRsp struct {
	Errcode   int    `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	Ticket    string `json:"ticket,omitempty"`
	ExpiresIn int    `json:"expires_in,omitempty"` // 有效期7200秒，开发者必须在自己的服务全局缓存jsapi_ticket
}

// ============ 自定义菜单 ============

type MenuGetRsp struct {
	Errcode         int                `json:"errcode,omitempty"`
	Errmsg          string             `json:"errmsg,omitempty"`
	Menu            *MenuInfo          `json:"menu,omitempty"`
	Conditionalmenu []*ConditionalMenu `json:"conditionalmenu,omitempty"`
}

type MenuInfo struct {
	Button []map[string]interface{} `json:"button"`
	MenuId int64                    `json:"menuid,omitempty"`
}

type ConditionalMenu struct {
	Button    []map[string]interface{} `json:"button"`
	Matchrule map[string]interface{}   `json:"matchrule"`
	MenuId    int64                    `json:"menuid"`
}

type MenuAddConditionalRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	MenuId  string `json:"menuid,omitempty"`
}

type MenuTryMatchRsp struct {
	Errcode int                      `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Button  []map[string]interface{} `json:"button,omitempty"`
}

type CurrentSelfMenuInfoRsp struct {
	Errcode      int                    `json:"errcode,omitempty"`
	Errmsg       string                 `json:"errmsg,omitempty"`
	IsMenuOpen   int                    `json:"is_menu_open"`
	SelfMenuInfo map[string]interface{} `json:"selfmenu_info,omitempty"`
}

// ============ 素材管理 ============

type MediaUploadRsp struct {
	Errcode   int    `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	Type      string `json:"type,omitempty"`
	MediaId   string `json:"media_id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

type MediaGetRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type MaterialAddNewsRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	MediaId string `json:"media_id,omitempty"`
}

type MaterialUploadImgRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Url     string `json:"url,omitempty"`
}

type MaterialAddMaterialRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	MediaId string `json:"media_id,omitempty"`
	Url     string `json:"url,omitempty"`
}

type MaterialGetMaterialRsp struct {
	Errcode  int         `json:"errcode,omitempty"`
	Errmsg   string      `json:"errmsg,omitempty"`
	NewsItem []*NewsItem `json:"news_item,omitempty"`
	DownUrl  string      `json:"down_url,omitempty"`
}

type NewsItem struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"content_source_url"`
}

type MaterialGetMaterialCountRsp struct {
	Errcode    int    `json:"errcode,omitempty"`
	Errmsg     string `json:"errmsg,omitempty"`
	VoiceCount int    `json:"voice_count"`
	VideoCount int    `json:"video_count"`
	ImageCount int    `json:"image_count"`
	NewsCount  int    `json:"news_count"`
}

type MaterialBatchGetMaterialRsp struct {
	Errcode    int             `json:"errcode,omitempty"`
	Errmsg     string          `json:"errmsg,omitempty"`
	TotalCount int             `json:"total_count"`
	ItemCount  int             `json:"item_count"`
	Item       []*MaterialItem `json:"item,omitempty"`
}

type MaterialItem struct {
	MediaId    string                 `json:"media_id"`
	Name       string                 `json:"name,omitempty"`
	UpdateTime int64                  `json:"update_time"`
	Url        string                 `json:"url,omitempty"`
	Content    map[string]interface{} `json:"content,omitempty"`
}

// ============ 客服消息 ============

type KfAccountAddRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type KfAccountListRsp struct {
	Errcode int          `json:"errcode,omitempty"`
	Errmsg  string       `json:"errmsg,omitempty"`
	KfList  []*KfAccount `json:"kf_list,omitempty"`
}

type KfAccount struct {
	KfAccount        string `json:"kf_account"`
	KfHeadimgurl     string `json:"kf_headimgurl,omitempty"`
	KfId             string `json:"kf_id,omitempty"`
	KfNick           string `json:"kf_nick,omitempty"`
	KfWx             string `json:"kf_wx,omitempty"`
	InviteWx         string `json:"invite_wx,omitempty"`
	InviteExpireTime int64  `json:"invite_expire_time,omitempty"`
	InviteStatus     string `json:"invite_status,omitempty"`
}

type KfSessionCreateRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type KfSessionGetRsp struct {
	Errcode    int    `json:"errcode,omitempty"`
	Errmsg     string `json:"errmsg,omitempty"`
	KfAccount  string `json:"kf_account,omitempty"`
	CreateTime int64  `json:"createtime,omitempty"`
}

type KfSessionListRsp struct {
	Errcode     int            `json:"errcode,omitempty"`
	Errmsg      string         `json:"errmsg,omitempty"`
	SessionList []*SessionInfo `json:"sessionlist,omitempty"`
}

type SessionInfo struct {
	KfAccount  string `json:"kf_account"`
	Openid     string `json:"openid"`
	CreateTime int64  `json:"createtime"`
}

type KfSessionWaitCaseRsp struct {
	Errcode  int         `json:"errcode,omitempty"`
	Errmsg   string      `json:"errmsg,omitempty"`
	Count    int         `json:"count"`
	WaitCase []*WaitCase `json:"waitcaselist,omitempty"`
}

type WaitCase struct {
	LatestTime int64  `json:"latest_time"`
	Openid     string `json:"openid"`
}

type MsgRecordListRsp struct {
	Errcode    int          `json:"errcode,omitempty"`
	Errmsg     string       `json:"errmsg,omitempty"`
	RecordList []*MsgRecord `json:"recordlist,omitempty"`
	Number     int          `json:"number"`
	Msgid      int64        `json:"msgid"`
}

type MsgRecord struct {
	Worker   string `json:"worker"`
	Openid   string `json:"openid"`
	OperCode int    `json:"opercode"`
	Time     int64  `json:"time"`
	Text     string `json:"text"`
}

// ============ 用户管理 ============

type UserInfoRsp struct {
	Errcode        int    `json:"errcode,omitempty"`
	Errmsg         string `json:"errmsg,omitempty"`
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Language       string `json:"language,omitempty"`
	SubscribeTime  int64  `json:"subscribe_time,omitempty"`
	Unionid        string `json:"unionid,omitempty"`
	Remark         string `json:"remark,omitempty"`
	Groupid        int    `json:"groupid,omitempty"`
	TagidList      []int  `json:"tagid_list,omitempty"`
	SubscribeScene string `json:"subscribe_scene,omitempty"`
	QrScene        int    `json:"qr_scene,omitempty"`
	QrSceneStr     string `json:"qr_scene_str,omitempty"`
}

type UserInfoBatchRsp struct {
	Errcode      int            `json:"errcode,omitempty"`
	Errmsg       string         `json:"errmsg,omitempty"`
	UserInfoList []*UserInfoRsp `json:"user_info_list,omitempty"`
}

type UserListRsp struct {
	Errcode    int       `json:"errcode,omitempty"`
	Errmsg     string    `json:"errmsg,omitempty"`
	Total      int       `json:"total"`
	Count      int       `json:"count"`
	Data       *FansData `json:"data,omitempty"`
	NextOpenid string    `json:"next_openid,omitempty"`
}

type BlackListRsp struct {
	Errcode    int       `json:"errcode,omitempty"`
	Errmsg     string    `json:"errmsg,omitempty"`
	Total      int       `json:"total"`
	Count      int       `json:"count"`
	Data       *FansData `json:"data,omitempty"`
	NextOpenid string    `json:"next_openid,omitempty"`
}

// ============ 基础接口 ============

type ApiDomainIpRsp struct {
	Errcode int      `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	IpList  []string `json:"ip_list,omitempty"`
}

type CallbackIpRsp struct {
	Errcode int      `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	IpList  []string `json:"ip_list,omitempty"`
}

type ApiQuotaRsp struct {
	Errcode int        `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Quota   *QuotaInfo `json:"quota,omitempty"`
}

type QuotaInfo struct {
	DailyLimit int `json:"daily_limit"`
	Used       int `json:"used"`
	Remain     int `json:"remain"`
}

type RidInfoRsp struct {
	Errcode int                    `json:"errcode,omitempty"`
	Errmsg  string                 `json:"errmsg,omitempty"`
	Request map[string]interface{} `json:"request,omitempty"`
}

// ============ 数据统计 ============

type UserSummaryRsp struct {
	Errcode int            `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	List    []*SummaryItem `json:"list,omitempty"`
}

type SummaryItem struct {
	RefDate      string `json:"ref_date"`
	UserSource   int    `json:"user_source,omitempty"`
	NewUser      int    `json:"new_user,omitempty"`
	CancelUser   int    `json:"cancel_user,omitempty"`
	CumulateUser int    `json:"cumulate_user,omitempty"`
}

type UserCumulateRsp struct {
	Errcode int             `json:"errcode,omitempty"`
	Errmsg  string          `json:"errmsg,omitempty"`
	List    []*CumulateItem `json:"list,omitempty"`
}

type CumulateItem struct {
	RefDate      string `json:"ref_date"`
	CumulateUser int    `json:"cumulate_user"`
}

type ArticleSummaryRsp struct {
	Errcode int            `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	List    []*ArticleItem `json:"list,omitempty"`
}

type ArticleItem struct {
	RefDate          string                   `json:"ref_date"`
	Msgid            string                   `json:"msgid,omitempty"`
	Title            string                   `json:"title,omitempty"`
	Details          []map[string]interface{} `json:"details,omitempty"`
	IntPageReadUser  int                      `json:"int_page_read_user,omitempty"`
	IntPageReadCount int                      `json:"int_page_read_count,omitempty"`
	OriPageReadUser  int                      `json:"ori_page_read_user,omitempty"`
	OriPageReadCount int                      `json:"ori_page_read_count,omitempty"`
	ShareUser        int                      `json:"share_user,omitempty"`
	ShareCount       int                      `json:"share_count,omitempty"`
	AddToFavUser     int                      `json:"add_to_fav_user,omitempty"`
	AddToFavCount    int                      `json:"add_to_fav_count,omitempty"`
}

type UpstreamMsgRsp struct {
	Errcode int        `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	List    []*MsgItem `json:"list,omitempty"`
}

type MsgItem struct {
	RefDate  string `json:"ref_date"`
	MsgType  int    `json:"msg_type"`
	MsgUser  int    `json:"msg_user"`
	MsgCount int    `json:"msg_count"`
}

type InterfaceSummaryRsp struct {
	Errcode int              `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	List    []*InterfaceItem `json:"list,omitempty"`
}

type InterfaceItem struct {
	RefDate       string `json:"ref_date"`
	RefHour       int    `json:"ref_hour,omitempty"`
	CallbackCount int    `json:"callback_count"`
	FailCount     int    `json:"fail_count"`
	TotalTimeCost int    `json:"total_time_cost"`
	MaxTimeCost   int    `json:"max_time_cost"`
}

// ============ 草稿箱与发布 ============

type DraftAddRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	MediaId string `json:"media_id,omitempty"`
}

type DraftGetRsp struct {
	Errcode  int                      `json:"errcode,omitempty"`
	Errmsg   string                   `json:"errmsg,omitempty"`
	NewsItem []map[string]interface{} `json:"news_item,omitempty"`
}

type FreepublishSubmitRsp struct {
	Errcode   int    `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	PublishId string `json:"publish_id,omitempty"`
	MsgDataId int64  `json:"msg_data_id,omitempty"`
}

type FreepublishGetRsp struct {
	Errcode       int                    `json:"errcode,omitempty"`
	Errmsg        string                 `json:"errmsg,omitempty"`
	PublishStatus int                    `json:"publish_status,omitempty"`
	ArticleId     string                 `json:"article_id,omitempty"`
	ArticleDetail map[string]interface{} `json:"article_detail,omitempty"`
	FailIdx       []int                  `json:"fail_idx,omitempty"`
}

type FreepublishArticleRsp struct {
	Errcode  int                      `json:"errcode,omitempty"`
	Errmsg   string                   `json:"errmsg,omitempty"`
	NewsItem []map[string]interface{} `json:"news_item,omitempty"`
}

type FreepublishBatchGetRsp struct {
	Errcode    int                      `json:"errcode,omitempty"`
	Errmsg     string                   `json:"errmsg,omitempty"`
	TotalCount int                      `json:"total_count"`
	ItemCount  int                      `json:"item_count"`
	Item       []map[string]interface{} `json:"item,omitempty"`
}

// ============ 评论管理 ============

type CommentListRsp struct {
	Errcode int        `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Total   int        `json:"total"`
	Comment []*Comment `json:"comment,omitempty"`
}

type Comment struct {
	UserCommentId int64         `json:"user_comment_id"`
	Openid        string        `json:"openid"`
	CreateTime    int64         `json:"create_time"`
	Content       string        `json:"content"`
	CommentType   int           `json:"comment_type"`
	Reply         *CommentReply `json:"reply,omitempty"`
}

type CommentReply struct {
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

// ============ AI开放能力 ============

type VoiceTranslateRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}

type QrcodeImgScanRsp struct {
	Errcode     int             `json:"errcode,omitempty"`
	Errmsg      string          `json:"errmsg,omitempty"`
	CodeResults []*QrcodeResult `json:"code_results,omitempty"`
	ImgSize     *ImgSize        `json:"img_size,omitempty"`
}

type QrcodeResult struct {
	TypeName string     `json:"typeName"`
	Data     string     `json:"data"`
	Pos      *QrcodePos `json:"pos,omitempty"`
}

type QrcodePos struct {
	LeftTop     *Position `json:"left_top"`
	RightTop    *Position `json:"right_top"`
	RightBottom *Position `json:"right_bottom"`
	LeftBottom  *Position `json:"left_bottom"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ImgSize struct {
	W int `json:"w"`
	H int `json:"h"`
}

type AiCropImgRsp struct {
	Errcode int           `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Results []*CropResult `json:"results,omitempty"`
	ImgSize *ImgSize      `json:"img_size,omitempty"`
}

type CropResult struct {
	CropLeft   int `json:"crop_left"`
	CropTop    int `json:"crop_top"`
	CropRight  int `json:"crop_right"`
	CropBottom int `json:"crop_bottom"`
}

type OcrIdCardRsp struct {
	Errcode     int    `json:"errcode,omitempty"`
	Errmsg      string `json:"errmsg,omitempty"`
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
	Addr        string `json:"addr,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	ValidDate   string `json:"valid_date,omitempty"`
}

type OcrBankCardRsp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Number  string `json:"number,omitempty"`
}

type OcrDrivingRsp struct {
	Errcode             int    `json:"errcode,omitempty"`
	Errmsg              string `json:"errmsg,omitempty"`
	PlateNum            string `json:"plate_num,omitempty"`
	VehicleType         string `json:"vehicle_type,omitempty"`
	Owner               string `json:"owner,omitempty"`
	Addr                string `json:"addr,omitempty"`
	UseCharacter        string `json:"use_character,omitempty"`
	Model               string `json:"model,omitempty"`
	Vin                 string `json:"vin,omitempty"`
	EngineNum           string `json:"engine_num,omitempty"`
	RegisterDate        string `json:"register_date,omitempty"`
	IssueDate           string `json:"issue_date,omitempty"`
	PlateNumB           string `json:"plate_num_b,omitempty"`
	Record              string `json:"record,omitempty"`
	PassengersNum       string `json:"passengers_num,omitempty"`
	TotalQuality        string `json:"total_quality,omitempty"`
	TotalpRepareQuality string `json:"totalprepare_quality,omitempty"`
}

type OcrDrivingLicenseRsp struct {
	Errcode      int    `json:"errcode,omitempty"`
	Errmsg       string `json:"errmsg,omitempty"`
	IdNum        string `json:"id_num,omitempty"`
	Name         string `json:"name,omitempty"`
	Nationality  string `json:"nationality,omitempty"`
	Sex          string `json:"sex,omitempty"`
	Address      string `json:"address,omitempty"`
	BirthDate    string `json:"birth_date,omitempty"`
	IssueDate    string `json:"issue_date,omitempty"`
	CarClass     string `json:"car_class,omitempty"`
	ValidFrom    string `json:"valid_from,omitempty"`
	ValidTo      string `json:"valid_to,omitempty"`
	OfficialSeal string `json:"official_seal,omitempty"`
}

type OcrBizLicenseRsp struct {
	Errcode             int           `json:"errcode,omitempty"`
	Errmsg              string        `json:"errmsg,omitempty"`
	RegNum              string        `json:"reg_num,omitempty"`
	Serial              string        `json:"serial,omitempty"`
	LegalRepresentative string        `json:"legal_representative,omitempty"`
	EnterpriseName      string        `json:"enterprise_name,omitempty"`
	TypeOfOrganization  string        `json:"type_of_organization,omitempty"`
	Address             string        `json:"address,omitempty"`
	TypeOfEnterprise    string        `json:"type_of_enterprise,omitempty"`
	BusinessScope       string        `json:"business_scope,omitempty"`
	RegisteredCapital   string        `json:"registered_capital,omitempty"`
	PaidInCapital       string        `json:"paid_in_capital,omitempty"`
	ValidPeriod         string        `json:"valid_period,omitempty"`
	RegisteredDate      string        `json:"registered_date,omitempty"`
	CertPosition        *CertPosition `json:"cert_position,omitempty"`
	ImgSize             *ImgSize      `json:"img_size,omitempty"`
}

type CertPosition struct {
	Pos *QrcodePos `json:"pos,omitempty"`
}

type OcrCommonRsp struct {
	Errcode int        `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Items   []*OcrItem `json:"items,omitempty"`
	ImgSize *ImgSize   `json:"img_size,omitempty"`
}

type OcrItem struct {
	Text string     `json:"text"`
	Pos  *QrcodePos `json:"pos,omitempty"`
}
