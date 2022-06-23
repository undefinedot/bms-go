package config

type Mysql struct {
	Host         string `mapstructure:"host"`
	DbName       string `mapstructure:"db_name"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Post         string `mapstructure:"post"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

// Dns 拼接MySQL的dns
func (m *Mysql) Dns() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Post + ")/" + m.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
