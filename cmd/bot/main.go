package main

import (
	"log"
	"time"

	"bybit-bot_ruslan/internal/app"
	"bybit-bot_ruslan/internal/executor"
	mdbybit "bybit-bot_ruslan/internal/marketdata/bybit"
	"bybit-bot_ruslan/internal/strategy/trading_engine"
)

func main() {

	market := mdbybit.New()

	exec := executor.NewMockExecutor()

	strategy := trading_engine.New()

	a := app.New(
		market,
		exec,
		strategy,
		"BTCUSDT",
		5*time.Minute,
	)

	log.Println("Bot started")

	a.Run()
}
