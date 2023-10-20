package ascp

import (
	"sync"
)

// PhysicsInventory 结构体
type PhysicsInventory struct {
	// 仓库编码，string（50)    卖家下唯一主键
	ErpWarehouseCode string `json:"erp_warehouse_code,omitempty" xml:"erp_warehouse_code,omitempty"`
	// ERP货品id，ERP内部系统货品唯一识别标识
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 仓可用正品库存数量
	AvaliableQuantity int64 `json:"avaliable_quantity,omitempty" xml:"avaliable_quantity,omitempty"`
	// 仓实际正品库存总数
	TotalQuantity int64 `json:"total_quantity,omitempty" xml:"total_quantity,omitempty"`
}

var poolPhysicsInventory = sync.Pool{
	New: func() any {
		return new(PhysicsInventory)
	},
}

// GetPhysicsInventory() 从对象池中获取PhysicsInventory
func GetPhysicsInventory() *PhysicsInventory {
	return poolPhysicsInventory.Get().(*PhysicsInventory)
}

// ReleasePhysicsInventory 释放PhysicsInventory
func ReleasePhysicsInventory(v *PhysicsInventory) {
	v.ErpWarehouseCode = ""
	v.ScItemId = ""
	v.AvaliableQuantity = 0
	v.TotalQuantity = 0
	poolPhysicsInventory.Put(v)
}
