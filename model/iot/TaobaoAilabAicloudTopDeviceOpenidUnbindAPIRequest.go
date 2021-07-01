package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openTaoBaoId解绑设备 API请求
taobao.ailab.aicloud.top.device.openid.unbind

openTaoBaoId解绑设备
*/
type TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest struct {
    model.Params
    // 淘宝openId
    _openId   string
    // 设备uuid
    _uuid   string
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
}

// 初始化TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceOpenidUnbindRequest() *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest{
    return &TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.openid.unbind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenId Setter
// 淘宝openId
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetOpenId(_openId string) error {
    r._openId = _openId
    r.Set("open_id", _openId)
    return nil
}

// OpenId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetOpenId() string {
    return r._openId
}
// Uuid Setter
// 设备uuid
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetUuid() string {
    return r._uuid
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetUserId() string {
    return r._userId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetExt() string {
    return r._ext
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetUtdId() string {
    return r._utdId
}