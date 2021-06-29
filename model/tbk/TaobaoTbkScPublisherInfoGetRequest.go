package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户备案信息查询 API请求
taobao.tbk.sc.publisher.info.get

查询已生成的渠道id或会员运营id的相关信息。
*/
type TaobaoTbkScPublisherInfoGetRequest struct {
    model.Params
    // 类型，必选 1:渠道信息；2:会员信息
    infoType   int64
    // 渠道独占 - 渠道关系ID
    relationId   int64
    // 第几页
    pageNo   int64
    // 每页大小
    pageSize   int64
    // 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
    relationApp   string
    // 会员独占 - 会员运营ID
    specialId   string
    // 淘宝客外部用户标记，如自身系统账户ID；微信ID等
    externalId   string
    // 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
    externalType   int64
}

// 初始化TaobaoTbkScPublisherInfoGetRequest对象
func NewTaobaoTbkScPublisherInfoGetRequest() *TaobaoTbkScPublisherInfoGetRequest{
    return &TaobaoTbkScPublisherInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkScPublisherInfoGetRequest) GetApiMethodName() string {
    return "taobao.tbk.sc.publisher.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkScPublisherInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InfoType Setter
// 类型，必选 1:渠道信息；2:会员信息
func (r *TaobaoTbkScPublisherInfoGetRequest) SetInfoType(infoType int64) error {
    r.infoType = infoType
    r.Set("info_type", infoType)
    return nil
}

// InfoType Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetInfoType() int64 {
    return r.infoType
}
// RelationId Setter
// 渠道独占 - 渠道关系ID
func (r *TaobaoTbkScPublisherInfoGetRequest) SetRelationId(relationId int64) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

// RelationId Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetRelationId() int64 {
    return r.relationId
}
// PageNo Setter
// 第几页
func (r *TaobaoTbkScPublisherInfoGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoTbkScPublisherInfoGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// RelationApp Setter
// 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
func (r *TaobaoTbkScPublisherInfoGetRequest) SetRelationApp(relationApp string) error {
    r.relationApp = relationApp
    r.Set("relation_app", relationApp)
    return nil
}

// RelationApp Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetRelationApp() string {
    return r.relationApp
}
// SpecialId Setter
// 会员独占 - 会员运营ID
func (r *TaobaoTbkScPublisherInfoGetRequest) SetSpecialId(specialId string) error {
    r.specialId = specialId
    r.Set("special_id", specialId)
    return nil
}

// SpecialId Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetSpecialId() string {
    return r.specialId
}
// ExternalId Setter
// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
func (r *TaobaoTbkScPublisherInfoGetRequest) SetExternalId(externalId string) error {
    r.externalId = externalId
    r.Set("external_id", externalId)
    return nil
}

// ExternalId Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetExternalId() string {
    return r.externalId
}
// ExternalType Setter
// 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
func (r *TaobaoTbkScPublisherInfoGetRequest) SetExternalType(externalType int64) error {
    r.externalType = externalType
    r.Set("external_type", externalType)
    return nil
}

// ExternalType Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetExternalType() int64 {
    return r.externalType
}
