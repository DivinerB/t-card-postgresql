package configs

import (
	"t-card/config/app_config"
	"t-card/config/bot_config"
	"t-card/config/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
	bot_config.InitBotConfig()
}
