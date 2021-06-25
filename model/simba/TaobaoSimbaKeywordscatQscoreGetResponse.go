package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词和类目出价的质量得分 APIResponse
taobao.simba.keywordscat.qscore.get

取得一个推广组的所有关键词和类目出价的质量得分列表
*/
type TaobaoSimbaKeywordscatQscoreGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordscatQscoreGetResponse `json:"taobao_simba_keywordscat_qscore_get_response,omitempty"`
}

type TaobaoSimbaKeywordscatQscoreGetResponse struct {

    // 类目出价和词的质量得分对象
    Qscore   *Qscore `json:"qscore,omitempty"`

}