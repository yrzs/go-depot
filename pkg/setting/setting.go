package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		} else {
			vp.AddConfigPath("configs/")
			continue
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp: vp}, nil
}
