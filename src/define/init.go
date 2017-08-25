package define

import (
    "flag"
    "github.com/go-ini/ini"
    "log"
)

var (
    GatewayIP   string
    GatewayPort string


    Environment string
    Config      string
)

func init () {
    flag.StringVar(&Environment, "env","dev","server run environment")
    flag.StringVar(&Config, "conf","","config path")

    flag.Parse()

    conf, err := ini.Load(Config)
    if err != nil {
        log.Panic("init config error")
    }

    GatewayIP = conf.Section(Environment).Key("GatewayIP").String()
    GatewayPort = conf.Section(Environment).Key("GatewayPort").String()
}