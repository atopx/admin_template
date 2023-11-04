package config

type Config struct {
	Server *ServerConfig `yaml:"server"`
	Redis  *RedisConfig  `yaml:"redis"`
	Mysql  *MysqlConfig  `yaml:"mysql"`
}

type ServerConfig struct {
	Mode     string
	Host     string
	Port     int
	Loglevel string
}

type RedisConfig struct {
	Host     string
	Port     int
	Db       int
	Password string
}

type MysqlConfig struct {
	User  string
	Pass  string
	Host  string
	Port  int
	Name  string
	Debug bool
}
