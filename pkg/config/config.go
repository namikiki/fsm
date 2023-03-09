package config

import (
	"log"

	"fsm/pkg/types"

	"github.com/BurntSushi/toml"
)

func NewConfig() *types.Config {
	var cf types.Config
	if _, err := toml.DecodeFile("config.toml", &cf); err != nil {
		panic(err)
	}
	log.Println(cf)
	return &cf
}
