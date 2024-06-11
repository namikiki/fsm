package configs

import (
	"fsm/pkg/types"

	"github.com/BurntSushi/toml"
)

// NewConfig 读取本地配置文件
// todo 读取不到配置文件后，生成默认配置文件
func NewConfig() *types.Config {
	var cf types.Config
	if _, err := toml.DecodeFile("config.toml", &cf); err != nil {
		panic(err)
	}
	return &cf
}
