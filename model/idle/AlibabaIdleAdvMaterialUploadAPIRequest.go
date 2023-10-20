package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleadvmaterialuploadAPIRequest 闲鱼用户增长素材中心素材上传接口 API请求
// alibaba.idle.adv.material.upload
//
// 闲鱼用户增长素材中心素材上传接口
type AlibabaidleadvmaterialuploadAPIRequest struct {
	model.Params
	// 素材上传参数
	_uploadTopParam *IdleAdvMaterialUploadTopParam
}

// NewAlibabaidleadvmaterialuploadRequest 初始化AlibabaidleadvmaterialuploadAPIRequest对象
func NewAlibabaidleadvmaterialuploadRequest() *AlibabaidleadvmaterialuploadAPIRequest {
	return &AlibabaidleadvmaterialuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleadvmaterialuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.adv.material.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleadvmaterialuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleadvmaterialuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUploadTopParam is UploadTopParam Setter
// 素材上传参数
func (r *AlibabaidleadvmaterialuploadAPIRequest) SetUploadTopParam(_uploadTopParam *IdleAdvMaterialUploadTopParam) error {
	r._uploadTopParam = _uploadTopParam
	r.Set("upload_top_param", _uploadTopParam)
	return nil
}

// GetUploadTopParam UploadTopParam Getter
func (r AlibabaidleadvmaterialuploadAPIRequest) GetUploadTopParam() *IdleAdvMaterialUploadTopParam {
	return r._uploadTopParam
}