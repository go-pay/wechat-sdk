package mini

import "github.com/go-pay/xtime"

const (
	Success = 0

	MsgTypeText     = 0 // 文本消息
	MsgTypeImage    = 1 // 图片消息
	MsgTypeLink     = 2 // 图文链接
	MsgTypeMiniPage = 3 // 小程序卡片

	TypingTyping = 0 // 对用户下发"正在输入"状态，
	TypingCancel = 1 // 取消对用户的"正在输入"状态

	HostDefault = "https://api.weixin.qq.com"
)

type MsgType int8

type TypingStatus int8

//type Config struct {
//	Appid       string
//	Secret      string
//	AccessToken string
//	Host        string
//}

type ErrorCode struct {
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type Code2Session struct {
	Openid     string `json:"openid,omitempty"`      // 用户唯一标识
	SessionKey string `json:"session_key,omitempty"` // 会话密钥
	Unionid    string `json:"unionid,omitempty"`     // 用户在开放平台的唯一标识符
	Errcode    int    `json:"errcode,omitempty"`     // 错误码
	Errmsg     string `json:"errmsg,omitempty"`      // 错误信息
}

type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in,omitempty"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	Errcode     int    `json:"errcode,omitempty"`      // 错误码
	Errmsg      string `json:"errmsg,omitempty"`       // 错误信息
}

type PaidUnionId struct {
	Unionid string `json:"unionid,omitempty"` // 用户在开放平台的唯一标识符
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type CheckEncryptedResult struct {
	Vaild      bool   `json:"vaild"`             // 是否是合法的数据
	CreateTime int    `json:"create_time"`       // 加密数据生成的时间戳
	Errcode    int    `json:"errcode,omitempty"` // 错误码
	Errmsg     string `json:"errmsg,omitempty"`  // 错误信息
}

// 微信小程序解密后 用户手机号
type UserPhone struct {
	PhoneNumber     string         `json:"phoneNumber,omitempty"`
	PurePhoneNumber string         `json:"purePhoneNumber,omitempty"`
	CountryCode     string         `json:"countryCode,omitempty"`
	Watermark       *watermarkInfo `json:"watermark,omitempty"`
}

// 微信小程序解密后 用户信息
type UserInfo struct {
	OpenId    string         `json:"openId,omitempty"`
	NickName  string         `json:"nickName,omitempty"`
	Gender    int            `json:"gender,omitempty"`
	City      string         `json:"city,omitempty"`
	Province  string         `json:"province,omitempty"`
	Country   string         `json:"country,omitempty"`
	AvatarUrl string         `json:"avatarUrl,omitempty"`
	UnionId   string         `json:"unionId,omitempty"`
	Watermark *watermarkInfo `json:"watermark,omitempty"`
}

type watermarkInfo struct {
	Appid     string `json:"appid,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

type UploadTempMedia struct {
	MediaId   string     `json:"media_id"`          // 媒体文件上传后，获取标识，3天内有效。
	Type      string     `json:"type"`              // 文件类型
	CreatedAt xtime.Time `json:"created_at"`        // 媒体文件上传时间戳
	Errcode   int        `json:"errcode,omitempty"` // 错误码
	Errmsg    string     `json:"errmsg,omitempty"`  // 错误信息
}

type PhoneNumberRsp struct {
	PhoneInfo *UserPhone `json:"phone_info,omitempty"` // 用户手机号信息
	Errcode   int        `json:"errcode,omitempty"`    // 错误码
	Errmsg    string     `json:"errmsg,omitempty"`     // 错误信息
}

type PluginOpenPid struct {
	Openpid string `json:"openpid,omitempty"` // 插件用户的唯一标识
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type UserEncryptKey struct {
	Errcode     int        `json:"errcode"`
	Errmsg      string     `json:"errmsg"`
	KeyInfoList []*KeyInfo `json:"key_info_list"`
}

type KeyInfo struct {
	EncryptKey string `json:"encrypt_key"`
	Version    int    `json:"version"`
	ExpireIn   int    `json:"expire_in"`
	Iv         string `json:"iv"`
	CreateTime int    `json:"create_time"`
}

type MediaAssetSingleFileUploadRsp struct {
	MediaId int    `json:"media_id"` // 媒体文件id。
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MediaAssetPullUploadRsp struct {
	TaskId  int    `json:"task_id"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MediaAssetGetTaskRsp struct {
	TaskInfo *PullUploadTaskInfo `json:"task_info"`
	Errcode  int                 `json:"errcode"`
	Errmsg   string              `json:"errmsg"`
}

type PullUploadTaskInfo struct {
	Id         int    `json:"id"`
	TaskType   int    `json:"task_type"`
	Status     int    `json:"status"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	CreateTime int    `json:"create_time"`
	FinishTime int    `json:"finish_time"`
	MediaId    int    `json:"media_id"`
}

type MediaAssetApplyUploadRsp struct {
	UploadId string `json:"upload_id"`
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
}

type MediaAssetUploadPartRsp struct {
	Etag    string `json:"etag"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MediaAssetCommitUploadRsp struct {
	MediaId int    `json:"media_id"` // 媒体文件id。
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MediaAssetListMediaRsp struct {
	Errcode       int          `json:"errcode"`
	Errmsg        string       `json:"errmsg"`
	MediaInfoList []*MediaInfo `json:"media_info_list"`
}

type MediaInfo struct {
	MediaId     int          `json:"media_id"`
	CreateTime  int          `json:"create_time"`
	ExpireTime  int          `json:"expire_time"`
	DramaId     int          `json:"drama_id"`
	FileSize    string       `json:"file_size"`
	Duration    int          `json:"duration"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CoverUrl    string       `json:"cover_url"`
	OriginalUrl string       `json:"original_url"`
	Mp4Url      string       `json:"mp4_url"`
	HlsUrl      string       `json:"hls_url"`
	AuditDetail *AuditDetail `json:"audit_detail"`
}

type AuditDetail struct {
	Status                 int      `json:"status"`
	CreateTime             int      `json:"create_time"`
	AuditTime              int      `json:"audit_time"`
	Reason                 string   `json:"reason"`
	EvidenceMaterialIdList []string `json:"evidence_material_id_list"`
}

type MediaAssetGetMediaRsp struct {
	Errcode   int        `json:"errcode"`
	Errmsg    string     `json:"errmsg"`
	MediaInfo *MediaInfo `json:"media_info"`
}

type MediaAssetGetMediaLinkRsp struct {
	Errcode   int                `json:"errcode"`
	Errmsg    string             `json:"errmsg"`
	MediaInfo *MediaPlaybackInfo `json:"media_info"`
}

type MediaPlaybackInfo struct {
	MediaId     int    `json:"media_id"`
	Duration    int    `json:"duration"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CoverUrl    string `json:"cover_url"`
	Mp4Url      string `json:"mp4_url"`
	HlsUrl      string `json:"hls_url"`
}

type MediaAssetDeleteMediaRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ==================== 订阅消息相关 ====================

// SubscribeMessageSendRsp 发送订阅消息响应
type SubscribeMessageSendRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// TemplateListRsp 获取模板列表响应
type TemplateListRsp struct {
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
	Data    []*TemplateInfo `json:"data"`
}

// TemplateInfo 模板信息
type TemplateInfo struct {
	PriTmplId string `json:"priTmplId"` // 模板id
	Title     string `json:"title"`     // 模板标题
	Content   string `json:"content"`   // 模板内容
	Example   string `json:"example"`   // 模板示例
	Type      int    `json:"type"`      // 模板类型，2为一次性订阅，3为长期订阅
}

// PubTemplateTitleListRsp 获取模板标题列表响应
type PubTemplateTitleListRsp struct {
	Errcode int                     `json:"errcode"`
	Errmsg  string                  `json:"errmsg"`
	Count   int                     `json:"count"`
	Data    []*PubTemplateTitleInfo `json:"data"`
}

// PubTemplateTitleInfo 模板标题信息
type PubTemplateTitleInfo struct {
	Tid        int    `json:"tid"`        // 模板标题id
	Title      string `json:"title"`      // 模板标题
	Type       int    `json:"type"`       // 模板类型，2为一次性订阅，3为长期订阅
	CategoryId string `json:"categoryId"` // 模板所属类目id
}

// AddTemplateRsp 添加模板响应
type AddTemplateRsp struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	PriTmplId string `json:"priTmplId"` // 添加至账号下的模板id
}

// DeleteTemplateRsp 删除模板响应
type DeleteTemplateRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CategoryListRsp 获取类目列表响应
type CategoryListRsp struct {
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
	Data    []*CategoryInfo `json:"data"`
}

// CategoryInfo 类目信息
type CategoryInfo struct {
	Id   int    `json:"id"`   // 类目id
	Name string `json:"name"` // 类目名称
}

// ==================== 小程序码相关 ====================

// QRCodeRsp 小程序码/二维码响应（返回图片 Buffer）
type QRCodeRsp struct {
	Buffer  []byte // 图片 Buffer
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

// LineColor 小程序码线条颜色
type LineColor struct {
	R string `json:"r"` // red，0-255
	G string `json:"g"` // green，0-255
	B string `json:"b"` // blue，0-255
}

// ==================== URL Scheme/Link 相关 ====================

// GenerateSchemeRsp 生成 URL Scheme 响应
type GenerateSchemeRsp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Openlink string `json:"openlink"` // 生成的小程序 scheme 码
}

// QuerySchemeRsp 查询 URL Scheme 响应
type QuerySchemeRsp struct {
	Errcode    int              `json:"errcode"`
	Errmsg     string           `json:"errmsg"`
	SchemeInfo *SchemeInfo      `json:"scheme_info"`
	QuotaInfo  *SchemeQuotaInfo `json:"quota_info"`
}

// SchemeInfo Scheme 信息
type SchemeInfo struct {
	Appid      string `json:"appid"`       // 小程序 appid
	Path       string `json:"path"`        // 小程序页面路径
	Query      string `json:"query"`       // 小程序页面query
	CreateTime int64  `json:"create_time"` // 创建时间
	ExpireTime int64  `json:"expire_time"` // 到期失效时间
	EnvVersion string `json:"env_version"` // 要打开的小程序版本
}

// SchemeQuotaInfo Scheme 配额信息
type SchemeQuotaInfo struct {
	RemainVisitQuota int64 `json:"remain_visit_quota"` // URL Scheme（小程序链接）剩余访问次数
}

// GenerateUrlLinkRsp 生成 URL Link 响应
type GenerateUrlLinkRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UrlLink string `json:"url_link"` // 生成的小程序 URL Link
}

// QueryUrlLinkRsp 查询 URL Link 响应
type QueryUrlLinkRsp struct {
	Errcode      int           `json:"errcode"`
	Errmsg       string        `json:"errmsg"`
	UrlLinkInfo  *UrlLinkInfo  `json:"url_link_info"`
	UrlLinkQuota *UrlLinkQuota `json:"url_link_quota"`
	VisitOpenid  []string      `json:"visit_openid"` // 访问用户的 openid 列表
}

// UrlLinkInfo URL Link 信息
type UrlLinkInfo struct {
	Appid      string `json:"appid"`       // 小程序 appid
	Path       string `json:"path"`        // 小程序页面路径
	Query      string `json:"query"`       // 小程序页面query
	CreateTime int64  `json:"create_time"` // 创建时间
	ExpireTime int64  `json:"expire_time"` // 到期失效时间
	EnvVersion string `json:"env_version"` // 要打开的小程序版本
}

// UrlLinkQuota URL Link 配额信息
type UrlLinkQuota struct {
	RemainVisitQuota int64 `json:"remain_visit_quota"` // URL Link 剩余访问次数
}

// GenerateShortLinkRsp 生成 Short Link 响应
type GenerateShortLinkRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Link    string `json:"link"` // 生成的小程序 Short Link
}

// ==================== 内容安全相关 ====================

// MsgSecCheckRsp 文本内容安全检测响应
type MsgSecCheckRsp struct {
	Errcode int               `json:"errcode"`
	Errmsg  string            `json:"errmsg"`
	Result  *SecCheckResult   `json:"result"`
	Detail  []*SecCheckDetail `json:"detail"`
	TraceId string            `json:"trace_id"` // 唯一请求标识，标记单次请求
}

// SecCheckResult 检测结果
type SecCheckResult struct {
	Suggest string `json:"suggest"` // 建议，有risky、pass、review三种值
	Label   int    `json:"label"`   // 命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
}

// SecCheckDetail 详细检测结果
type SecCheckDetail struct {
	Strategy string `json:"strategy"` // 策略类型
	Errcode  int    `json:"errcode"`  // 错误码，仅当该值为0时，该项结果有效
	Suggest  string `json:"suggest"`  // 建议，有risky、pass、review三种值
	Label    int    `json:"label"`    // 命中标签枚举值
	Prob     int    `json:"prob"`     // 0-100，代表置信度，越高代表越有可能属于当前返回的标签（label）
	Level    int    `json:"level"`    // 命中的自定义关键词的等级，可能返回1、2、3
	Keyword  string `json:"keyword"`  // 命中的自定义关键词
}

// ImgSecCheckRsp 图片内容安全检测响应
type ImgSecCheckRsp struct {
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
	Result  *SecCheckResult `json:"result"`
	TraceId string          `json:"trace_id"` // 唯一请求标识，标记单次请求
}

// MediaCheckAsyncRsp 音视频内容安全异步检测响应
type MediaCheckAsyncRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	TraceId string `json:"trace_id"` // 唯一请求标识，标记单次请求
}

// ==================== 数据分析相关 ====================

// DailySummaryRsp 获取用户访问小程序数据概况响应
type DailySummaryRsp struct {
	Errcode int                 `json:"errcode"`
	Errmsg  string              `json:"errmsg"`
	List    []*DailySummaryItem `json:"list"`
}

// DailySummaryItem 日概况数据项
type DailySummaryItem struct {
	RefDate         string  `json:"ref_date"`          // 日期，格式为 yyyyMMdd
	VisitTotal      int     `json:"visit_total"`       // 累计用户数
	SharePv         int     `json:"share_pv"`          // 转发次数
	ShareUv         int     `json:"share_uv"`          // 转发人数
	VisitUv         int     `json:"visit_uv"`          // 访问人数
	VisitUvNew      int     `json:"visit_uv_new"`      // 新用户数
	StayTimeUv      float64 `json:"stay_time_uv"`      // 人均停留时长 (浮点型，单位：秒)
	StayTimeSession float64 `json:"stay_time_session"` // 次均停留时长 (浮点型，单位：秒)
	VisitDepth      float64 `json:"visit_depth"`       // 平均访问深度 (浮点型)
}

// VisitTrendRsp 获取用户访问小程序数据趋势响应
type VisitTrendRsp struct {
	Errcode int               `json:"errcode"`
	Errmsg  string            `json:"errmsg"`
	List    []*VisitTrendItem `json:"list"`
}

// VisitTrendItem 访问趋势数据项
type VisitTrendItem struct {
	RefDate         string  `json:"ref_date"`          // 日期，格式为 yyyyMMdd
	SessionCnt      int     `json:"session_cnt"`       // 打开次数
	VisitPv         int     `json:"visit_pv"`          // 访问次数
	VisitUv         int     `json:"visit_uv"`          // 访问人数
	VisitUvNew      int     `json:"visit_uv_new"`      // 新用户数
	StayTimeUv      float64 `json:"stay_time_uv"`      // 人均停留时长 (浮点型，单位：秒)
	StayTimeSession float64 `json:"stay_time_session"` // 次均停留时长 (浮点型，单位：秒)
	VisitDepth      float64 `json:"visit_depth"`       // 平均访问深度 (浮点型)
}

// RetainInfoRsp 获取用户访问小程序留存响应
type RetainInfoRsp struct {
	Errcode    int           `json:"errcode"`
	Errmsg     string        `json:"errmsg"`
	RefDate    string        `json:"ref_date"`     // 日期，格式为 yyyyMMdd
	VisitUvNew []*RetainItem `json:"visit_uv_new"` // 新增用户留存
	VisitUv    []*RetainItem `json:"visit_uv"`     // 活跃用户留存
}

// RetainItem 留存数据项
type RetainItem struct {
	Key   int     `json:"key"`   // 标识，0开始，0表示当天，1表示1天后，依此类推，key取值分别是：0,1,2,3,4,5,6,7,14,30
	Value float64 `json:"value"` // 留存率
}

// VisitPageRsp 获取访问页面数据响应
type VisitPageRsp struct {
	Errcode int              `json:"errcode"`
	Errmsg  string           `json:"errmsg"`
	RefDate string           `json:"ref_date"` // 日期，格式为 yyyyMMdd
	List    []*VisitPageItem `json:"list"`
}

// VisitPageItem 访问页面数据项
type VisitPageItem struct {
	PagePath       string  `json:"page_path"`        // 页面路径
	PageVisitPv    int     `json:"page_visit_pv"`    // 访问次数
	PageVisitUv    int     `json:"page_visit_uv"`    // 访问人数
	PageStaytimePv float64 `json:"page_staytime_pv"` // 次均停留时长
	EntrypagePv    int     `json:"entrypage_pv"`     // 进入页次数
	ExitpagePv     int     `json:"exitpage_pv"`      // 退出页次数
	PageSharePv    int     `json:"page_share_pv"`    // 转发次数
	PageShareUv    int     `json:"page_share_uv"`    // 转发人数
}

// UserPortraitRsp 获取小程序用户画像分布响应
type UserPortraitRsp struct {
	Errcode    int             `json:"errcode"`
	Errmsg     string          `json:"errmsg"`
	RefDate    string          `json:"ref_date"`            // 日期，格式为 yyyyMMdd
	VisitUv    int             `json:"visit_uv"`            // 活跃用户数
	VisitUvNew int             `json:"visit_uv_new"`        // 新用户数
	Province   []*PortraitItem `json:"province,omitempty"`  // 省份分布
	City       []*PortraitItem `json:"city,omitempty"`      // 城市分布
	Genders    []*PortraitItem `json:"genders,omitempty"`   // 性别分布
	Platforms  []*PortraitItem `json:"platforms,omitempty"` // 终端类型分布
	Devices    []*PortraitItem `json:"devices,omitempty"`   // 机型分布
	Ages       []*PortraitItem `json:"ages,omitempty"`      // 年龄分布
}

// PortraitItem 画像分布数据项
type PortraitItem struct {
	Id    int     `json:"id,omitempty"`    // 属性值id
	Name  string  `json:"name,omitempty"`  // 属性值名称
	Value float64 `json:"value,omitempty"` // 该属性值对应的用户数量占比
}

// PerformanceDataRsp 获取小程序性能数据响应
type PerformanceDataRsp struct {
	Errcode int              `json:"errcode"`
	Errmsg  string           `json:"errmsg"`
	Body    *PerformanceBody `json:"body"`
}

// PerformanceBody 性能数据主体
type PerformanceBody struct {
	Tables []*PerformanceTable `json:"tables"`
}

// PerformanceTable 性能数据表
type PerformanceTable struct {
	RefDate   string              `json:"refDate"` // 日期，格式为 yyyyMMdd
	Cost      int                 `json:"cost"`    // 耗时，单位ms
	Ratio     float64             `json:"ratio"`   // 占比
	FieldList []*PerformanceField `json:"fieldList,omitempty"`
}

// PerformanceField 性能数据字段
type PerformanceField struct {
	FieldName  string `json:"fieldName"`  // 字段名
	FieldValue string `json:"fieldValue"` // 字段值
}

// VisitDistributionRsp 获取用户小程序访问分布数据响应
type VisitDistributionRsp struct {
	Errcode int                      `json:"errcode"`
	Errmsg  string                   `json:"errmsg"`
	RefDate string                   `json:"ref_date"` // 日期，格式为 yyyyMMdd
	List    []*VisitDistributionItem `json:"list,omitempty"`
}

// VisitDistributionItem 访问分布数据项
type VisitDistributionItem struct {
	Index    int                   `json:"index"` // 分布类型
	ItemList []*DistributionDetail `json:"item_list,omitempty"`
}

// DistributionDetail 分布详情
type DistributionDetail struct {
	Key   int     `json:"key,omitempty"`   // 场景id 或 访问深度 或 停留时长
	Value float64 `json:"value,omitempty"` // 该场景id访问pv占比 或 该访问深度访问人数占比 或 该停留时长访问人数占比
}

// ==================== 图像处理相关 ====================

// AiCropRsp 图片智能裁剪响应
type AiCropRsp struct {
	Errcode int           `json:"errcode"`
	Errmsg  string        `json:"errmsg"`
	Results []*CropResult `json:"results"`
}

// CropResult 裁剪结果
type CropResult struct {
	CropLeft   int `json:"crop_left"`   // 裁剪区域左上角横坐标
	CropTop    int `json:"crop_top"`    // 裁剪区域左上角纵坐标
	CropRight  int `json:"crop_right"`  // 裁剪区域右下角横坐标
	CropBottom int `json:"crop_bottom"` // 裁剪区域右下角纵坐标
}

// ScanQRCodeRsp 条码/二维码识别响应
type ScanQRCodeRsp struct {
	Errcode     int           `json:"errcode"`
	Errmsg      string        `json:"errmsg"`
	CodeResults []*CodeResult `json:"code_results"`
	ImgSize     *ImageSize    `json:"img_size"`
}

// CodeResult 识别结果
type CodeResult struct {
	TypeName string   `json:"type_name"` // 类型名称，如 QR_CODE、EAN_13 等
	Data     string   `json:"data"`      // 二维码/条码内容
	Pos      *CodePos `json:"pos"`       // 位置信息
}

// CodePos 位置信息
type CodePos struct {
	LeftTop     *Point `json:"left_top"`     // 左上角坐标
	RightTop    *Point `json:"right_top"`    // 右上角坐标
	RightBottom *Point `json:"right_bottom"` // 右下角坐标
	LeftBottom  *Point `json:"left_bottom"`  // 左下角坐标
}

// Point 坐标点
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// ImageSize 图片尺寸
type ImageSize struct {
	W int `json:"w"` // 宽度
	H int `json:"h"` // 高度
}

// SuperResolutionRsp 图片高清化响应（返回图片 Buffer）
type SuperResolutionRsp struct {
	Buffer  []byte // 图片 Buffer
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

// OcrIdCardRsp 身份证OCR识别响应
type OcrIdCardRsp struct {
	Errcode int     `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
	Type    string  `json:"type"`    // 正面或背面，Front / Back
	IdCard  *IdCard `json:"id_card"` // 身份证信息
}

// IdCard 身份证信息
type IdCard struct {
	Name        string `json:"name,omitempty"`        // 姓名
	Id          string `json:"id,omitempty"`          // 身份证号
	Addr        string `json:"addr,omitempty"`        // 地址
	Gender      string `json:"gender,omitempty"`      // 性别
	Nationality string `json:"nationality,omitempty"` // 民族
	ValidDate   string `json:"valid_date,omitempty"`  // 有效期
}

// OcrBankCardRsp 银行卡OCR识别响应
type OcrBankCardRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Number  string `json:"number"` // 银行卡号
}

// OcrDrivingRsp 驾驶证OCR识别响应
type OcrDrivingRsp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	IdNum       string `json:"id_num"`      // 证号
	Name        string `json:"name"`        // 姓名
	Sex         string `json:"sex"`         // 性别
	Nationality string `json:"nationality"` // 国籍
	Address     string `json:"address"`     // 住址
	BirthDate   string `json:"birth_date"`  // 出生日期
	IssueDate   string `json:"issue_date"`  // 初次领证日期
	CarClass    string `json:"car_class"`   // 准驾车型
	ValidFrom   string `json:"valid_from"`  // 有效期限起始日
	ValidTo     string `json:"valid_to"`    // 有效期限终止日
}

// OcrVehicleLicenseRsp 行驶证OCR识别响应
type OcrVehicleLicenseRsp struct {
	Errcode       int    `json:"errcode"`
	Errmsg        string `json:"errmsg"`
	VehicleType   string `json:"vehicle_type"`   // 车辆类型
	Owner         string `json:"owner"`          // 所有人
	Addr          string `json:"addr"`           // 住址
	UseCharacter  string `json:"use_character"`  // 使用性质
	Model         string `json:"model"`          // 品牌型号
	Vin           string `json:"vin"`            // 车辆识别代号
	EngineNum     string `json:"engine_num"`     // 发动机号码
	RegisterDate  string `json:"register_date"`  // 注册日期
	IssueDate     string `json:"issue_date"`     // 发证日期
	PlateNumB     string `json:"plate_num_b"`    // 号牌号码
	Record        string `json:"record"`         // 档案编号
	PassengersNum string `json:"passengers_num"` // 核定载人数
	TotalQuality  string `json:"total_quality"`  // 总质量
	TotalMass     string `json:"total_mass"`     // 整备质量
}

// OcrBusinessLicenseRsp 营业执照OCR识别响应
type OcrBusinessLicenseRsp struct {
	Errcode             int           `json:"errcode"`
	Errmsg              string        `json:"errmsg"`
	RegNum              string        `json:"reg_num"`              // 注册号
	Serial              string        `json:"serial"`               // 编号
	LegalRepresentative string        `json:"legal_representative"` // 法定代表人姓名
	EnterpriseName      string        `json:"enterprise_name"`      // 企业名称
	TypeOfEnterprise    string        `json:"type_of_enterprise"`   // 组成形式
	Address             string        `json:"address"`              // 经营场所/企业住所
	TypeOfOrganization  string        `json:"type_of_organization"` // 公司类型
	BusinessScope       string        `json:"business_scope"`       // 经营范围
	RegisteredCapital   string        `json:"registered_capital"`   // 注册资本
	PaidInCapital       string        `json:"paid_in_capital"`      // 实收资本
	ValidPeriod         string        `json:"valid_period"`         // 营业期限
	RegisteredDate      string        `json:"registered_date"`      // 注册日期/成立日期
	CertPosition        *CertPosition `json:"cert_position"`        // 营业执照位置
	ImgSize             *ImageSize    `json:"img_size"`             // 图片大小
}

// CertPosition 证件位置
type CertPosition struct {
	Pos *CodePos `json:"pos"` // 位置信息
}

// OcrCommonRsp 通用印刷体OCR识别响应
type OcrCommonRsp struct {
	Errcode int        `json:"errcode"`
	Errmsg  string     `json:"errmsg"`
	Items   []*OcrItem `json:"items"`
}

// OcrItem OCR识别项
type OcrItem struct {
	Text string   `json:"text"` // 识别的文本
	Pos  *CodePos `json:"pos"`  // 位置信息
}

// OcrPlateNumberRsp 车牌OCR识别响应
type OcrPlateNumberRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Number  string `json:"number"` // 车牌号码
}

// ==================== 即时配送相关 ====================

// DeliveryGetAllAccountRsp 拉取已绑定账号响应
type DeliveryGetAllAccountRsp struct {
	Errcode  int             `json:"errcode"`
	Errmsg   string          `json:"errmsg"`
	ShopList []*DeliveryShop `json:"shop_list"`
}

// DeliveryShop 配送商家信息
type DeliveryShop struct {
	DeliveryId  string `json:"delivery_id"`  // 配送公司ID
	ShopId      string `json:"shopid"`       // 商家ID
	AuditResult int    `json:"audit_result"` // 审核状态
}

// DeliveryPreAddOrderRsp 配送单预下单响应
type DeliveryPreAddOrderRsp struct {
	Errcode          int    `json:"errcode"`
	Errmsg           string `json:"errmsg"`
	Fee              int    `json:"fee"`               // 实际运费(单位：元)，运费减去优惠券费用
	Deliverfee       int    `json:"deliverfee"`        // 运费(单位：元)
	Couponfee        int    `json:"couponfee"`         // 优惠券费用(单位：元)
	Tips             int    `json:"tips"`              // 小费(单位：元)
	Insurancefee     int    `json:"insurancefee"`      // 保价费(单位：元)
	Distance         int    `json:"distance"`          // 配送距离(单位：米)
	DispatchDuration int    `json:"dispatch_duration"` // 预计骑手接单时间(单位：秒)
	DeliveryToken    string `json:"delivery_token"`    // 配送token
}

// DeliveryAddOrderRsp 配送单下单响应
type DeliveryAddOrderRsp struct {
	Errcode          int    `json:"errcode"`
	Errmsg           string `json:"errmsg"`
	Fee              int    `json:"fee"`               // 实际运费(单位：元)
	Deliverfee       int    `json:"deliverfee"`        // 运费(单位：元)
	Couponfee        int    `json:"couponfee"`         // 优惠券费用(单位：元)
	Tips             int    `json:"tips"`              // 小费(单位：元)
	Insurancefee     int    `json:"insurancefee"`      // 保价费(单位：元)
	Distance         int    `json:"distance"`          // 配送距离(单位：米)
	WaybillId        string `json:"waybill_id"`        // 配送单号
	OrderStatus      int    `json:"order_status"`      // 配送状态
	FinishCode       int    `json:"finish_code"`       // 收货码
	PickupCode       int    `json:"pickup_code"`       // 取货码
	DispatchDuration int    `json:"dispatch_duration"` // 预计骑手接单时间(单位：秒)
}

// DeliveryAddTipsRsp 配送单增加小费响应
type DeliveryAddTipsRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// DeliveryCancelOrderRsp 配送单取消响应
type DeliveryCancelOrderRsp struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Deduct_fee int    `json:"deduct_fee"` // 扣除的违约金(单位：元)
	Desc       string `json:"desc"`       // 说明
}

// DeliveryGetOrderRsp 配送单查询响应
type DeliveryGetOrderRsp struct {
	Errcode     int     `json:"errcode"`
	Errmsg      string  `json:"errmsg"`
	OrderStatus int     `json:"order_status"` // 配送状态
	WaybillId   string  `json:"waybill_id"`   // 配送单号
	RiderName   string  `json:"rider_name"`   // 骑手姓名
	RiderPhone  string  `json:"rider_phone"`  // 骑手电话
	RiderLng    float64 `json:"rider_lng"`    // 骑手位置经度
	RiderLat    float64 `json:"rider_lat"`    // 骑手位置纬度
}

// DeliveryMockUpdateOrderRsp 模拟配送公司更新配送单状态响应
type DeliveryMockUpdateOrderRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// DeliveryAbnormalConfirmRsp 异常件退回商家商圈接口响应
type DeliveryAbnormalConfirmRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ==================== 物流助手相关 ====================

// ExpressAddOrderRsp 生成运单响应
type ExpressAddOrderRsp struct {
	Errcode            int    `json:"errcode"`
	Errmsg             string `json:"errmsg"`
	OrderId            string `json:"order_id"`            // 订单ID
	WaybillId          string `json:"waybill_id"`          // 运单号
	WaybillData        string `json:"waybill_data"`        // 面单信息
	DeliveryResultcode int    `json:"delivery_resultcode"` // 配送侧错误码
	DeliveryResultmsg  string `json:"delivery_resultmsg"`  // 配送侧错误信息
}

// ExpressCancelOrderRsp 取消运单响应
type ExpressCancelOrderRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ExpressGetAllAccountRsp 获取所有绑定的物流账号响应
type ExpressGetAllAccountRsp struct {
	Errcode int               `json:"errcode"`
	Errmsg  string            `json:"errmsg"`
	Count   int               `json:"count"`
	List    []*ExpressAccount `json:"list"`
}

// ExpressAccount 物流账号信息
type ExpressAccount struct {
	BizId           string `json:"biz_id"`            // 快递公司客户编码
	DeliveryId      string `json:"delivery_id"`       // 快递公司ID
	CreateTime      int64  `json:"create_time"`       // 账号绑定时间
	UpdateTime      int64  `json:"update_time"`       // 账号更新时间
	StatusCode      int    `json:"status_code"`       // 账号状态
	Alias           string `json:"alias"`             // 账号别名
	RemarkWrongMsg  string `json:"remark_wrong_msg"`  // 账号状态说明
	QuotaNum        int    `json:"quota_num"`         // 电子面单余额
	QuotaUpdateTime int64  `json:"quota_update_time"` // 电子面单余额更新时间
}

// ExpressGetQuotaRsp 获取电子面单余额响应
type ExpressGetQuotaRsp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	QuotaNum int    `json:"quota_num"` // 电子面单余额
}

// ExpressGetPathRsp 查询运单轨迹响应
type ExpressGetPathRsp struct {
	Errcode      int                `json:"errcode"`
	Errmsg       string             `json:"errmsg"`
	OrderId      string             `json:"order_id"`       // 订单ID
	WaybillId    string             `json:"waybill_id"`     // 运单号
	DeliveryId   string             `json:"delivery_id"`    // 快递公司ID
	PathItemNum  int                `json:"path_item_num"`  // 轨迹节点数量
	PathItemList []*ExpressPathItem `json:"path_item_list"` // 轨迹节点列表
}

// ExpressPathItem 轨迹节点
type ExpressPathItem struct {
	ActionTime int    `json:"action_time"` // 轨迹节点时间
	ActionType int    `json:"action_type"` // 轨迹节点类型
	ActionMsg  string `json:"action_msg"`  // 轨迹节点详情
}

// ExpressGetAllDeliveryRsp 获取支持的快递公司列表响应
type ExpressGetAllDeliveryRsp struct {
	Errcode int                `json:"errcode"`
	Errmsg  string             `json:"errmsg"`
	Count   int                `json:"count"`
	Data    []*ExpressDelivery `json:"data"`
}

// ExpressDelivery 快递公司信息
type ExpressDelivery struct {
	DeliveryId   string `json:"delivery_id"`   // 快递公司ID
	DeliveryName string `json:"delivery_name"` // 快递公司名称
}

// ExpressGetPrinterRsp 获取打印员响应
type ExpressGetPrinterRsp struct {
	Errcode   int              `json:"errcode"`
	Errmsg    string           `json:"errmsg"`
	Count     int              `json:"count"`
	Openid    []*PrinterOpenid `json:"openid"`
	TagidList string           `json:"tagid_list"` // 用户标签ID列表
}

// PrinterOpenid 打印员openid
type PrinterOpenid struct {
	Openid string `json:"openid"`
}

// ExpressUpdatePrinterRsp 配置面单打印员响应
type ExpressUpdatePrinterRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ExpressGetContactRsp 获取面单联系人信息响应
type ExpressGetContactRsp struct {
	Errcode      int             `json:"errcode"`
	Errmsg       string          `json:"errmsg"`
	WaybillId    string          `json:"waybill_id"` // 运单号
	SenderInfo   *ExpressContact `json:"sender"`     // 寄件人信息
	ReceiverInfo *ExpressContact `json:"receiver"`   // 收件人信息
}

// ExpressContact 联系人信息
type ExpressContact struct {
	Name     string `json:"name"`     // 姓名
	Mobile   string `json:"mobile"`   // 电话
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
	Area     string `json:"area"`     // 区县
	Address  string `json:"address"`  // 详细地址
}

// ==================== 搜索相关 ====================

// SearchSubmitPagesRsp 提交小程序页面响应
type SearchSubmitPagesRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SearchDeletePageRsp 删除已提交的小程序页面响应
type SearchDeletePageRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ==================== 运维中心相关 ====================

// GetDomainInfoRsp 获取域名配置响应
type GetDomainInfoRsp struct {
	Errcode         int      `json:"errcode"`
	Errmsg          string   `json:"errmsg"`
	Requestdomain   []string `json:"requestdomain"`   // request合法域名
	Wsrequestdomain []string `json:"wsrequestdomain"` // socket合法域名
	Uploaddomain    []string `json:"uploaddomain"`    // uploadFile合法域名
	Downloaddomain  []string `json:"downloaddomain"`  // downloadFile合法域名
	Udpdomain       []string `json:"udpdomain"`       // udp合法域名
	Tcpdomain       []string `json:"tcpdomain"`       // tcp合法域名
}

// ModifyDomainRsp 修改服务器域名响应
type ModifyDomainRsp struct {
	Errcode         int      `json:"errcode"`
	Errmsg          string   `json:"errmsg"`
	Requestdomain   []string `json:"requestdomain"`   // request合法域名
	Wsrequestdomain []string `json:"wsrequestdomain"` // socket合法域名
	Uploaddomain    []string `json:"uploaddomain"`    // uploadFile合法域名
	Downloaddomain  []string `json:"downloaddomain"`  // downloadFile合法域名
	Udpdomain       []string `json:"udpdomain"`       // udp合法域名
	Tcpdomain       []string `json:"tcpdomain"`       // tcp合法域名
}

// SetWebviewDomainRsp 设置业务域名响应
type SetWebviewDomainRsp struct {
	Errcode       int      `json:"errcode"`
	Errmsg        string   `json:"errmsg"`
	Webviewdomain []string `json:"webviewdomain"` // 业务域名
}

// GetQrcodeJumppublishRsp 获取小程序码扫码打开的页面响应
type GetQrcodeJumppublishRsp struct {
	Errcode    int               `json:"errcode"`
	Errmsg     string            `json:"errmsg"`
	PrefixList []*QrcodeJumpRule `json:"prefix_list"`
}

// QrcodeJumpRule 二维码跳转规则
type QrcodeJumpRule struct {
	Prefix        string   `json:"prefix"`          // 二维码规则
	PermitSubRule int      `json:"permit_sub_rule"` // 是否独占符合二维码前缀匹配规则的所有子规则
	Path          string   `json:"path"`            // 小程序功能页面
	OpenVersion   int      `json:"open_version"`    // 测试范围
	DebugUrl      []string `json:"debug_url"`       // 测试链接
	State         int      `json:"state"`           // 发布标志位
}

// SetQrcodeJumpRsp 设置小程序码扫码打开的页面响应
type SetQrcodeJumpRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ==================== 小程序管理相关 ====================

// GetAccountBasicInfoRsp 获取小程序基本信息响应
type GetAccountBasicInfoRsp struct {
	Errcode        int            `json:"errcode"`
	Errmsg         string         `json:"errmsg"`
	Appid          string         `json:"appid"`           // 小程序appid
	AccountType    int            `json:"account_type"`    // 帐号类型
	PrincipalType  int            `json:"principal_type"`  // 主体类型
	PrincipalName  string         `json:"principal_name"`  // 主体名称
	Credential     string         `json:"credential"`      // 主体标识
	RealnameStatus int            `json:"realname_status"` // 实名验证状态
	WxVerifyInfo   *WxVerifyInfo  `json:"wx_verify_info"`  // 微信认证信息
	SignatureInfo  *SignatureInfo `json:"signature_info"`  // 功能介绍信息
	HeadImageInfo  *HeadImageInfo `json:"head_image_info"` // 头像信息
	NicknameInfo   *NicknameInfo  `json:"nickname_info"`   // 名称信息
}

// WxVerifyInfo 微信认证信息
type WxVerifyInfo struct {
	QualificationVerify   bool  `json:"qualification_verify"`     // 是否资质认证
	NamingVerify          bool  `json:"naming_verify"`            // 是否名称认证
	AnnualReview          bool  `json:"annual_review"`            // 是否年审
	AnnualReviewBeginTime int64 `json:"annual_review_begin_time"` // 年审开始时间
	AnnualReviewEndTime   int64 `json:"annual_review_end_time"`   // 年审结束时间
}

// SignatureInfo 功能介绍信息
type SignatureInfo struct {
	Signature       string `json:"signature"`         // 功能介绍
	ModifyUsedCount int    `json:"modify_used_count"` // 功能介绍已使用修改次数
	ModifyQuota     int    `json:"modify_quota"`      // 功能介绍修改次数总额度
}

// HeadImageInfo 头像信息
type HeadImageInfo struct {
	HeadImageUrl    string `json:"head_image_url"`    // 头像url
	ModifyUsedCount int    `json:"modify_used_count"` // 头像已使用修改次数
	ModifyQuota     int    `json:"modify_quota"`      // 头像修改次数总额度
}

// NicknameInfo 名称信息
type NicknameInfo struct {
	Nickname        string `json:"nickname"`          // 小程序名称
	ModifyUsedCount int    `json:"modify_used_count"` // 名称已使用修改次数
	ModifyQuota     int    `json:"modify_quota"`      // 名称修改次数总额度
}

// GetPageRsp 获取已上传的代码的页面列表响应
type GetPageRsp struct {
	Errcode  int      `json:"errcode"`
	Errmsg   string   `json:"errmsg"`
	PageList []string `json:"page_list"` // 页面配置列表
}

// GetCategoryRsp 获取授权小程序帐号的可选类目响应
type GetCategoryRsp struct {
	Errcode      int             `json:"errcode"`
	Errmsg       string          `json:"errmsg"`
	CategoryList []*CategoryItem `json:"category_list"`
}

// CategoryItem 类目信息
type CategoryItem struct {
	FirstClass  string `json:"first_class"`  // 一级类目名称
	SecondClass string `json:"second_class"` // 二级类目名称
	ThirdClass  string `json:"third_class"`  // 三级类目名称
	FirstId     int    `json:"first_id"`     // 一级类目ID
	SecondId    int    `json:"second_id"`    // 二级类目ID
	ThirdId     int    `json:"third_id"`     // 三级类目ID
}

// GetExtConfigRsp 获取小程序的第三方提交代码的页面配置响应
type GetExtConfigRsp struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	ExtConfig string `json:"ext_config"` // 第三方自定义的配置
}

// SetExtConfigRsp 设置小程序的第三方提交代码的页面配置响应
type SetExtConfigRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ==================== 服务市场相关 ====================

// InvokeServiceRsp 调用服务平台提供的服务响应
type InvokeServiceRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data    string `json:"data"` // 服务返回的数据
}

// ==================== 虚拟支付相关 ====================

// XpayQueryUserBalanceRsp 查询代币余额响应
type XpayQueryUserBalanceRsp struct {
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
	Balance        int    `json:"balance"`         // 代币总余额
	PresentBalance int    `json:"present_balance"` // 赠送代币余额
	SumSave        int    `json:"sum_save"`        // 历史总充值金额
	SumPresent     int    `json:"sum_present"`     // 历史总赠送金额
	SumBalance     int    `json:"sum_balance"`     // 历史总增加代币数
	SumCost        int    `json:"sum_cost"`        // 历史总消耗代币数
	FirstSaveFlag  bool   `json:"first_save_flag"` // 首充活动资格标识
}

// XpayCurrencyPayRsp 扣减代币响应
type XpayCurrencyPayRsp struct {
	Errcode           int    `json:"errcode"`
	Errmsg            string `json:"errmsg"`
	OrderId           string `json:"order_id"`            // 订单号
	Balance           int    `json:"balance"`             // 总余额（含付费和赠送）
	UsedPresentAmount int    `json:"used_present_amount"` // 使用的赠送代币数量
}

// XpayQueryOrderRsp 查询创建的订单响应
type XpayQueryOrderRsp struct {
	Errcode int        `json:"errcode"`
	Errmsg  string     `json:"errmsg"`
	Order   *XpayOrder `json:"order"`
}

// XpayOrder 虚拟支付订单信息
type XpayOrder struct {
	OrderId    string `json:"order_id"`    // 订单号
	Status     int    `json:"status"`      // 订单状态（0-10）
	OrderType  int    `json:"order_type"`  // 订单类型：0-标准虚拟支付，1-退款，7-Apple iOS，8-Apple iOS退款
	OrderFee   int    `json:"order_fee"`   // 订单金额（分）
	PaidFee    int    `json:"paid_fee"`    // 用户支付金额
	CreateTime int    `json:"create_time"` // 创建时间
	UpdateTime int    `json:"update_time"` // 更新时间
	PaidTime   int    `json:"paid_time"`   // 支付/退款时间（unix时间戳）
	WxOrderId  string `json:"wx_order_id"` // 微信内部订单号
	SettState  int    `json:"sett_state"`  // 结算状态（0-3）
}

// XpayCancelCurrencyPayRsp 代币支付退款响应
type XpayCancelCurrencyPayRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	OrderId string `json:"order_id"` // 退款订单号
}

// XpayPresentCurrencyRsp 代币赠送响应
type XpayPresentCurrencyRsp struct {
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
	Balance        int    `json:"balance"`         // 赠送后用户代币余额
	OrderId        string `json:"order_id"`        // 赠送订单号
	PresentBalance int    `json:"present_balance"` // 用户累计赠送代币数
}

// XpayDownloadBillRsp 下载小程序账单响应
type XpayDownloadBillRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Url     string `json:"url"` // 下载链接（30分钟有效）
}

// XpayRefundOrderRsp 启动订单退款任务响应
type XpayRefundOrderRsp struct {
	Errcode         int    `json:"errcode"`
	Errmsg          string `json:"errmsg"`
	RefundOrderId   string `json:"refund_order_id"`    // 退款订单号
	RefundWxOrderId string `json:"refund_wx_order_id"` // 微信侧退款订单号
	PayOrderId      string `json:"pay_order_id"`       // 关联的支付订单号
	PayWxOrderId    string `json:"pay_wx_order_id"`    // 关联的微信侧支付订单号
}

// XpayCreateWithdrawOrderRsp 创建提现单响应
type XpayCreateWithdrawOrderRsp struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	WithdrawNo   string `json:"withdraw_no"`    // 提现单号
	WxWithdrawNo string `json:"wx_withdraw_no"` // 微信侧提现单号
}

// XpayQueryWithdrawOrderRsp 查询提现单响应
type XpayQueryWithdrawOrderRsp struct {
	Errcode                  int    `json:"errcode"`
	Errmsg                   string `json:"errmsg"`
	WithdrawNo               string `json:"withdraw_no"`                // 提现单号
	Status                   int    `json:"status"`                     // 状态：1-处理中，2-成功，3-失败
	WithdrawAmount           string `json:"withdraw_amount"`            // 提现金额
	WxWithdrawNo             string `json:"wx_withdraw_no"`             // 微信侧提现单号
	WithdrawSuccessTimestamp string `json:"withdraw_success_timestamp"` // 成功时间戳（秒）
	CreateTime               string `json:"create_time"`                // 创建时间
	FailReason               string `json:"fail_reason"`                // 失败原因
}

// XpayQueryUploadGoodsRsp 查询批量上传道具任务响应
type XpayQueryUploadGoodsRsp struct {
	Errcode    int              `json:"errcode"`
	Errmsg     string           `json:"errmsg"`
	UploadItem []*XpayGoodsItem `json:"upload_item"` // 上传道具列表
	Status     int              `json:"status"`      // 任务状态：0-无任务，1-进行中，2-失败/部分失败，3-成功
}

// XpayGoodsItem 道具信息
type XpayGoodsItem struct {
	Id           string `json:"id"`            // 道具ID
	Name         string `json:"name"`          // 道具名称
	Price        int    `json:"price"`         // 价格（分）
	Remark       string `json:"remark"`        // 备注
	ItemUrl      string `json:"item_url"`      // 道具图片URL
	UploadStatus int    `json:"upload_status"` // 状态：0-上传中，1-ID已存在，2-成功，3-失败
	Errmsg       string `json:"errmsg"`        // 失败原因
}

// XpayQueryPublishGoodsRsp 查询批量发布道具任务响应
type XpayQueryPublishGoodsRsp struct {
	Errcode     int                `json:"errcode"`
	Errmsg      string             `json:"errmsg"`
	PublishItem []*XpayPublishItem `json:"publish_item"` // 发布道具列表
	Status      int                `json:"status"`       // 任务状态：0-无任务，1-进行中，2-失败/部分失败，3-成功
}

// XpayPublishItem 发布道具信息
type XpayPublishItem struct {
	Id            string `json:"id"`             // 道具ID
	PublishStatus int    `json:"publish_status"` // 状态：0-上传中，1-ID已存在，2-成功，3-失败
	Errmsg        string `json:"errmsg"`         // 失败原因
}

// XpayQueryBizBalanceRsp 查询商家账户可提现余额响应
type XpayQueryBizBalanceRsp struct {
	Errcode          int                   `json:"errcode"`
	Errmsg           string                `json:"errmsg"`
	BalanceAvailable *XpayBalanceAvailable `json:"balance_available"` // 可提现余额信息
}

// XpayBalanceAvailable 可提现余额信息
type XpayBalanceAvailable struct {
	Amount       string `json:"amount"`        // 可提现金额（元）
	CurrencyCode string `json:"currency_code"` // 货币代码（通常为CNY）
}

// XpayQueryTransferAccountRsp 查询广告金充值账户响应
type XpayQueryTransferAccountRsp struct {
	Errcode  int                    `json:"errcode"`
	Errmsg   string                 `json:"errmsg"`
	AcctList []*XpayTransferAccount `json:"acct_list"` // 充值账户列表
}

// XpayTransferAccount 广告金充值账户信息
type XpayTransferAccount struct {
	TransferAccountName       string `json:"transfer_account_name"`        // 充值账户名称
	TransferAccountUid        int    `json:"transfer_account_uid"`         // 充值账户UID
	TransferAccountAgencyId   int    `json:"transfer_account_agency_id"`   // 服务商账户ID
	TransferAccountAgencyName string `json:"transfer_account_agency_name"` // 服务商账户名称
	State                     int    `json:"state"`                        // 状态：0-待审核，1-通过，2-拒绝
	BindResult                int    `json:"bind_result"`                  // 绑定结果：1-成功，2-失败
	ErrorMsg                  string `json:"error_msg"`                    // 错误信息
}

// XpayQueryAdverFundsRsp 查询广告金发放记录响应
type XpayQueryAdverFundsRsp struct {
	Errcode        int              `json:"errcode"`
	Errmsg         string           `json:"errmsg"`
	AdverFundsList []*XpayAdverFund `json:"adver_funds_list"` // 发放记录列表
	TotalPage      int              `json:"total_page"`       // 总页数
}

// XpayAdverFund 广告金发放记录
type XpayAdverFund struct {
	SettleBegin  int    `json:"settle_begin"`  // 结算周期开始时间（unix时间戳）
	SettleEnd    int    `json:"settle_end"`    // 结算周期结束时间（unix时间戳）
	TotalAmount  int    `json:"total_amount"`  // 发放金额（分）
	RemainAmount int    `json:"remain_amount"` // 可用余额（分）
	ExpireTime   int    `json:"expire_time"`   // 过期时间（unix时间戳）
	FundType     int    `json:"fund_type"`     // 发放原因代码
	FundId       string `json:"fund_id"`       // 发放ID
}

// XpayCreateFundsBillRsp 充值广告金响应
type XpayCreateFundsBillRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	BillId  string `json:"bill_id"` // 充值单ID
}

// XpayQueryFundsBillRsp 查询广告金充值记录响应
type XpayQueryFundsBillRsp struct {
	Errcode   int              `json:"errcode"`
	Errmsg    string           `json:"errmsg"`
	BillList  []*XpayFundsBill `json:"bill_list"`  // 充值记录列表
	TotalPage int              `json:"total_page"` // 总页数
}

// XpayFundsBill 广告金充值记录
type XpayFundsBill struct {
	BillId              string `json:"bill_id"`               // 充值单ID
	OperTime            int    `json:"oper_time"`             // 充值时间（unix时间戳）
	SettleBegin         int    `json:"settle_begin"`          // 结算周期开始时间
	SettleEnd           int    `json:"settle_end"`            // 结算周期结束时间
	FundId              string `json:"fund_id"`               // 广告金ID
	TransferAccountName string `json:"transfer_account_name"` // 充值账户名称
	TransferAccountUid  int    `json:"transfer_account_uid"`  // 充值账户UID
	TransferAmount      int    `json:"transfer_amount"`       // 充值金额（分）
	Status              int    `json:"status"`                // 状态：0-处理中，1-成功，2-失败
	RequestId           string `json:"request_id"`            // 请求ID
}

// XpayQueryRecoverBillRsp 查询广告金回收记录响应
type XpayQueryRecoverBillRsp struct {
	Errcode   int                `json:"errcode"`
	Errmsg    string             `json:"errmsg"`
	BillList  []*XpayRecoverBill `json:"bill_list"`  // 回收记录列表
	TotalPage int                `json:"total_page"` // 总页数
}

// XpayRecoverBill 广告金回收记录
type XpayRecoverBill struct {
	BillId             string   `json:"bill_id"`              // 回收单ID
	RecoverTime        int      `json:"recover_time"`         // 回收时间（unix时间戳）
	SettleBegin        int      `json:"settle_begin"`         // 结算周期开始时间
	SettleEnd          int      `json:"settle_end"`           // 结算周期结束时间
	FundId             string   `json:"fund_id"`              // 对应广告金发放ID
	RecoverAccountName string   `json:"recover_account_name"` // 回收账户名称
	RecoverAmount      int      `json:"recover_amount"`       // 回收金额（分）
	RefundOrderList    []string `json:"refund_order_list"`    // 关联退款订单号列表
}

// XpayQuerySubscribeContractRsp 查询签约关系响应
type XpayQuerySubscribeContractRsp struct {
	Errcode            int    `json:"errcode"`
	Errmsg             string `json:"errmsg"`
	AuthorizationState string `json:"authorization_state"` // SIGNED-签约生效，TERMINATED-签约终止，UNBINDUSER-未签约
}

// XpayGetComplaintListRsp 获取投诉列表响应
type XpayGetComplaintListRsp struct {
	Errcode    int              `json:"errcode"`
	Errmsg     string           `json:"errmsg"`
	Total      int              `json:"total"`      // 总记录数
	Complaints []*XpayComplaint `json:"complaints"` // 投诉列表
}

// XpayComplaint 投诉信息
type XpayComplaint struct {
	ComplaintId           string                `json:"complaint_id"`            // 投诉ID
	ComplaintTime         string                `json:"complaint_time"`          // 投诉时间
	ComplaintDetail       string                `json:"complaint_detail"`        // 投诉描述
	ComplaintState        string                `json:"complaint_state"`         // 状态：PENDING、PROCESSING、PROCESSED
	PayerPhone            string                `json:"payer_phone"`             // 投诉人电话
	PayerOpenid           string                `json:"payer_openid"`            // 投诉人openid
	ComplaintOrderInfo    []*XpayComplaintOrder `json:"complaint_order_info"`    // 关联订单信息
	ComplaintFullRefunded bool                  `json:"complaint_full_refunded"` // 是否已全额退款
	IncomingUserResponse  bool                  `json:"incoming_user_response"`  // 是否有待处理的用户回复
	UserComplaintTimes    int                   `json:"user_complaint_times"`    // 用户投诉次数
	ComplaintMediaList    []*XpayComplaintMedia `json:"complaint_media_list"`    // 附件媒体列表
	ProblemDescription    string                `json:"problem_description"`     // 问题描述分类
	ProblemType           string                `json:"problem_type"`            // 类型：REFUND、SERVICE_NOT_WORK、OTHERS
	ApplyRefundAmount     int                   `json:"apply_refund_amount"`     // 申请退款金额（分）
	UserTagList           []string              `json:"user_tag_list"`           // 用户标签（TRUSTED、HIGH_RISK等）
	ServiceOrderInfo      []*XpayComplaintOrder `json:"service_order_info"`      // 关联服务订单信息
}

// XpayComplaintOrder 投诉关联订单
type XpayComplaintOrder struct {
	OrderId   string `json:"order_id"`
	WxOrderId string `json:"wx_order_id"`
}

// XpayComplaintMedia 投诉媒体信息
type XpayComplaintMedia struct {
	MediaType string   `json:"media_type"` // 媒体类型：USER_COMPLAINT_IMAGE、OPERATION_IMAGE
	MediaUrl  []string `json:"media_url"`  // 媒体URL列表
}

// XpayGetComplaintDetailRsp 获取投诉详情响应
type XpayGetComplaintDetailRsp struct {
	Errcode   int            `json:"errcode"`
	Errmsg    string         `json:"errmsg"`
	Complaint *XpayComplaint `json:"complaint"` // 投诉详情
}

// XpayGetNegotiationHistoryRsp 获取协商历史响应
type XpayGetNegotiationHistoryRsp struct {
	Errcode int                    `json:"errcode"`
	Errmsg  string                 `json:"errmsg"`
	Total   int                    `json:"total"`   // 总记录数
	History []*XpayNegotiationItem `json:"history"` // 协商历史列表
}

// XpayNegotiationItem 协商历史记录
type XpayNegotiationItem struct {
	LogId              string                `json:"log_id"`               // 操作流水ID
	Operator           string                `json:"operator"`             // 操作角色
	OperateTime        string                `json:"operate_time"`         // 操作时间
	OperateType        string                `json:"operate_type"`         // 操作类型
	OperateDetails     string                `json:"operate_details"`      // 操作内容
	ComplaintMediaList []*XpayComplaintMedia `json:"complaint_media_list"` // 附件媒体列表
}

// XpayUploadVpFileRsp 上传媒体文件响应
type XpayUploadVpFileRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	FileId  string `json:"file_id"` // 文件ID
}

// XpayGetUploadFileSignRsp 获取上传文件签名响应
type XpayGetUploadFileSignRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Sign    string `json:"sign"`    // Authorization头部值
	CosUrl  string `json:"cos_url"` // 转换后的URL（convert_cos=true时有效，30分钟有效）
}
