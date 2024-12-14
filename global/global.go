package global

import (
	"sync"
	"waverley/wave-seo/config"

	"github.com/spf13/viper"
)

var (
	WAVE_CONFIG config.Config
	WAVE_VP     *viper.Viper
	lock        sync.RWMutex
)
