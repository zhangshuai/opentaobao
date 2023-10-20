package ascpffo

import (
	"sync"
)

// OnWayInventoryQueryDto 结构体
type OnWayInventoryQueryDto struct {
	// 货品Id列表，最多30个
	ScItemIdList []int64 `json:"sc_item_id_list,omitempty" xml:"sc_item_id_list>int64,omitempty"`
	// 入库仓编码
	InboundStoreCode string `json:"inbound_store_code,omitempty" xml:"inbound_store_code,omitempty"`
	// 出库仓编码
	OutboundStoreCode string `json:"outbound_store_code,omitempty" xml:"outbound_store_code,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 库存类型(1 采购在途，2 调拨在途，3 销售在途，4 销退在途)
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
}

var poolOnWayInventoryQueryDto = sync.Pool{
	New: func() any {
		return new(OnWayInventoryQueryDto)
	},
}

// GetOnWayInventoryQueryDto() 从对象池中获取OnWayInventoryQueryDto
func GetOnWayInventoryQueryDto() *OnWayInventoryQueryDto {
	return poolOnWayInventoryQueryDto.Get().(*OnWayInventoryQueryDto)
}

// ReleaseOnWayInventoryQueryDto 释放OnWayInventoryQueryDto
func ReleaseOnWayInventoryQueryDto(v *OnWayInventoryQueryDto) {
	v.ScItemIdList = v.ScItemIdList[:0]
	v.InboundStoreCode = ""
	v.OutboundStoreCode = ""
	v.BizType = 0
	v.InventoryType = 0
	poolOnWayInventoryQueryDto.Put(v)
}
