package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴权 API请求
alibaba.guard.access.auth

刷卡鉴权
*/
type AlibabaGuardAccessAuthRequest struct {
    model.Params
    // 请求
    paramIdentifyAuthDTO   *IdentifyAuthDto
}

// 初始化AlibabaGuardAccessAuthRequest对象
func NewAlibabaGuardAccessAuthRequest() *AlibabaGuardAccessAuthRequest{
    return &AlibabaGuardAccessAuthRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGuardAccessAuthRequest) GetApiMethodName() string {
    return "alibaba.guard.access.auth"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGuardAccessAuthRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamIdentifyAuthDTO Setter
// 请求
func (r *AlibabaGuardAccessAuthRequest) SetParamIdentifyAuthDTO(paramIdentifyAuthDTO *IdentifyAuthDto) error {
    r.paramIdentifyAuthDTO = paramIdentifyAuthDTO
    r.Set("param_identify_auth_d_t_o", paramIdentifyAuthDTO)
    return nil
}

// ParamIdentifyAuthDTO Getter
func (r AlibabaGuardAccessAuthRequest) GetParamIdentifyAuthDTO() *IdentifyAuthDto {
    return r.paramIdentifyAuthDTO
}