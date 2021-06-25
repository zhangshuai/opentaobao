package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词按地域进行细分的数据 APIResponse
taobao.simba.insight.wordsareadata.get

获取关键词按地域细分的详细数据，目前地域只能细化到省级别，返回的结果中包括市，是为了方便以后扩展，目前结果中市的值等于省。
*/
type TaobaoSimbaInsightWordsareadataGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaInsightWordsareadataGetResponse `json:"taobao_simba_insight_wordsareadata_get_response,omitempty"`
}

type TaobaoSimbaInsightWordsareadataGetResponse struct {

    // 地域细分数据
    WordAreadataList   []InsightWordsAreaDistributeDataDTO `json:"word_areadata_list,omitempty"`

}