package tbk

import (
	"sync"
)

// MiniProgramDto 结构体
type MiniProgramDto struct {
	// 小程序APPID
	MiniProgramAppid string `json:"mini_program_appid,omitempty" xml:"mini_program_appid,omitempty"`
	// 小程序路径
	MiniProgramPath string `json:"mini_program_path,omitempty" xml:"mini_program_path,omitempty"`
	// 小程序码url地址
	MiniProgramQrcodeUrl string `json:"mini_program_qrcode_url,omitempty" xml:"mini_program_qrcode_url,omitempty"`
}

var poolMiniProgramDto = sync.Pool{
	New: func() any {
		return new(MiniProgramDto)
	},
}

// GetMiniProgramDto() 从对象池中获取MiniProgramDto
func GetMiniProgramDto() *MiniProgramDto {
	return poolMiniProgramDto.Get().(*MiniProgramDto)
}

// ReleaseMiniProgramDto 释放MiniProgramDto
func ReleaseMiniProgramDto(v *MiniProgramDto) {
	v.MiniProgramAppid = ""
	v.MiniProgramPath = ""
	v.MiniProgramQrcodeUrl = ""
	poolMiniProgramDto.Put(v)
}
