package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交图片翻译任务 API请求
alibaba.seaking.imagetranslate.submit

提交图片翻译任务
*/
type AlibabaSeakingImagetranslateSubmitRequest struct {
    model.Params
    // token来源站点
    tokenFrom   string
    // 子任务列表
    imageTranslateDetailList   []ImageTranslateDetailDto
    // 用户token
    token   string
}

// 初始化AlibabaSeakingImagetranslateSubmitRequest对象
func NewAlibabaSeakingImagetranslateSubmitRequest() *AlibabaSeakingImagetranslateSubmitRequest{
    return &AlibabaSeakingImagetranslateSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagetranslateSubmitRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagetranslate.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagetranslateSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingImagetranslateSubmitRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingImagetranslateSubmitRequest) GetTokenFrom() string {
    return r.tokenFrom
}
// ImageTranslateDetailList Setter
// 子任务列表
func (r *AlibabaSeakingImagetranslateSubmitRequest) SetImageTranslateDetailList(imageTranslateDetailList []ImageTranslateDetailDto) error {
    r.imageTranslateDetailList = imageTranslateDetailList
    r.Set("image_translate_detail_list", imageTranslateDetailList)
    return nil
}

// ImageTranslateDetailList Getter
func (r AlibabaSeakingImagetranslateSubmitRequest) GetImageTranslateDetailList() []ImageTranslateDetailDto {
    return r.imageTranslateDetailList
}
// Token Setter
// 用户token
func (r *AlibabaSeakingImagetranslateSubmitRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaSeakingImagetranslateSubmitRequest) GetToken() string {
    return r.token
}