package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过id播放歌曲 API请求
taobao.ailab.aicloud.top.device.control.playbyid

通过id播放歌曲
*/
type TaobaoAilabAicloudTopDeviceControlPlaybyidRequest struct {
    model.Params
    // 用户信息
    param0   *OpenBaseInfo
    // 设备id
    param1   string
    // 音频id
    param2   string
    // 音频来源
    param3   string
    // 音频类型，如果没有音频类型默认填children_song
    param4   string
}

// 初始化TaobaoAilabAicloudTopDeviceControlPlaybyidRequest对象
func NewTaobaoAilabAicloudTopDeviceControlPlaybyidRequest() *TaobaoAilabAicloudTopDeviceControlPlaybyidRequest{
    return &TaobaoAilabAicloudTopDeviceControlPlaybyidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.playbyid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) GetParam1() string {
    return r.param1
}
// Param2 Setter
// 音频id
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) SetParam2(param2 string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) GetParam2() string {
    return r.param2
}
// Param3 Setter
// 音频来源
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) SetParam3(param3 string) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

// Param3 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) GetParam3() string {
    return r.param3
}
// Param4 Setter
// 音频类型，如果没有音频类型默认填children_song
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) SetParam4(param4 string) error {
    r.param4 = param4
    r.Set("param4", param4)
    return nil
}

// Param4 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidRequest) GetParam4() string {
    return r.param4
}
