package conf

import (
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	p := &Config
	util.ReadConfig(p)
	if err := p.Validate(); err != nil {
		panic(err)
	}
}

// Logger logger config
type Logger struct {
	Level string `json:"level" yaml:"level"`
}

// ConfigTpl ...
type ConfigTpl struct {
	SrvAddr  string `json:"addr" yaml:"addr"`
	CertFile string `json:"cert_file" yaml:"cert_file"`
	KeyFile  string `json:"key_file" yaml:"key_file"`
	Logger   Logger `json:"logger" yaml:"logger"`
}

// Config ...
var Config ConfigTpl

// Validate 用于完成基本的配置验证和初始化工作。
func (c *ConfigTpl) Validate() error {
	if len(c.SrvAddr) == 0 {
		c.SrvAddr = ":8081"
	}
	logger.SetLevel(c.Logger.Level)
	logger.SetJSONLog()
	return nil
}
