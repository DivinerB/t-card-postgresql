package bot_config

import (
	"os"
	"strconv"
)

var BOT_TOKEN = "8167186092:AAHR7f8gZ-70BYt5TIEdSBXe1qDxE646pZI"
var BOT_ID = "@tcardbackendbot"
var CHANNEL_ID int64 = 123
var ADMIN_USER_ID int64 = 123
var TMA_URL = "https://tcard-frontend.vercel.app/"

func InitBotConfig() {
	botTokenEnv := os.Getenv("BOT_TOKEN")

	if botTokenEnv != "" {
		BOT_TOKEN = botTokenEnv
	}

	botIDEnv := os.Getenv("BOT_ID")

	if botIDEnv != "" {
		BOT_ID = botIDEnv
	}

	channelIDEnv := os.Getenv("CHANNEL_ID")

	if channelIDEnv != "" {
		numInt64, err := strconv.ParseInt(channelIDEnv, 10, 64)
		println(err)
		num := int64(numInt64)
		CHANNEL_ID = num
	}

	adminUserIDEnv := os.Getenv("ADMIN_USER_ID")

	if adminUserIDEnv != "" {
		numInt64, err := strconv.ParseInt(channelIDEnv, 10, 64)
		println(err)
		num := int64(numInt64)
		ADMIN_USER_ID = num
	}

	tmaEnv := os.Getenv("TMA_URL")

	if tmaEnv != "" {
		TMA_URL = tmaEnv
	}
}
