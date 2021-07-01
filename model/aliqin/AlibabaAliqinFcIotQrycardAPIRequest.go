package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询终端信息 API请求
alibaba.aliqin.fc.iot.qrycard

查询终端信息
*/
type AlibabaAliqinFcIotQrycardAPIRequest struct {
    model.Params
    // 外部计费来源
    _billSource   string
    // 外部计费号
    _billReal   string
    // ICCID
    _iccid   string
}

// 初始化AlibabaAliqinFcIotQrycardAPIRequest对象
func NewAlibabaAliqinFcIotQrycardRequest() *AlibabaAliqinFcIotQrycardAPIRequest{
    return &AlibabaAliqinFcIotQrycardAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.qrycard"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillSource Setter
// 外部计费来源
func (r *AlibabaAliqinFcIotQrycardAPIRequest) SetBillSource(_billSource string) error {
    r._billSource = _billSource
    r.Set("bill_source", _billSource)
    return nil
}

// BillSource Getter
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetBillSource() string {
    return r._billSource
}
// BillReal Setter
// 外部计费号
func (r *AlibabaAliqinFcIotQrycardAPIRequest) SetBillReal(_billReal string) error {
    r._billReal = _billReal
    r.Set("bill_real", _billReal)
    return nil
}

// BillReal Getter
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetBillReal() string {
    return r._billReal
}
// Iccid Setter
// ICCID
func (r *AlibabaAliqinFcIotQrycardAPIRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetIccid() string {
    return r._iccid
}