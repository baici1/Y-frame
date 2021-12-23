package configs

type HTTPServer struct {
	API              string `json:"Api" yaml:"api" mapstructure:"API"`
	Web              string `json:"Web" yaml:"web" mapstructure:"Web"`
	AllowCrossDomain bool   `json:"AllowCrossDomain" yaml:"allow_cross_domain" mapstructure:"AllowCrossDomain"`
}
