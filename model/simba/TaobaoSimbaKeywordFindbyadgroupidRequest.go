package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据推广单元id获取关键词 API请求
taobao.simba.keyword.findbyadgroupid

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordFindbyadgroupidRequest struct {
    model.Params
    // 推广单元id
    adgroupId   int64
}

// 初始化TaobaoSimbaKeywordFindbyadgroupidRequest对象
func NewTaobaoSimbaKeywordFindbyadgroupidRequest() *TaobaoSimbaKeywordFindbyadgroupidRequest{
    return &TaobaoSimbaKeywordFindbyadgroupidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordFindbyadgroupidRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.findbyadgroupid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordFindbyadgroupidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaKeywordFindbyadgroupidRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaKeywordFindbyadgroupidRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
