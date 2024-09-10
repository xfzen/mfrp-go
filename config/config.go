package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Frp struct {
		Conf string
	}
}
