package tbk

// MiniProgramDto 结构体
type MiniProgramDto struct {
	// 小程序APPID
	Miniprogramappid string `json:"mini_program_appid,omitempty" xml:"mini_program_appid,omitempty"`
	// 小程序路径
	Miniprogrampath string `json:"mini_program_path,omitempty" xml:"mini_program_path,omitempty"`
	// 小程序码url地址
	Miniprogramqrcodeurl string `json:"mini_program_qrcode_url,omitempty" xml:"mini_program_qrcode_url,omitempty"`
}
