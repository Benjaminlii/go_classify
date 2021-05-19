package config

// 基础配置类
type BasicConfig struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

// Mysql db相关配置
type Mysql struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DB        string `yaml:"db"`
	UserName  string `yaml:"username"`
	PassWorld string `yaml:"password"`
}

// Redis redis相关配置
type Redis struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DB        string `yaml:"db"`
	PassWorld string `yaml:"password"`
}
