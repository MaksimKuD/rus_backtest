package main

import (
	"log"
	"time"

	"bybit-bot_ruslan/internal/app"
	"bybit-bot_ruslan/internal/executor"
	"bybit-bot_ruslan/internal/logger"
	mdbybit "bybit-bot_ruslan/internal/marketdata/bybit"
	"bybit-bot_ruslan/internal/strategy/trading_engine"
)

func main() {

	logFile := logger.InitLogger()
	defer logFile.Close()

	log.Println("🚀 Bot started")

	market := mdbybit.New()
	exec := executor.NewMockExecutor()
	strategy := trading_engine.New()

	app := app.New(
		market,
		exec,
		strategy,
		"BTCUSDT",
		5*time.Minute,
	)

	app.Run()
}
