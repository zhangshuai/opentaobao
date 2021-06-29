package tmallcarenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车EPC版本压缩库新增接口 API请求
tmall.carcenter.vehicle.version.insert

汽车EPC版本压缩库新增接口
*/
type TmallCarcenterVehicleVersionInsertRequest struct {
    model.Params
    // 版本压缩库入参
    dto   *VersionVehicleInfoOriginalDto
}

// 初始化TmallCarcenterVehicleVersionInsertRequest对象
func NewTmallCarcenterVehicleVersionInsertRequest() *TmallCarcenterVehicleVersionInsertRequest{
    return &TmallCarcenterVehicleVersionInsertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleVersionInsertRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicle.version.insert"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleVersionInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Dto Setter
// 版本压缩库入参
func (r *TmallCarcenterVehicleVersionInsertRequest) SetDto(dto *VersionVehicleInfoOriginalDto) error {
    r.dto = dto
    r.Set("dto", dto)
    return nil
}

// Dto Getter
func (r TmallCarcenterVehicleVersionInsertRequest) GetDto() *VersionVehicleInfoOriginalDto {
    return r.dto
}