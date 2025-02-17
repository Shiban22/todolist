package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
	Redis struct {
		Host string
		Type string
		Pass string
		Key  string
	}
}
