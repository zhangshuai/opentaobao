package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动列表接口 API请求
taobao.singletreasure.activity.query

查询活动列表接口
*/
type TaobaoSingletreasureActivityQueryRequest struct {
    model.Params
    // 查询对象
    query   *PageQueryDto
}

// 初始化TaobaoSingletreasureActivityQueryRequest对象
func NewTaobaoSingletreasureActivityQueryRequest() *TaobaoSingletreasureActivityQueryRequest{
    return &TaobaoSingletreasureActivityQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityQueryRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询对象
func (r *TaobaoSingletreasureActivityQueryRequest) SetQuery(query *PageQueryDto) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r TaobaoSingletreasureActivityQueryRequest) GetQuery() *PageQueryDto {
    return r.query
}