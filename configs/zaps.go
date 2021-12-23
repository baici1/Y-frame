package configs

type Zaps struct {
	AppDebug          bool   `yaml:"AppDebug" json:"AppDebug,omitempty" mapstructure:"AppDebug"`
	GinLogName        string `yaml:"GinLogName" json:"gin_log_name,omitempty" mapstructure:"GinLogName"`
	GoSkeletonLogName string `yaml:"GoSkeletonLogName" json:"go_skeleton_log_name,omitempty" mapstructure:"GoSkeletonLogName"`
	TextFormat        string `yaml:"TextFormat" json:"text_format,omitempty" mapstructure:"TextFormat"`
	TimePrecision     string `yaml:"TimePrecision" json:"time_precision,omitempty" mapstructure:"TimePrecision"`
	MaxSize           int    `yaml:"MaxSize" json:"max_size,omitempty" mapstructure:"MaxSize"`
	MaxBackups        int    `yaml:"MaxBackups" json:"max_backups,omitempty" mapstructure:"MaxBackups"`
	MaxAge            int    `yaml:"MaxAge" json:"max_age,omitempty" mapstructure:"MaxAge"`
	Compress          bool   `yaml:"Compress" json:"compress,omitempty" mapstructure:"Compress"`
}
