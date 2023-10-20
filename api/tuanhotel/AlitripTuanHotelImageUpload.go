package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelImageUpload 图片上传接口
// alitrip.tuan.hotel.image.upload
//
// 用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址
func AlitripTuanHotelImageUpload(clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelImageUploadAPIRequest, session string) (*tuanhotel.AlitripTuanHotelImageUploadAPIResponse, error) {
	var resp tuanhotel.AlitripTuanHotelImageUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
