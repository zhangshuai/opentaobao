package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪商品各实体通用配置 API请求
taobao.xhotel.entity.config

飞猪商品各实体通用配置服务
*/
type TaobaoXhotelEntityConfigRequest struct {
    model.Params
    // 实体编码，例如酒店编码，房价编码
    entityCode   string
    // 格式为：[{"invoice_status":"1"},{"invoice_has":"1"},{"invoice_canvat":"1"},{"invoice_provider":"2"},{"invoice_channel":"1"},{"invoice_payer":"1"},{"invoice_type":"1,2,3"},{"invoice_desc":"测试发票描叙123"}]  key-value形式的配置信息其中invoice_status为发票状态，取值为0(无效),1(有效);invoice_has是否能开发票，1(能),0(不能);invoice_canvat是否能提供增值税专用发票1(能),0(不能)；invoice_provider发票提供方1(酒店提供),2(卖家邮寄)；invoice_channel开票渠道，目前仅支持1(线下开具)；invoice_payer邮费承担方，1(卖家承担)，2(买家承担-邮费到付);invoice_type开票类型 1(代订房费),2(旅游费) 3(差旅费) 4(卖家提供Email版收据)  如果多个请以英文逗号隔开，例如1,2,3。invoice_desc发票补充说明，如果没有可以填写无，长度不要超过20字符。该接口是全量更新的形式，每次更新必须要携带这些节点信息，如果想失效掉该rp上的发票信息，那么可以将invoice_status设置为0。
    configData   string
    // 默认是taobao。和酒店，房型等实体上的vendor保持一致。
    vendor   string
    // 实体id,例如酒店id,房型id(飞猪平台的id),rpid等，目前仅支持rpid
    entityId   int64
    // 实体类型：1：卖家维度，2：酒店维度，3：房型维度，4：宝贝维度，5：rp维度。请注意目前仅支持类型5
    type   int64
}

// 初始化TaobaoXhotelEntityConfigRequest对象
func NewTaobaoXhotelEntityConfigRequest() *TaobaoXhotelEntityConfigRequest{
    return &TaobaoXhotelEntityConfigRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelEntityConfigRequest) GetApiMethodName() string {
    return "taobao.xhotel.entity.config"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelEntityConfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntityCode Setter
// 实体编码，例如酒店编码，房价编码
func (r *TaobaoXhotelEntityConfigRequest) SetEntityCode(entityCode string) error {
    r.entityCode = entityCode
    r.Set("entity_code", entityCode)
    return nil
}

// EntityCode Getter
func (r TaobaoXhotelEntityConfigRequest) GetEntityCode() string {
    return r.entityCode
}
// ConfigData Setter
// 格式为：[{"invoice_status":"1"},{"invoice_has":"1"},{"invoice_canvat":"1"},{"invoice_provider":"2"},{"invoice_channel":"1"},{"invoice_payer":"1"},{"invoice_type":"1,2,3"},{"invoice_desc":"测试发票描叙123"}]  key-value形式的配置信息其中invoice_status为发票状态，取值为0(无效),1(有效);invoice_has是否能开发票，1(能),0(不能);invoice_canvat是否能提供增值税专用发票1(能),0(不能)；invoice_provider发票提供方1(酒店提供),2(卖家邮寄)；invoice_channel开票渠道，目前仅支持1(线下开具)；invoice_payer邮费承担方，1(卖家承担)，2(买家承担-邮费到付);invoice_type开票类型 1(代订房费),2(旅游费) 3(差旅费) 4(卖家提供Email版收据)  如果多个请以英文逗号隔开，例如1,2,3。invoice_desc发票补充说明，如果没有可以填写无，长度不要超过20字符。该接口是全量更新的形式，每次更新必须要携带这些节点信息，如果想失效掉该rp上的发票信息，那么可以将invoice_status设置为0。
func (r *TaobaoXhotelEntityConfigRequest) SetConfigData(configData string) error {
    r.configData = configData
    r.Set("config_data", configData)
    return nil
}

// ConfigData Getter
func (r TaobaoXhotelEntityConfigRequest) GetConfigData() string {
    return r.configData
}
// Vendor Setter
// 默认是taobao。和酒店，房型等实体上的vendor保持一致。
func (r *TaobaoXhotelEntityConfigRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelEntityConfigRequest) GetVendor() string {
    return r.vendor
}
// EntityId Setter
// 实体id,例如酒店id,房型id(飞猪平台的id),rpid等，目前仅支持rpid
func (r *TaobaoXhotelEntityConfigRequest) SetEntityId(entityId int64) error {
    r.entityId = entityId
    r.Set("entity_id", entityId)
    return nil
}

// EntityId Getter
func (r TaobaoXhotelEntityConfigRequest) GetEntityId() int64 {
    return r.entityId
}
// Type Setter
// 实体类型：1：卖家维度，2：酒店维度，3：房型维度，4：宝贝维度，5：rp维度。请注意目前仅支持类型5
func (r *TaobaoXhotelEntityConfigRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoXhotelEntityConfigRequest) GetType() int64 {
    return r.type
}