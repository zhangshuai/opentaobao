package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdGetAPIRequest 淘宝客-服务商-获取人群标签 API请求
// taobao.tbk.sc.ucrowd.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群信息，包括人群名称描述及覆盖会员数。
type TaobaoTbkScUcrowdGetAPIRequest struct {
	model.Params
	// 人群标签id
	_ucrowdId int64
}

// NewTaobaoTbkScUcrowdGetRequest 初始化TaobaoTbkScUcrowdGetAPIRequest对象
func NewTaobaoTbkScUcrowdGetRequest() *TaobaoTbkScUcrowdGetAPIRequest {
	return &TaobaoTbkScUcrowdGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScUcrowdGetAPIRequest) Reset() {
	r._ucrowdId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScUcrowdGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScUcrowdGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScUcrowdGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUcrowdId is UcrowdId Setter
// 人群标签id
func (r *TaobaoTbkScUcrowdGetAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaoTbkScUcrowdGetAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}

var poolTaobaoTbkScUcrowdGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScUcrowdGetRequest()
	},
}

// GetTaobaoTbkScUcrowdGetRequest 从 sync.Pool 获取 TaobaoTbkScUcrowdGetAPIRequest
func GetTaobaoTbkScUcrowdGetAPIRequest() *TaobaoTbkScUcrowdGetAPIRequest {
	return poolTaobaoTbkScUcrowdGetAPIRequest.Get().(*TaobaoTbkScUcrowdGetAPIRequest)
}

// ReleaseTaobaoTbkScUcrowdGetAPIRequest 将 TaobaoTbkScUcrowdGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScUcrowdGetAPIRequest(v *TaobaoTbkScUcrowdGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkScUcrowdGetAPIRequest.Put(v)
}
