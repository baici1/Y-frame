package configs

import "time"

type Casbin struct {
	ConfPosition          string        `yaml:"ConfPosition" json:"conf_position" mapstructure:"ConfPosition"`
	AutoLoadPolicySeconds time.Duration `yaml:"AutoLoadPolicySeconds" json:"auto_load_policy_seconds" mapstructure:"AutoLoadPolicySeconds"`
}
