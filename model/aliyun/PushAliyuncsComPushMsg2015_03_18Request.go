package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送 API请求
push.aliyuncs.com.pushMsg.2015-03-18

消息推送  ,支持指定用户/账号/广播等模式
*/
type PushAliyuncsComPushMsg2015_03_18Request struct {
    model.Params
    // 用户账号列表,以换行区分,仅sendType为3时有效
    account   string
    // 防打扰时长,取值范围为1~23
    antiHarassDuration   int64
    // 防打扰开始时间点,取值范围为0~23
    antiHarassStartTime   int64
    // 应用标识
    appId   int64
    // 批次编号,用于统计活动推送效果
    batchNumber   string
    // 消息体,UTF-8编码
    body   string
    // 设备编号列表,以换行区分,仅sendType为4时有效
    deviceId   string
    // 设备类型,取值范围为:0~3云推送支持多种设备, 各种设备类型编号如下:IOS设备:deviceType&amp;1=1; Andriod设备:deviceType&amp;2=2;如果存在此字段,则 向指定的设备类型推送消息。默认为全部(3);
    deviceType   int64
    // 推送时间,若空表示立即推送,推送时间不能早于当前 时间
    pushTime   string
    // 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根 据用户账号列表文件发送消息4:指定设备,根据设备编 码列表文件发送消息默认值为1
    sendType   int64
    // 标签名称,仅支持1个标签,仅sendType为2时有效
    tag   string
    // 离线消息保存时长,取值范围为1~72,若不填,则表 示不保存离线消息
    timeout   int64
    // 标题
    title   string
}

// 初始化PushAliyuncsComPushMsg2015_03_18Request对象
func NewPushAliyuncsComPushMsg2015_03_18Request() *PushAliyuncsComPushMsg2015_03_18Request{
    return &PushAliyuncsComPushMsg2015_03_18Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r PushAliyuncsComPushMsg2015_03_18Request) GetApiMethodName() string {
    return "push.aliyuncs.com.pushMsg.2015-03-18"
}

// IRequest interface 方法, 获取API参数
func (r PushAliyuncsComPushMsg2015_03_18Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Account Setter
// 用户账号列表,以换行区分,仅sendType为3时有效
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetAccount(account string) error {
    r.account = account
    r.Set("Account", account)
    return nil
}

// Account Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetAccount() string {
    return r.account
}
// AntiHarassDuration Setter
// 防打扰时长,取值范围为1~23
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetAntiHarassDuration(antiHarassDuration int64) error {
    r.antiHarassDuration = antiHarassDuration
    r.Set("AntiHarassDuration", antiHarassDuration)
    return nil
}

// AntiHarassDuration Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetAntiHarassDuration() int64 {
    return r.antiHarassDuration
}
// AntiHarassStartTime Setter
// 防打扰开始时间点,取值范围为0~23
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetAntiHarassStartTime(antiHarassStartTime int64) error {
    r.antiHarassStartTime = antiHarassStartTime
    r.Set("AntiHarassStartTime", antiHarassStartTime)
    return nil
}

// AntiHarassStartTime Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetAntiHarassStartTime() int64 {
    return r.antiHarassStartTime
}
// AppId Setter
// 应用标识
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetAppId(appId int64) error {
    r.appId = appId
    r.Set("AppId", appId)
    return nil
}

// AppId Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetAppId() int64 {
    return r.appId
}
// BatchNumber Setter
// 批次编号,用于统计活动推送效果
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetBatchNumber(batchNumber string) error {
    r.batchNumber = batchNumber
    r.Set("BatchNumber", batchNumber)
    return nil
}

// BatchNumber Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetBatchNumber() string {
    return r.batchNumber
}
// Body Setter
// 消息体,UTF-8编码
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetBody(body string) error {
    r.body = body
    r.Set("Body", body)
    return nil
}

// Body Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetBody() string {
    return r.body
}
// DeviceId Setter
// 设备编号列表,以换行区分,仅sendType为4时有效
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("DeviceId", deviceId)
    return nil
}

// DeviceId Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetDeviceId() string {
    return r.deviceId
}
// DeviceType Setter
// 设备类型,取值范围为:0~3云推送支持多种设备, 各种设备类型编号如下:IOS设备:deviceType&amp;1=1; Andriod设备:deviceType&amp;2=2;如果存在此字段,则 向指定的设备类型推送消息。默认为全部(3);
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("DeviceType", deviceType)
    return nil
}

// DeviceType Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetDeviceType() int64 {
    return r.deviceType
}
// PushTime Setter
// 推送时间,若空表示立即推送,推送时间不能早于当前 时间
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetPushTime(pushTime string) error {
    r.pushTime = pushTime
    r.Set("PushTime", pushTime)
    return nil
}

// PushTime Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetPushTime() string {
    return r.pushTime
}
// SendType Setter
// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根 据用户账号列表文件发送消息4:指定设备,根据设备编 码列表文件发送消息默认值为1
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetSendType(sendType int64) error {
    r.sendType = sendType
    r.Set("SendType", sendType)
    return nil
}

// SendType Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetSendType() int64 {
    return r.sendType
}
// Tag Setter
// 标签名称,仅支持1个标签,仅sendType为2时有效
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetTag(tag string) error {
    r.tag = tag
    r.Set("Tag", tag)
    return nil
}

// Tag Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetTag() string {
    return r.tag
}
// Timeout Setter
// 离线消息保存时长,取值范围为1~72,若不填,则表 示不保存离线消息
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetTimeout(timeout int64) error {
    r.timeout = timeout
    r.Set("Timeout", timeout)
    return nil
}

// Timeout Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetTimeout() int64 {
    return r.timeout
}
// Title Setter
// 标题
func (r *PushAliyuncsComPushMsg2015_03_18Request) SetTitle(title string) error {
    r.title = title
    r.Set("Title", title)
    return nil
}

// Title Getter
func (r PushAliyuncsComPushMsg2015_03_18Request) GetTitle() string {
    return r.title
}
