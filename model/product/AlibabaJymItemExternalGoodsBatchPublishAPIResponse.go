package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchPublishAPIResponse 交易猫外部商家批量发布商品接口 API返回值
// alibaba.jym.item.external.goods.batch.publish
//
// 供外部B端商家接入，提交批量发布商品请求，返回批量发布任务结果
type AlibabaJymItemExternalGoodsBatchPublishAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemExternalGoodsBatchPublishAPIResponseModel
}

// AlibabaJymItemExternalGoodsBatchPublishAPIResponseModel is 交易猫外部商家批量发布商品接口 成功返回结果
type AlibabaJymItemExternalGoodsBatchPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_batch_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 商品批量发布接口返回
	Result *GoodsBatchResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}