package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefundfetchgoodsAPIResponse saas 售后逆向 商户发起逆向取货 API返回值
// alibaba.tcls.aelophy.refund.fetchgoods
//
// saas 售后逆向 商户发起逆向取货
type AlibabatclsaelophyrefundfetchgoodsAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophyrefundfetchgoodsAPIResponseModel
}

// AlibabatclsaelophyrefundfetchgoodsAPIResponseModel is saas 售后逆向 商户发起逆向取货 成功返回结果
type AlibabatclsaelophyrefundfetchgoodsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_fetchgoods_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabatclsaelophyrefundfetchgoodsApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
