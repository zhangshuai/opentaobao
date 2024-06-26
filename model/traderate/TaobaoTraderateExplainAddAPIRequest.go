package traderate

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTraderateExplainAddAPIRequest 商城评价解释接口 API请求
// taobao.traderate.explain.add
//
// 商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
type TaobaoTraderateExplainAddAPIRequest struct {
	model.Params
	// 评价解释内容，最大长度：500个汉字
	_reply string
	// 子订单ID
	_oid int64
}

// NewTaobaoTraderateExplainAddRequest 初始化TaobaoTraderateExplainAddAPIRequest对象
func NewTaobaoTraderateExplainAddRequest() *TaobaoTraderateExplainAddAPIRequest {
	return &TaobaoTraderateExplainAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTraderateExplainAddAPIRequest) Reset() {
	r._reply = ""
	r._oid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTraderateExplainAddAPIRequest) GetApiMethodName() string {
	return "taobao.traderate.explain.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTraderateExplainAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTraderateExplainAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReply is Reply Setter
// 评价解释内容，最大长度：500个汉字
func (r *TaobaoTraderateExplainAddAPIRequest) SetReply(_reply string) error {
	r._reply = _reply
	r.Set("reply", _reply)
	return nil
}

// GetReply Reply Getter
func (r TaobaoTraderateExplainAddAPIRequest) GetReply() string {
	return r._reply
}

// SetOid is Oid Setter
// 子订单ID
func (r *TaobaoTraderateExplainAddAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// GetOid Oid Getter
func (r TaobaoTraderateExplainAddAPIRequest) GetOid() int64 {
	return r._oid
}

var poolTaobaoTraderateExplainAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTraderateExplainAddRequest()
	},
}

// GetTaobaoTraderateExplainAddRequest 从 sync.Pool 获取 TaobaoTraderateExplainAddAPIRequest
func GetTaobaoTraderateExplainAddAPIRequest() *TaobaoTraderateExplainAddAPIRequest {
	return poolTaobaoTraderateExplainAddAPIRequest.Get().(*TaobaoTraderateExplainAddAPIRequest)
}

// ReleaseTaobaoTraderateExplainAddAPIRequest 将 TaobaoTraderateExplainAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoTraderateExplainAddAPIRequest(v *TaobaoTraderateExplainAddAPIRequest) {
	v.Reset()
	poolTaobaoTraderateExplainAddAPIRequest.Put(v)
}
