package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpwarehouseinventoryqueryAPIRequest AliExpress在仓库存查询API API请求
// aliexpress.ascp.warehouse.inventory.query
//
// AliExpress在仓库存查询查询API
type AliexpressascpwarehouseinventoryqueryAPIRequest struct {
	model.Params
	// 查询DTO
	_warehouseInventoryQueryDto *WarehouseInventoryQueryDto
}

// NewAliexpressascpwarehouseinventoryqueryRequest 初始化AliexpressascpwarehouseinventoryqueryAPIRequest对象
func NewAliexpressascpwarehouseinventoryqueryRequest() *AliexpressascpwarehouseinventoryqueryAPIRequest {
	return &AliexpressascpwarehouseinventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascpwarehouseinventoryqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.warehouse.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascpwarehouseinventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascpwarehouseinventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseInventoryQueryDto is WarehouseInventoryQueryDto Setter
// 查询DTO
func (r *AliexpressascpwarehouseinventoryqueryAPIRequest) SetWarehouseInventoryQueryDto(_warehouseInventoryQueryDto *WarehouseInventoryQueryDto) error {
	r._warehouseInventoryQueryDto = _warehouseInventoryQueryDto
	r.Set("warehouse_inventory_query_dto", _warehouseInventoryQueryDto)
	return nil
}

// GetWarehouseInventoryQueryDto WarehouseInventoryQueryDto Getter
func (r AliexpressascpwarehouseinventoryqueryAPIRequest) GetWarehouseInventoryQueryDto() *WarehouseInventoryQueryDto {
	return r._warehouseInventoryQueryDto
}
