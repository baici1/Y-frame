package configs

type Token struct {
	JwtTokenRefreshExpireAt int64  `yaml:"JwtTokenRefreshExpireAt" json:"jwt_token_refresh_expire_at,omitempty" mapstructure:"JwtTokenRefreshExpireAt"`
	BindContextKeyName      string `yaml:"BindContextKeyName" json:"bind_context_key_name,omitempty" mapstructure:"BindContextKeyName"`
	JwtTokenSignKey         string `yaml:"JwtTokenSignKey" json:"jwt_token_sign_key,omitempty" mapstructure:"JwtTokenSignKey"`
	JwtTokenOnlineUsers     int    `yaml:"JwtTokenOnlineUsers" json:"jwt_token_online_users,omitempty" mapstructure:"JwtTokenOnlineUsers"`
	JwtTokenCreatedExpireAt int64  `yaml:"JwtTokenCreatedExpireAt" json:"jwt_token_created_expire_at,omitempty" mapstructure:"JwtTokenCreatedExpireAt"`
	JwtTokenRefreshAllowSec int    `yaml:"JwtTokenRefreshAllowSec" json:"jwt_token_refresh_allow_sec,omitempty" mapstructure:"JwtTokenRefreshAllowSec"`
}
