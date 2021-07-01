package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加技能 API请求
yunos.tvpubadmin.device.yks.skill.add

添加技能
*/
type YunosTvpubadminDeviceYksSkillAddAPIRequest struct {
    model.Params
    // 技能id
    _skillId   int64
    // 设备id
    _botId   int64
    // 技能名称
    _name   string
    // 图片地址
    _iconImageUrl   string
}

// 初始化YunosTvpubadminDeviceYksSkillAddAPIRequest对象
func NewYunosTvpubadminDeviceYksSkillAddRequest() *YunosTvpubadminDeviceYksSkillAddAPIRequest{
    return &YunosTvpubadminDeviceYksSkillAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.yks.skill.add"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkillId Setter
// 技能id
func (r *YunosTvpubadminDeviceYksSkillAddAPIRequest) SetSkillId(_skillId int64) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetSkillId() int64 {
    return r._skillId
}
// BotId Setter
// 设备id
func (r *YunosTvpubadminDeviceYksSkillAddAPIRequest) SetBotId(_botId int64) error {
    r._botId = _botId
    r.Set("bot_id", _botId)
    return nil
}

// BotId Getter
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetBotId() int64 {
    return r._botId
}
// Name Setter
// 技能名称
func (r *YunosTvpubadminDeviceYksSkillAddAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetName() string {
    return r._name
}
// IconImageUrl Setter
// 图片地址
func (r *YunosTvpubadminDeviceYksSkillAddAPIRequest) SetIconImageUrl(_iconImageUrl string) error {
    r._iconImageUrl = _iconImageUrl
    r.Set("icon_image_url", _iconImageUrl)
    return nil
}

// IconImageUrl Getter
func (r YunosTvpubadminDeviceYksSkillAddAPIRequest) GetIconImageUrl() string {
    return r._iconImageUrl
}