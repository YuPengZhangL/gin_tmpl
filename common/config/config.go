package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var conf *Config

type Config struct {
	App   *App   `yaml:"app"`
	MySQL *MySQL `yaml:"mysql"`
	Redis *Redis `yaml:"redis"`
}

type App struct {
	Name       string `yaml:"name"`
	Release    bool   `yaml:"release"`
	Port       int    `yaml:"port"`
	LogLevel   string `yaml:"log_level"`
	LogPath    string `yaml:"log_path"`
	JWTSignKey string `yaml:"jwt_sign_key"`
	JWTIssuer  string `yaml:"jwt_issuer"`
}

type MySQL struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Redis struct {
	RedisHost   string `yaml:"host"`
	RedisPort   string `yaml:"port"`
	RedisPasswd string `yaml:"password"`
}

func InitConfig(file string) error {
	conf = &Config{}
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(f, conf)
	return err
}

func InitMockConfig() {
	conf = &Config{
		App: &App{
			Name:       "gg_web_tmpl_test",
			Release:    true,
			Port:       8081,
			LogLevel:   "debug",
			LogPath:    "./logs",
			JWTSignKey: "test",
			JWTIssuer:  "test",
		},
		MySQL: &MySQL{
			User:     "test",
			Password: "test",
			DB:       "test",
			Host:     "127.0.0.1",
			Port:     3306,
		},
		Redis: &Redis{
			RedisHost: "127.0.0.1",
			RedisPort: "6379",
		},
	}
}

func GetConf() *Config {
	return conf
}
