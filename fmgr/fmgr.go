package fmgr

import (
	"context"
	"fmt"
	"os"

	mconfig "mfrp/config"
	"mfrp/pkg/utils"

	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	"github.com/fatedier/frp/server"
	"github.com/zeromicro/go-zero/core/logx"
)

var strictConfigMode bool

func CreateFrpsService(c mconfig.Config) (*server.Service, error) {
	logx.Debugf("CreateService Frp Conf: %v", utils.PrettyJson(c.Frp))

	cfg, isLegacyFormat, err := config.LoadServerConfig(c.Frp.Conf, strictConfigMode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isLegacyFormat {
		fmt.Printf("WARNING: ini format is deprecated and the support will be removed in the future, " +
			"please use yaml/json/toml format instead!\n")
	}

	warning, err := validation.ValidateServerConfig(cfg)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("frps: the configuration file %s syntax is ok\n", c.Frp.Conf)

	svr, err := server.NewService(cfg)
	if err != nil {
		return nil, err
	}

	logx.Infof("frps started successfully")

	go svr.Run(context.Background())

	return svr, nil
}
