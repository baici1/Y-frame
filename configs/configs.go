package configs

type Server struct {
	Zaps      Zaps              `json:"zaps" yaml:"Zaps" mapstructure:"Zaps"`
	Gormv2    Gormv2            `json:"gormv_2" yaml:"Gormv2" mapstructure:"Gormv2"`
	Redis     Redis             `json:"redis" yaml:"Redis" mapstructure:"Redis"`
	System    HTTPServer        `json:"system" yaml:"HTTPServer" mapstructure:"HTTPServer"`
	Token     Token             `json:"token" yaml:"Token" mapstructure:"Token"`
	File      FileUploadSetting `json:"file" yaml:"FileUploadSetting" mapstructure:"FileUploadSetting"`
	Captcha   Captcha           `json:"captcha" yaml:"Captcha" mapstructure:"Captcha"`
	SnowFlake SnowFlake         `json:"SnowFlake" yaml:"SnowFlake" mapstructure:"SnowFlake"`
}
