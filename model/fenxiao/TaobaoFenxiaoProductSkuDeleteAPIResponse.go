package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品SKU删除接口 API返回值 
taobao.fenxiao.product.sku.delete

根据sku properties删除sku数据
*/
type TaobaoFenxiaoProductSkuDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductSkuDeleteAPIResponseModel
}

// 产品SKU删除接口 成功返回结果
type TaobaoFenxiaoProductSkuDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"fenxiao_product_sku_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作结果
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    // 操作时间
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
}