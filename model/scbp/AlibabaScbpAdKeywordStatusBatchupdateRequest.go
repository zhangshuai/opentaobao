package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量启动暂停推广词状态 API请求
alibaba.scbp.ad.keyword.status.batchupdate

批量启动暂停关键词推广状态
*/
type AlibabaScbpAdKeywordStatusBatchupdateRequest struct {
    model.Params
    // 系统自动生成
    keywordUpdateDtoList   []KeywordUpdateDto
}

// 初始化AlibabaScbpAdKeywordStatusBatchupdateRequest对象
func NewAlibabaScbpAdKeywordStatusBatchupdateRequest() *AlibabaScbpAdKeywordStatusBatchupdateRequest{
    return &AlibabaScbpAdKeywordStatusBatchupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordStatusBatchupdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.status.batchupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordStatusBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordUpdateDtoList Setter
// 系统自动生成
func (r *AlibabaScbpAdKeywordStatusBatchupdateRequest) SetKeywordUpdateDtoList(keywordUpdateDtoList []KeywordUpdateDto) error {
    r.keywordUpdateDtoList = keywordUpdateDtoList
    r.Set("keyword_update_dto_list", keywordUpdateDtoList)
    return nil
}

// KeywordUpdateDtoList Getter
func (r AlibabaScbpAdKeywordStatusBatchupdateRequest) GetKeywordUpdateDtoList() []KeywordUpdateDto {
    return r.keywordUpdateDtoList
}
