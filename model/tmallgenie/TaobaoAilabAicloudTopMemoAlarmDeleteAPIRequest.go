package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest 天猫精灵闹钟删除 API请求
// taobao.ailab.aicloud.top.memo.alarm.delete
//
// 天猫精灵闹钟删除
type TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest struct {
	model.Params
	// schema
	_schema string
	// 手持设备ID
	_utdId string
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// 企业用户ID
	_userId string
	// 闹钟ID
	_memoId int64
}

// NewTaobaoAilabAicloudTopMemoAlarmDeleteRequest 初始化TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest对象
func NewTaobaoAilabAicloudTopMemoAlarmDeleteRequest() *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest {
	return &TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.alarm.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// Get Schema Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetSchema() string {
	return r._schema
}

// Set is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// Get UtdId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// Set is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// Get Ext Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetExt() string {
	return r._ext
}

// Set is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetUserId() string {
	return r._userId
}

// Set is MemoId Setter
// 闹钟ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// Get MemoId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetMemoId() int64 {
	return r._memoId
}
