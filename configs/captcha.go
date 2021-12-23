package configs

import "time"

type Captcha struct {
	Expiration time.Duration `yaml:"Expiration" json:"expiration,omitempty" mapstructure:"Expiration"`
	StdWidth   int           `yaml:"StdWidth" json:"std_width,omitempty" mapstructure:"StdWidth"`
	StdHeight  int           `yaml:"StdHeight" json:"std_height,omitempty" mapstructure:"StdHeight"`
	Lang       string        `yaml:"Lang" json:"lang,omitempty" mapstructure:"Lang"`
	Length     int           `yaml:"Length" json:"length,omitempty" mapstructure:"Length"`
	CollectNum int           `yaml:"CollectNum" json:"collect_num,omitempty" mapstructure:"CollectNum"`
}
