package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieTracksSearchAPIRequest 查询音频 API请求
// alibaba.ailabs.aligenie.tracks.search
//
// 搜索类目下的音频信息
type AlibabaAilabsAligenieTracksSearchAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 一级类目，如：儿童、新闻、商业财经有声书等
	_param1 string
	// 二级类目，如儿童下有：儿歌、童谣、国学等
	_param2 string
	// 搜索的单个音频名称
	_param3 string
	// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
	_param4 int64
	// 每页数量（不超过50）
	_param5 int64
}

// NewAlibabaAilabsAligenieTracksSearchRequest 初始化AlibabaAilabsAligenieTracksSearchAPIRequest对象
func NewAlibabaAilabsAligenieTracksSearchRequest() *AlibabaAilabsAligenieTracksSearchAPIRequest {
	return &AlibabaAilabsAligenieTracksSearchAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._param1 = ""
	r._param2 = ""
	r._param3 = ""
	r._param4 = 0
	r._param5 = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.tracks.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetExt() string {
	return r._ext
}

// SetParam1 is Param1 Setter
// 一级类目，如：儿童、新闻、商业财经有声书等
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam2 is Param2 Setter
// 二级类目，如儿童下有：儿歌、童谣、国学等
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetParam2(_param2 string) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetParam2() string {
	return r._param2
}

// SetParam3 is Param3 Setter
// 搜索的单个音频名称
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetParam3(_param3 string) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// GetParam3 Param3 Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetParam3() string {
	return r._param3
}

// SetParam4 is Param4 Setter
// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetParam4(_param4 int64) error {
	r._param4 = _param4
	r.Set("param4", _param4)
	return nil
}

// GetParam4 Param4 Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetParam4() int64 {
	return r._param4
}

// SetParam5 is Param5 Setter
// 每页数量（不超过50）
func (r *AlibabaAilabsAligenieTracksSearchAPIRequest) SetParam5(_param5 int64) error {
	r._param5 = _param5
	r.Set("param5", _param5)
	return nil
}

// GetParam5 Param5 Getter
func (r AlibabaAilabsAligenieTracksSearchAPIRequest) GetParam5() int64 {
	return r._param5
}

var poolAlibabaAilabsAligenieTracksSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsAligenieTracksSearchRequest()
	},
}

// GetAlibabaAilabsAligenieTracksSearchRequest 从 sync.Pool 获取 AlibabaAilabsAligenieTracksSearchAPIRequest
func GetAlibabaAilabsAligenieTracksSearchAPIRequest() *AlibabaAilabsAligenieTracksSearchAPIRequest {
	return poolAlibabaAilabsAligenieTracksSearchAPIRequest.Get().(*AlibabaAilabsAligenieTracksSearchAPIRequest)
}

// ReleaseAlibabaAilabsAligenieTracksSearchAPIRequest 将 AlibabaAilabsAligenieTracksSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsAligenieTracksSearchAPIRequest(v *AlibabaAilabsAligenieTracksSearchAPIRequest) {
	v.Reset()
	poolAlibabaAilabsAligenieTracksSearchAPIRequest.Put(v)
}
