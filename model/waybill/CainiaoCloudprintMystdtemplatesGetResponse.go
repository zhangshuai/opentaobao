package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户使用的菜鸟电子面单模板信息 APIResponse
cainiao.cloudprint.mystdtemplates.get

获取用户使用的菜鸟电子面单
*/
type CainiaoCloudprintMystdtemplatesGetAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCloudprintMystdtemplatesGetResponse `json:"cainiao_cloudprint_mystdtemplates_get_response,omitempty"`
}

type CainiaoCloudprintMystdtemplatesGetResponse struct {

    // 返回结果
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}