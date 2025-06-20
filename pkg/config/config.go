package config

import (
	"io"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// 全局配置实例
var (
	cfg  *AppConfig
	once sync.Once
)

// AppConfig 顶层配置结构
type AppConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"mysql"`
	Redis    RedisConfig    `yaml:"redis"`
	Mongo    MongoConfig    `yaml:"mongodb"`
	Log      LogConfig      `yaml:"log"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host         string `yaml:"host"` // 监听地址
	Port         int    `yaml:"port"`
	Mode         string `yaml:"mode"` // debug/release
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	//DSN             string `yaml:"dsn"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	DbName          string `yaml:"db_name"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Charset         string `yaml:"charset"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

// ... 其他配置结构
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type MongoConfig struct {
	Uri     string `yaml:"uri"`
	DBName  string `yaml:"db_name"`
	TimeOut int64  `yaml:"time_out"`
}

type LogConfig struct {
}

// Load 加载配置 (线程安全)
func Load(configPath string) error {
	var err error
	once.Do(func() {
		cfg, err = loadFromYAML(configPath)
	})
	return err
}

// Get 获取配置实例
func Get() *AppConfig {
	return cfg
}

// 私有加载方法
func loadFromYAML(path string) (*AppConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
