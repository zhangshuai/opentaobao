package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地点关联关系增量编辑 API请求
taobao.location.relation.edit

地点关联关系增量编辑
*/
type TaobaoLocationRelationEditAPIRequest struct {
    model.Params
    // 关系对象列表
    _locationRelationList   []LocationRelationDto
}

// 初始化TaobaoLocationRelationEditAPIRequest对象
func NewTaobaoLocationRelationEditRequest() *TaobaoLocationRelationEditAPIRequest{
    return &TaobaoLocationRelationEditAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLocationRelationEditAPIRequest) GetApiMethodName() string {
    return "taobao.location.relation.edit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLocationRelationEditAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LocationRelationList Setter
// 关系对象列表
func (r *TaobaoLocationRelationEditAPIRequest) SetLocationRelationList(_locationRelationList []LocationRelationDto) error {
    r._locationRelationList = _locationRelationList
    r.Set("location_relation_list", _locationRelationList)
    return nil
}

// LocationRelationList Getter
func (r TaobaoLocationRelationEditAPIRequest) GetLocationRelationList() []LocationRelationDto {
    return r._locationRelationList
}