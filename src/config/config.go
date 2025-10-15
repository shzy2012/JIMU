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

// clickhouse
type ClickhouseCfg struct {
	Endpoints []string `toml:"endpoints"`
	Username  string   `toml:"username"`
	Password  string   `toml:"password"`
	Database  string   `toml:"database"`
}

// Postgres
type PostgresCfg struct {
	Endpoint string `toml:"endpoint"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

// MongoDB
type MongoDBCfg struct {
	URI      string `toml:"uri"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

// ES
type ESCfg struct {
	Username string   `toml:"username"`
	Password string   `toml:"password"`
	Index    string   `toml:"index"`
	URLs     []string `toml:"urls"`
}

// APIs
type APICfg struct {
}

// Files
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
	Server     ServerCfg     `toml:"server"`
	MongoDB    MongoDBCfg    `toml:"mongodb"`
	MySQL      MySQLCfg      `toml:"mysql"`
	Postgres   PostgresCfg   `toml:"postgres"`
	Elastic    ESCfg         `toml:"elastic"`
	API        APICfg        `toml:"api"`
	File       FileCfg       `toml:"file"`
	Clickhouse ClickhouseCfg `toml:"clickhouse"`
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
	Server     ServerCfg
	MongoDB    MongoDBCfg
	ES         ESCfg
	MySQL      MySQLCfg
	Postgres   PostgresCfg
	API        APICfg
	File       FileCfg
	Clickhouse ClickhouseCfg
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

	cfg := NewConfig(cfgPath)
	if cfg == nil {
		log.Fatalln("config parse failed")
	}

	Server = cfg.Server
	MongoDB = cfg.MongoDB
	MySQL = cfg.MySQL
	Clickhouse = cfg.Clickhouse
	Postgres = cfg.Postgres
	ES = cfg.Elastic
	API = cfg.API
	File = cfg.File
	log.Printf("[config]=> %s init ok\n", cfgPath)
}
