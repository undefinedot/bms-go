package config

type Zap struct {
	Release string `mapstructure:"release"`
	Level   string `mapstructure:"level"`
	Format  string `mapstructure:"format"`
	Dir     string `mapstructure:"dir"`
	Name    string `mapstructure:"name"`
}
