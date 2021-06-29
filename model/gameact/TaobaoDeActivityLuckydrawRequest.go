package gameact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖 API请求
taobao.de.activity.luckydraw

用于激励平台对外提供抽奖功能，包括但不限于集分宝、红包、宝点、淘金币、淘彩票等
*/
type TaobaoDeActivityLuckydrawRequest struct {
    model.Params
    // 运营和cp约定的事件唯一标示
    eventKey   string
    // 时间戳
    sequenceId   int64
    // 用户的串ID
    accountId   string
    // 机器设备号
    machineId   string
    // 确认签名key
    confirmKey   string
    // 行为Key
    behaviorKey   string
    // 渠道
    channel   string
    // 使用市场
    market   string
    // 盒型号
    deviceModel   string
    // 魔盒分发渠道
    distribChannel   string
    // 魔盒UUID
    uuid   string
}

// 初始化TaobaoDeActivityLuckydrawRequest对象
func NewTaobaoDeActivityLuckydrawRequest() *TaobaoDeActivityLuckydrawRequest{
    return &TaobaoDeActivityLuckydrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityLuckydrawRequest) GetApiMethodName() string {
    return "taobao.de.activity.luckydraw"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityLuckydrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EventKey Setter
// 运营和cp约定的事件唯一标示
func (r *TaobaoDeActivityLuckydrawRequest) SetEventKey(eventKey string) error {
    r.eventKey = eventKey
    r.Set("event_key", eventKey)
    return nil
}

// EventKey Getter
func (r TaobaoDeActivityLuckydrawRequest) GetEventKey() string {
    return r.eventKey
}
// SequenceId Setter
// 时间戳
func (r *TaobaoDeActivityLuckydrawRequest) SetSequenceId(sequenceId int64) error {
    r.sequenceId = sequenceId
    r.Set("sequence_id", sequenceId)
    return nil
}

// SequenceId Getter
func (r TaobaoDeActivityLuckydrawRequest) GetSequenceId() int64 {
    return r.sequenceId
}
// AccountId Setter
// 用户的串ID
func (r *TaobaoDeActivityLuckydrawRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r TaobaoDeActivityLuckydrawRequest) GetAccountId() string {
    return r.accountId
}
// MachineId Setter
// 机器设备号
func (r *TaobaoDeActivityLuckydrawRequest) SetMachineId(machineId string) error {
    r.machineId = machineId
    r.Set("machine_id", machineId)
    return nil
}

// MachineId Getter
func (r TaobaoDeActivityLuckydrawRequest) GetMachineId() string {
    return r.machineId
}
// ConfirmKey Setter
// 确认签名key
func (r *TaobaoDeActivityLuckydrawRequest) SetConfirmKey(confirmKey string) error {
    r.confirmKey = confirmKey
    r.Set("confirm_key", confirmKey)
    return nil
}

// ConfirmKey Getter
func (r TaobaoDeActivityLuckydrawRequest) GetConfirmKey() string {
    return r.confirmKey
}
// BehaviorKey Setter
// 行为Key
func (r *TaobaoDeActivityLuckydrawRequest) SetBehaviorKey(behaviorKey string) error {
    r.behaviorKey = behaviorKey
    r.Set("behavior_key", behaviorKey)
    return nil
}

// BehaviorKey Getter
func (r TaobaoDeActivityLuckydrawRequest) GetBehaviorKey() string {
    return r.behaviorKey
}
// Channel Setter
// 渠道
func (r *TaobaoDeActivityLuckydrawRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TaobaoDeActivityLuckydrawRequest) GetChannel() string {
    return r.channel
}
// Market Setter
// 使用市场
func (r *TaobaoDeActivityLuckydrawRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

// Market Getter
func (r TaobaoDeActivityLuckydrawRequest) GetMarket() string {
    return r.market
}
// DeviceModel Setter
// 盒型号
func (r *TaobaoDeActivityLuckydrawRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

// DeviceModel Getter
func (r TaobaoDeActivityLuckydrawRequest) GetDeviceModel() string {
    return r.deviceModel
}
// DistribChannel Setter
// 魔盒分发渠道
func (r *TaobaoDeActivityLuckydrawRequest) SetDistribChannel(distribChannel string) error {
    r.distribChannel = distribChannel
    r.Set("distrib_channel", distribChannel)
    return nil
}

// DistribChannel Getter
func (r TaobaoDeActivityLuckydrawRequest) GetDistribChannel() string {
    return r.distribChannel
}
// Uuid Setter
// 魔盒UUID
func (r *TaobaoDeActivityLuckydrawRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r TaobaoDeActivityLuckydrawRequest) GetUuid() string {
    return r.uuid
}
