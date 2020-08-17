package main

import (
	"github.com/Zhousiru/libone/internal/api"
	"github.com/Zhousiru/libone/internal/config"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	api.Run(viper.GetString("addr"))

	return
}
