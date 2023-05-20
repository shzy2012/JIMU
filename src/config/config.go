package config

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/BurntSushi/toml"
)

type MongoDB struct {
	URI      string `yaml:"uri"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// 服务器配置
type ServerCfg struct {
	RunMode      string        `toml:"run_mode"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

// 全局配置文件
type Config struct {
	Server ServerCfg `toml:"server"`
	Mongo  MongoDB   `toml:"mongodb"`
}

// NewConfig 初始化配置文件
func NewConfig(path string) *Config {
	cfg := Config{}
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		log.Fatalf("[config load]=>%s\n", err)
	}
	return &cfg
}

var (
	Serve ServerCfg
	Mongo MongoDB
)

func init() {

	//检查文件策略:
	// 1.先从当前目录检查config.toml是否存在,
	// 2.然后再是否有config目录,如果存在检查config目录下config.toml是否存在

	cfgPath := ""
	paths := []string{
		"config-dev.toml",
		"config.toml",
		"../config-dev.toml",
		"../config.toml",
		path.Join("config", "config-dev.toml"),
		path.Join("config", "config.toml"),
		path.Join("src", "config", "config-dev.toml"),
		path.Join("src", "config", "config.toml"),
	}
	//检查目录是否存在
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			cfgPath = path
			break
		}
	}
	if cfgPath == "" {
		log.Fatalln("not found config")
	}

	log.Printf("[cfgpath]:%s\n", cfgPath)
	cfg := NewConfig(cfgPath)
	if cfg == nil {
		log.Fatalln("config parse failed")
	}

	Serve = cfg.Server
	Mongo = cfg.Mongo
	log.Printf("config from %s init ok\n", cfgPath)
}
