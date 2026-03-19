package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Log      LogConfig      `yaml:"log"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	JWT      JWTConfig      `yaml:"jwt"`
	Email    EmailConfig    `yaml:"email"`
	Upload   UploadConfig   `yaml:"upload"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DBName       string `yaml:"dbname"`
	Charset      string `yaml:"charset"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConfig struct {
	Secret            string `yaml:"secret"`
	ExpireHours       int    `yaml:"expire_hours"`
	RememberExpireDays int   `yaml:"remember_expire_days"`
}

type EmailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

type UploadConfig struct {
	MaxFileSize       int64    `yaml:"max_file_size"`
	ChunkSize         int64    `yaml:"chunk_size"`
	UploadDir         string   `yaml:"upload_dir"`
	AllowedExtensions []string `yaml:"allowed_extensions"`
}

func LoadConfig(path string) (*Config, error) {
	var err error
	once.Do(func() {
		var data []byte
		data, err = os.ReadFile(path)
		if err != nil {
			return
		}
		cfg = &Config{}
		err = yaml.Unmarshal(data, cfg)
	})
	return cfg, err
}

func GetConfig() *Config {
	return cfg
}
