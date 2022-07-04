package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	DefaultContextTimeout time.Duration
	OpenTracing           struct {
		ServiceName string
		AgentHost   string
		AgentPort   string
	}
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type ApiClientSettingS struct {
	AccessTokenIdentity  string
	AccessTokenValidity  bool
	AccessTokenExpire    int
	RefreshTokenValidity bool
	RefreshTokenExpire   int
	HttpSignValidity     bool
	HttpSignExpire       int64
	HttpSignAccount      struct {
		Key        string
		Secret     string
		SignName   string
		ExpireName string
	}
}

type WechatSettingS struct {
	Work struct {
		WebHook struct {
			EndPoint string
			Key      string
		}
	}
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
