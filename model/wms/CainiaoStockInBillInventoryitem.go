package wms

import (
	"sync"
)

// CainiaoStockInBillInventoryitem 结构体
type CainiaoStockInBillInventoryitem struct {
	// 商品保质期信息，失效日期
	DueDate string `json:"due_date,omitempty" xml:"due_date,omitempty"`
	// 商品保质期信息，生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 生产编码，同一商品可能因商家不同有不同编码
	ProduceCode string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
	// 生产地区，仓库采集的商品信息，供记录参考
	ProduceArea string `json:"produce_area,omitempty" xml:"produce_area,omitempty"`
	// 批次号
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 库存类型1 可销售库存 101残次品
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 数量
	ItemQty int64 `json:"item_qty,omitempty" xml:"item_qty,omitempty"`
}

var poolCainiaoStockInBillInventoryitem = sync.Pool{
	New: func() any {
		return new(CainiaoStockInBillInventoryitem)
	},
}

// GetCainiaoStockInBillInventoryitem() 从对象池中获取CainiaoStockInBillInventoryitem
func GetCainiaoStockInBillInventoryitem() *CainiaoStockInBillInventoryitem {
	return poolCainiaoStockInBillInventoryitem.Get().(*CainiaoStockInBillInventoryitem)
}

// ReleaseCainiaoStockInBillInventoryitem 释放CainiaoStockInBillInventoryitem
func ReleaseCainiaoStockInBillInventoryitem(v *CainiaoStockInBillInventoryitem) {
	v.DueDate = ""
	v.ProduceDate = ""
	v.ProduceCode = ""
	v.ProduceArea = ""
	v.BatchCode = ""
	v.InventoryType = 0
	v.ItemQty = 0
	poolCainiaoStockInBillInventoryitem.Put(v)
}
