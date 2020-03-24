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
	SrvAddr     string     `json:"addr" yaml:"addr"`
	TLS         TlsConfig  `json:"tls" yaml:"tls"`
	Logger      Logger     `json:"logger" yaml:"logger"`
	UrbsSetting UrlSetting `json:"urbs_setting" yaml:"urbs_setting"`
}

// TlsConfig the config struct for creating tls.Config.
type TlsConfig struct {
	CertPath string `json:"cert_path" yaml:"cert_path"`
	KeyPath  string `json:"key_path" yaml:"key_path"`
}

// UrlSetting ...
type UrlSetting struct {
	Addr string `json:"addr" yaml:"addr"`
}

// Config ...
var Config ConfigTpl

// Validate 用于完成基本的配置验证和初始化工作。
func (c *ConfigTpl) Validate() error {
	if len(c.SrvAddr) == 0 {
		c.SrvAddr = ":8080"
	}
	logger.SetLevel(c.Logger.Level)
	logger.SetJSONLog()
	return nil
}
