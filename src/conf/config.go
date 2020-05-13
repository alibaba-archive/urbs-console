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

// ConfigTpl ...
type ConfigTpl struct {
	SrvAddr       string      `json:"addr" yaml:"addr"`
	TLS           TlsConfig   `json:"tls" yaml:"tls"`
	Logger        Logger      `json:"logger" yaml:"logger"`
	UrbsSetting   UrbsSetting `json:"urbs_setting" yaml:"urbs_setting"`
	Thrid         Thrid       `json:"thrid" yaml:"thrid"` // 三方接口
	CorsWhiteList []string    `json:"cors_white_list" yaml:"cors_white_list"`
	MySQL         SQL         `json:"mysql" yaml:"mysql"`
	SuperAdmins   []string    `json:"superAdmins" yaml:"superAdmins"`
	HIDKey        string      `json:"hid_key" yaml:"hid_key"`
	AuthKeys      []string    `json:"auth_keys" yaml:"auth_keys"`
}

// TlsConfig the config struct for creating tls.Config.
type TlsConfig struct {
	CertPath string `json:"cert_path" yaml:"cert_path"`
	KeyPath  string `json:"key_path" yaml:"key_path"`
}

// Logger logger config
type Logger struct {
	Level string `json:"level" yaml:"level"`
}

// UrbsSetting ...
type UrbsSetting struct {
	Addr string `json:"addr" yaml:"addr"`
	Key  string `json:"key" yaml:"key"`
}

// Thrid ...
type Thrid struct {
	Key string `json:"key" yaml:"key"`
	// 验证调用者身份接口
	UserAuth UserAuth `json:"user_auth" yaml:"user_auth"`
	// 群组成员接口
	GroupMember GroupMember `json:"group_member" yaml:"group_member"`
}

// UserAuth ...
type UserAuth struct {
	URL       string `json:"url" yaml:"url"`
	CookieKey string `json:"cookie_key" yaml:"cookie_key"`
}

// GroupMember ...
type GroupMember struct {
	URL string `json:"url" yaml:"url"`
}

// SQL ...
type SQL struct {
	Host         string `json:"host" yaml:"host"`
	User         string `json:"user" yaml:"user"`
	Password     string `json:"password" yaml:"password"`
	Database     string `json:"database" yaml:"database"`
	Parameters   string `json:"parameters" yaml:"parameters"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"`
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
