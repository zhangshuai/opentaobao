package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateUpdateappAPIRequest 更新实例化应用 API请求
// taobao.miniapp.template.updateapp
//
// 商家应用c端模板实例化小程序更新，生成新的版本，但不会自动上线新版本
type TaobaoMiniappTemplateUpdateappAPIRequest struct {
	model.Params
	// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
	_clients []string
	// 应用id，如果应用在对应端上已有的最新版本所使用的模板id,模板version和extjson，与本次入参一致，则认为不需要更新，返回已有的版本。
	_appId string
	// 扩展信息。线上版本使用的template_id与传入的template_id一致时，可不填。
	_extJson string
	// 模板id
	_templateId string
	// 模板版本
	_templateVersion string
	// Logo更新，1月5次
	_icon string
	// 描述更新，1年5次
	_desc string
	// 简称更新，1年5次
	_alias string
}

// NewTaobaoMiniappTemplateUpdateappRequest 初始化TaobaoMiniappTemplateUpdateappAPIRequest对象
func NewTaobaoMiniappTemplateUpdateappRequest() *TaobaoMiniappTemplateUpdateappAPIRequest {
	return &TaobaoMiniappTemplateUpdateappAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) Reset() {
	r._clients = r._clients[:0]
	r._appId = ""
	r._extJson = ""
	r._templateId = ""
	r._templateVersion = ""
	r._icon = ""
	r._desc = ""
	r._alias = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.updateapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClients is Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) SetClients(_clients []string) error {
	r._clients = _clients
	r.Set("clients", _clients)
	return nil
}

// GetClients Clients Getter
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetClients() []string {
	return r._clients
}

// SetAppId is AppId Setter
// 应用id，如果应用在对应端上已有的最新版本所使用的模板id,模板version和extjson，与本次入参一致，则认为不需要更新，返回已有的版本。
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetAppId() string {
	return r._appId
}

// SetExtJson is ExtJson Setter
// 扩展信息。线上版本使用的template_id与传入的template_id一致时，可不填。
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetTemplateVersion is TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) SetTemplateVersion(_templateVersion string) error {
	r._templateVersion = _templateVersion
	r.Set("template_version", _templateVersion)
	return nil
}

// GetTemplateVersion TemplateVersion Getter
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetTemplateVersion() string {
	return r._templateVersion
}

// SetIcon is Icon Setter
// Logo更新，1月5次
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) SetIcon(_icon string) error {
	r._icon = _icon
	r.Set("icon", _icon)
	return nil
}

// GetIcon Icon Getter
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetIcon() string {
	return r._icon
}

// SetDesc is Desc Setter
// 描述更新，1年5次
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetDesc() string {
	return r._desc
}

// SetAlias is Alias Setter
// 简称更新，1年5次
func (r *TaobaoMiniappTemplateUpdateappAPIRequest) SetAlias(_alias string) error {
	r._alias = _alias
	r.Set("alias", _alias)
	return nil
}

// GetAlias Alias Getter
func (r TaobaoMiniappTemplateUpdateappAPIRequest) GetAlias() string {
	return r._alias
}

var poolTaobaoMiniappTemplateUpdateappAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappTemplateUpdateappRequest()
	},
}

// GetTaobaoMiniappTemplateUpdateappRequest 从 sync.Pool 获取 TaobaoMiniappTemplateUpdateappAPIRequest
func GetTaobaoMiniappTemplateUpdateappAPIRequest() *TaobaoMiniappTemplateUpdateappAPIRequest {
	return poolTaobaoMiniappTemplateUpdateappAPIRequest.Get().(*TaobaoMiniappTemplateUpdateappAPIRequest)
}

// ReleaseTaobaoMiniappTemplateUpdateappAPIRequest 将 TaobaoMiniappTemplateUpdateappAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappTemplateUpdateappAPIRequest(v *TaobaoMiniappTemplateUpdateappAPIRequest) {
	v.Reset()
	poolTaobaoMiniappTemplateUpdateappAPIRequest.Put(v)
}
