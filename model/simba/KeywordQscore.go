package simba

import (
	"sync"
)

// KeywordQscore 结构体
type KeywordQscore struct {
	// 主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 质量总分
	Qscore string `json:"qscore,omitempty" xml:"qscore,omitempty"`
	// 基础分
	CustScore string `json:"cust_score,omitempty" xml:"cust_score,omitempty"`
	// 创意得分
	CreativeScore string `json:"creative_score,omitempty" xml:"creative_score,omitempty"`
	// 相关性
	ReleScore string `json:"rele_score,omitempty" xml:"rele_score,omitempty"`
	// 买家体验分
	CvrScore string `json:"cvr_score,omitempty" xml:"cvr_score,omitempty"`
	// 推广计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 推广组id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 关键词id
	KeywordId int64 `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
}

var poolKeywordQscore = sync.Pool{
	New: func() any {
		return new(KeywordQscore)
	},
}

// GetKeywordQscore() 从对象池中获取KeywordQscore
func GetKeywordQscore() *KeywordQscore {
	return poolKeywordQscore.Get().(*KeywordQscore)
}

// ReleaseKeywordQscore 释放KeywordQscore
func ReleaseKeywordQscore(v *KeywordQscore) {
	v.Nick = ""
	v.Word = ""
	v.Qscore = ""
	v.CustScore = ""
	v.CreativeScore = ""
	v.ReleScore = ""
	v.CvrScore = ""
	v.CampaignId = 0
	v.AdgroupId = 0
	v.KeywordId = 0
	poolKeywordQscore.Put(v)
}
