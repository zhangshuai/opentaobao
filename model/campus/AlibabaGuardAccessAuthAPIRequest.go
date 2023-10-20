package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGuardAccessAuthAPIRequest 鉴权 API请求
// alibaba.guard.access.auth
//
// 刷卡鉴权
type AlibabaGuardAccessAuthAPIRequest struct {
	model.Params
	// 鉴权结果DTO
	_identifyAuthDto *IdentifyAuthDto
}

// NewAlibabaGuardAccessAuthRequest 初始化AlibabaGuardAccessAuthAPIRequest对象
func NewAlibabaGuardAccessAuthRequest() *AlibabaGuardAccessAuthAPIRequest {
	return &AlibabaGuardAccessAuthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGuardAccessAuthAPIRequest) GetApiMethodName() string {
	return "alibaba.guard.access.auth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGuardAccessAuthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaGuardAccessAuthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifyAuthDto is IdentifyAuthDto Setter
// 鉴权结果DTO
func (r *AlibabaGuardAccessAuthAPIRequest) SetIdentifyAuthDto(_identifyAuthDto *IdentifyAuthDto) error {
	r._identifyAuthDto = _identifyAuthDto
	r.Set("identify_auth_dto", _identifyAuthDto)
	return nil
}

// GetIdentifyAuthDto IdentifyAuthDto Getter
func (r AlibabaGuardAccessAuthAPIRequest) GetIdentifyAuthDto() *IdentifyAuthDto {
	return r._identifyAuthDto
}
