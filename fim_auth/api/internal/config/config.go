package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DBSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
}
