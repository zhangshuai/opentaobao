package qimen

import (
	"sync"
)

// StoreProcessConfirmRequest 结构体
type StoreProcessConfirmRequest struct {
	// 加工商品列表
	Materialitems []MaterialItem `json:"materialitems,omitempty" xml:"materialitems>material_item,omitempty"`
	// 加工商品列表
	Productitems []ProductItem `json:"productitems,omitempty" xml:"productitems>product_item,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 加工单编码
	ProcessOrderCode string `json:"processOrderCode,omitempty" xml:"processOrderCode,omitempty"`
	// 仓储系统加工单ID
	ProcessOrderId string `json:"processOrderId,omitempty" xml:"processOrderId,omitempty"`
	// 外部业务编码(一个合作伙伴中要求唯一多次确认时;每次传入要求唯一(一般传入WMS损益单据编码))
	OutBizCode string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 单据类型(CNJG=仓内加工作业单)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 加工单完成时间(YYYY-MM-DD HH:MM:SS)
	OrderCompleteTime string `json:"orderCompleteTime,omitempty" xml:"orderCompleteTime,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 仓库编码,统仓统配等无需ERP指定仓储编码的情况填OTHER
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 实际作业总数量
	ActualQty int64 `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStoreprocessConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolStoreProcessConfirmRequest = sync.Pool{
	New: func() any {
		return new(StoreProcessConfirmRequest)
	},
}

// GetStoreProcessConfirmRequest() 从对象池中获取StoreProcessConfirmRequest
func GetStoreProcessConfirmRequest() *StoreProcessConfirmRequest {
	return poolStoreProcessConfirmRequest.Get().(*StoreProcessConfirmRequest)
}

// ReleaseStoreProcessConfirmRequest 释放StoreProcessConfirmRequest
func ReleaseStoreProcessConfirmRequest(v *StoreProcessConfirmRequest) {
	v.Materialitems = v.Materialitems[:0]
	v.Productitems = v.Productitems[:0]
	v.OwnerCode = ""
	v.ProcessOrderCode = ""
	v.ProcessOrderId = ""
	v.OutBizCode = ""
	v.OrderType = ""
	v.OrderCompleteTime = ""
	v.Remark = ""
	v.WarehouseCode = ""
	v.ActualQty = 0
	v.ExtendProps = nil
	poolStoreProcessConfirmRequest.Put(v)
}
