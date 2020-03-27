package conf

import (
	"github.com/teambition/gear/logging"
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
	SrvAddr     string      `json:"addr" yaml:"addr"`
	TLS         TlsConfig   `json:"tls" yaml:"tls"`
	Logger      Logger      `json:"logger" yaml:"logger"`
	UrbsSetting UrbsSetting `json:"urbs_setting" yaml:"urbs_setting"`
	// 验证调用者身份
	UserAuth UserAuth `json:"user_auth" yaml:"user_auth"`
	// 拉取群组成员
	GroupMember GroupMember `json:"group_member" yaml:"group_member"`
}

// TlsConfig the config struct for creating tls.Config.
type TlsConfig struct {
	CertPath string `json:"cert_path" yaml:"cert_path"`
	KeyPath  string `json:"key_path" yaml:"key_path"`
}

// UserAuth ...
type UserAuth struct {
	Keys          []string      `json:"keys" yaml:"keys"`
	UserAuthThrid UserAuthThrid `json:"thrid" yaml:"thrid"`
}

// UserAuthThrid ...
type UserAuthThrid struct {
	URL     string                 `json:"url" yaml:"url"`
	TokenKV map[string]interface{} `json:"token_kv" yaml:"token_kv"`
	BodyKK  map[string]string      `json:"body_kk" yaml:"body_kk"`
	From    string                 `json:"from" yaml:"from"`
}

// GroupMember ...
type GroupMember struct {
	URL     string                 `json:"url" yaml:"url"`
	Key     string                 `json:"keys" yaml:"key"`
	TokenKV map[string]interface{} `json:"token_kv" yaml:"token_kv"`
	BodyKK  map[string]string      `json:"body_kk" yaml:"body_kk"`
}

// UrbsSetting ...
type UrbsSetting struct {
	Addr     string   `json:"addr" yaml:"addr"`
	AuthKeys []string `json:"auth_keys" yaml:"auth_keys"`
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

	logging.Default().SetJSONLog().SetLevel(logging.Level(logger.Level()))
	return nil
}
