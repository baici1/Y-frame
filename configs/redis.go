package configs

type Redis struct {
	Port               int    `yaml:"Port" json:"port,omitempty" mapstructure:"Port"`
	MaxIdle            int    `yaml:"MaxIdle" json:"max_idle,omitempty" mapstructure:"MaxIdle"`
	MaxActive          int    `yaml:"MaxActive" json:"max_active,omitempty" mapstructure:"MaxActive"`
	ConnFailRetryTimes int    `yaml:"ConnFailRetryTimes" json:"conn_fail_retry_times,omitempty" mapstructure:"ConnFailRetryTimes"`
	Host               string `yaml:"Host" json:"host,omitempty" mapstructure:"Host"`
	IdleTimeout        int    `yaml:"IdleTimeout" json:"idle_timeout,omitempty" mapstructure:"IdleTimeout"`
	IndexDb            int    `yaml:"IndexDb" json:"index_db,omitempty" mapstructure:"IndexDb"`
	ReConnectInterval  int    `yaml:"ReConnectInterval" json:"re_connect_interval,omitempty" mapstructure:"ReConnectInterval"`
	Auth               string `yaml:"Auth" json:"auth,omitempty" mapstructure:"Auth"`
}
