package configs

type SnowFlake struct {
	SnowFlakeMachineId int64 `yaml:"SnowFlakeMachineId" json:"snow_flake_machine_id" mapstructure:"SnowFlakeMachineId"`
}
