package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmemonotedeleteAPIRequest 天猫精灵备忘录删除 API请求
// taobao.ailab.aicloud.top.memo.note.delete
//
// 删除天猫精灵用户设置的备忘录
type TaobaoailabaicloudtopmemonotedeleteAPIRequest struct {
	model.Params
	// schema
	_schema string
	// 企业用户ID
	_userId string
	// 手持设备ID
	_utdId string
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// 备忘录ID
	_memoId int64
}

// NewTaobaoailabaicloudtopmemonotedeleteRequest 初始化TaobaoailabaicloudtopmemonotedeleteAPIRequest对象
func NewTaobaoailabaicloudtopmemonotedeleteRequest() *TaobaoailabaicloudtopmemonotedeleteAPIRequest {
	return &TaobaoailabaicloudtopmemonotedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmemonotedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.note.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmemonotedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmemonotedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoailabaicloudtopmemonotedeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopmemonotedeleteAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoailabaicloudtopmemonotedeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopmemonotedeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoailabaicloudtopmemonotedeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopmemonotedeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoailabaicloudtopmemonotedeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopmemonotedeleteAPIRequest) GetExt() string {
	return r._ext
}

// SetMemoId is MemoId Setter
// 备忘录ID
func (r *TaobaoailabaicloudtopmemonotedeleteAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoailabaicloudtopmemonotedeleteAPIRequest) GetMemoId() int64 {
	return r._memoId
}
