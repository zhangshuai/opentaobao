package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消上门揽件 API请求
tmall.nr.fulfill.cancel

新零售门店业务取消上门揽件
*/
type TmallNrFulfillCancelRequest struct {
    model.Params
    // 入参
    req   *NrCancelFulfillReqDto
}

// 初始化TmallNrFulfillCancelRequest对象
func NewTmallNrFulfillCancelRequest() *TmallNrFulfillCancelRequest{
    return &TmallNrFulfillCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrFulfillCancelRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrFulfillCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Req Setter
// 入参
func (r *TmallNrFulfillCancelRequest) SetReq(req *NrCancelFulfillReqDto) error {
    r.req = req
    r.Set("req", req)
    return nil
}

// Req Getter
func (r TmallNrFulfillCancelRequest) GetReq() *NrCancelFulfillReqDto {
    return r.req
}