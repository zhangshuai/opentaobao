package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest 负卖库存失效接口 API请求
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate
//
// 失效负卖库存数据
type AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest struct {
	model.Params
	// 入参
	_futureInventoryMainOperationQuest *Futureinventorymainoperationquest
}

// NewAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest 初始化AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest对象
func NewAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest {
	return &AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) Reset() {
	r._futureInventoryMainOperationQuest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFutureInventoryMainOperationQuest is FutureInventoryMainOperationQuest Setter
// 入参
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) SetFutureInventoryMainOperationQuest(_futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
	r._futureInventoryMainOperationQuest = _futureInventoryMainOperationQuest
	r.Set("future_inventory_main_operation_quest", _futureInventoryMainOperationQuest)
	return nil
}

// GetFutureInventoryMainOperationQuest FutureInventoryMainOperationQuest Getter
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
	return r._futureInventoryMainOperationQuest
}

var poolAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest()
	},
}

// GetAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest 从 sync.Pool 获取 AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest
func GetAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest {
	return poolAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest.Get().(*AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest)
}

// ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest 将 AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest(v *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) {
	v.Reset()
	poolAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest.Put(v)
}
