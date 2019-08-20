package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// all configuration keys
const (
	StartX = "start.x"
	StartY = "start.y"
)

func init() {
	pflag.String(StartX, "", "starting x axis point")
	pflag.String(StartY, "", "starting y axis point")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	_ = viper.BindEnv(StartX, "START_X")
	_ = viper.BindEnv(StartY, "START_Y")

	viper.SetDefault(StartX, 1)
	viper.SetDefault(StartY, 1)
}
