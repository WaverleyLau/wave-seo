package app

import "waverley/wave-seo/global"

func InitApp() {
	// 初始化全局信息
	InitGlobal()
}

func InitGlobal() {
	global.WAVE_VP = InitConfig()

}
