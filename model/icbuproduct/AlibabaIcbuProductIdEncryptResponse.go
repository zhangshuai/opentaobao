package icbuproduct

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ICBU国际站商品加密接口 API返回值 
alibaba.icbu.product.id.encrypt

ICBU国际站，对混淆的产品ID加密。
*/
type AlibabaIcbuProductIdEncryptAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuProductIdEncryptResponse
}

// ICBU国际站商品加密接口 成功返回结果
type AlibabaIcbuProductIdEncryptResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_product_id_encrypt_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 加密id
    SecretId   string `json:"secret_id,omitempty" xml:"secret_id,omitempty"`
}
