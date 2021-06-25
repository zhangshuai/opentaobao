package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改产品线 APIResponse
taobao.fenxiao.productcat.update

修改产品线
*/
type TaobaoFenxiaoProductcatUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductcatUpdateResponse `json:"taobao_fenxiao_productcat_update_response,omitempty"`
}

type TaobaoFenxiaoProductcatUpdateResponse struct {

    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}