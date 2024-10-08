package mtopopen

import (
	"sync"
)

// MsgSendRequest 结构体
type MsgSendRequest struct {
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 扩展参数
	ExtParams string `json:"ext_params,omitempty" xml:"ext_params,omitempty"`
	// 发生时间
	OccurTime int64 `json:"occur_time,omitempty" xml:"occur_time,omitempty"`
	// 类型（1-RETENTION-滞留即将收费、2-TAKEOUT-快递员取出）
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolMsgSendRequest = sync.Pool{
	New: func() any {
		return new(MsgSendRequest)
	},
}

// GetMsgSendRequest() 从对象池中获取MsgSendRequest
func GetMsgSendRequest() *MsgSendRequest {
	return poolMsgSendRequest.Get().(*MsgSendRequest)
}

// ReleaseMsgSendRequest 释放MsgSendRequest
func ReleaseMsgSendRequest(v *MsgSendRequest) {
	v.CpCode = ""
	v.MailNo = ""
	v.ExtParams = ""
	v.OccurTime = 0
	v.Type = 0
	poolMsgSendRequest.Put(v)
}
