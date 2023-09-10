package setting

import (
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()

	for _, config := range configs {
		if config != "" {
			configBaseName := path.Base(config)
			ext := path.Ext(config)
			vp.SetConfigName(strings.TrimSuffix(configBaseName, ext))
			configPath := path.Dir(config)
			vp.AddConfigPath(configPath)
			vp.SetConfigType(ext[1:])
		}
	}

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
