package config

import (
	"github.com/spf13/viper"
)

var (
	Config *viper.Viper
)

func init() {
	Config = viper.New()
}
