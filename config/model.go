package config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Log      LogConfig      `yaml:"log"`
	Auth     AuthConfig     `yaml:"auth"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	ConnectString string `yaml:"connect_string"`
	DBName        string `yaml:"database_name"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type AuthConfig struct {
	SecretKey             string `yaml:"secret_key"`
	AccessTokenExpireTime int    `yaml:"access_token_expire_minute"`
}
