package config

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/BurntSushi/toml"
)

// MySQL
type MySQLCfg struct {
	Endpoint string `toml:"endpoint"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type MongoDBCfg struct {
	URI      string `toml:"uri"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type ESCfg struct {
	Username string   `toml:"username"`
	Password string   `toml:"password"`
	Index    string   `toml:"index"`
	URLs     []string `toml:"urls"`
}

type APICfg struct {
}

type FileCfg struct {
}

// 服务器配置
type ServerCfg struct {
	RunMode      string        `toml:"run_mode"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

// 全局配置文件
type Config struct {
	Server  ServerCfg  `toml:"server"`
	MongoDB MongoDBCfg `toml:"mongodb"`
	MySQL   MySQLCfg   `toml:"mysql"`
	Elastic ESCfg      `toml:"elastic"`
	API     APICfg     `toml:"api"`
	File    FileCfg    `toml:"file"`
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
	Server  ServerCfg
	MongoDB MongoDBCfg
	ES      ESCfg
	API     APICfg
	File    FileCfg
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

	Server = cfg.Server
	MongoDB = cfg.MongoDB
	ES = cfg.Elastic
	API = cfg.API
	File = cfg.File
	log.Printf("config from %s init ok\n", cfgPath)
}
