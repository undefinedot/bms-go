package config

type JWT struct {
	SecretKey  string `mapstructure:"secret_key"`
	Issuer     string `mapstructure:"issuer"`
	ExpireTime int64  `mapstructure:"expire_time"`
}
