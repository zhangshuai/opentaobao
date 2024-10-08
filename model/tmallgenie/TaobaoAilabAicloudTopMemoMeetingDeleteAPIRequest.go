package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest 天猫精灵会议删除 API请求
// taobao.ailab.aicloud.top.memo.meeting.delete
//
// 天猫精灵会议删除
type TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest struct {
	model.Params
	// schema
	_schema string
	// 企业用户ID
	_userId string
	// 手持设备ID
	_utdId string
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// 会议ID
	_memoId int64
}

// NewTaobaoAilabAicloudTopMemoMeetingDeleteRequest 初始化TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest对象
func NewTaobaoAilabAicloudTopMemoMeetingDeleteRequest() *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest {
	return &TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._memoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.meeting.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) GetExt() string {
	return r._ext
}

// SetMemoId is MemoId Setter
// 会议ID
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) GetMemoId() int64 {
	return r._memoId
}

var poolTaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMemoMeetingDeleteRequest()
	},
}

// GetTaobaoAilabAicloudTopMemoMeetingDeleteRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest
func GetTaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest() *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest {
	return poolTaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest.Get().(*TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest 将 TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest(v *TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest.Put(v)
}
