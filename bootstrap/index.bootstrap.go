package bootstrap

import (
	"log"
	"t-card/bot"
	configs "t-card/config"
	"t-card/config/app_config"
	"t-card/config/bot_config"
	"t-card/config/cors_config"
	"t-card/database"
	chat "t-card/internal/bot"
	"t-card/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootStrapApp() {

	// LOAD .env FILE
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	// INIT CONFIG
	configs.InitConfig()

	tgBot := bot.TgBot{
		Bot:         bot.InitBot(bot_config.BOT_TOKEN),
		BotID:       bot_config.BOT_ID,
		ChannelID:   bot_config.CHANNEL_ID,
		AdminUserID: bot_config.ADMIN_USER_ID,
		TmaURL:      bot_config.TMA_URL,
	}

	go func() {
		tgBot.Bot.Start()
	}()

	defer tgBot.Bot.Stop()
	registerBotHandlers(tgBot)

	// DATABASE CONNECTION
	database.ConnectDatabase()

	// INIT GIN ENGINE
	app := gin.Default()

	// CORS
	app.Use(cors_config.CorsConfig)

	// INIT ROUTE
	routes.InitRoute(app)

	// RUN APP
	app.Run(app_config.APP_PORT)
}

func registerBotHandlers(tgBot bot.TgBot) {
	commandHandler := chat.NewCommandHandler(&tgBot)
	tgBot.Bot.Handle("/start", commandHandler.StartHandler)
}
