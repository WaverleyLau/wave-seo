package global

import (
	"fmt"
	"sync"
	"waverley/wave-seo/config"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var (
	WAVE_CONFIG   config.Config
	WAVE_VP       *viper.Viper
	WAVE_VALIDATE *validator.Validate
	lock          sync.RWMutex
)

func Init() {
	fmt.Printf("initializing global config")
	// 初始化验证器
	WAVE_VALIDATE = validator.New()
}
