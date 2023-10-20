package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordidschangedgetAPIResponse 获取修改的词ID API返回值
// taobao.simba.keywordids.changed.get
//
// 获取修改的词ID
type TaobaosimbakeywordidschangedgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbakeywordidschangedgetAPIResponseModel
}

// TaobaosimbakeywordidschangedgetAPIResponseModel is 获取修改的词ID 成功返回结果
type TaobaosimbakeywordidschangedgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordids_changed_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词的ID列表
	ChangedKeywordIds []int64 `json:"changed_keyword_ids,omitempty" xml:"changed_keyword_ids>int64,omitempty"`
}
