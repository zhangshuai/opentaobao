package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankGroupListAPIRequest 图片银行分组信息获取 API请求
// alibaba.icbu.photobank.group.list
//
// 图片银行分组信息获取
type AlibabaIcbuPhotobankGroupListAPIRequest struct {
	model.Params
	// 补充信息
	_extraContext string
	// 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
	_id int64
}

// NewAlibabaIcbuPhotobankGroupListRequest 初始化AlibabaIcbuPhotobankGroupListAPIRequest对象
func NewAlibabaIcbuPhotobankGroupListRequest() *AlibabaIcbuPhotobankGroupListAPIRequest {
	return &AlibabaIcbuPhotobankGroupListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuPhotobankGroupListAPIRequest) Reset() {
	r._extraContext = ""
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.photobank.group.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtraContext is ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuPhotobankGroupListAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetId is Id Setter
// 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
func (r *AlibabaIcbuPhotobankGroupListAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetId() int64 {
	return r._id
}

var poolAlibabaIcbuPhotobankGroupListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuPhotobankGroupListRequest()
	},
}

// GetAlibabaIcbuPhotobankGroupListRequest 从 sync.Pool 获取 AlibabaIcbuPhotobankGroupListAPIRequest
func GetAlibabaIcbuPhotobankGroupListAPIRequest() *AlibabaIcbuPhotobankGroupListAPIRequest {
	return poolAlibabaIcbuPhotobankGroupListAPIRequest.Get().(*AlibabaIcbuPhotobankGroupListAPIRequest)
}

// ReleaseAlibabaIcbuPhotobankGroupListAPIRequest 将 AlibabaIcbuPhotobankGroupListAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuPhotobankGroupListAPIRequest(v *AlibabaIcbuPhotobankGroupListAPIRequest) {
	v.Reset()
	poolAlibabaIcbuPhotobankGroupListAPIRequest.Put(v)
}
