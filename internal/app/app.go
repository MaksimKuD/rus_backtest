package app

import (
	"log"
	"time"

	"bybit-bot_ruslan/internal/executor"
	"bybit-bot_ruslan/internal/marketdata/bybit"
	"bybit-bot_ruslan/internal/strategy"
)

type App struct {
	market     *bybit.Bybit
	exec       executor.Executor
	strategy   strategy.Strategy
	symbol     string
	tf         time.Duration
	lastCandle time.Time
}

func New(
	market *bybit.Bybit,
	exec executor.Executor,
	strat strategy.Strategy,
	symbol string,
	tf time.Duration,
) *App {

	return &App{
		market:   market,
		exec:     exec,
		strategy: strat,
		symbol:   symbol,
		tf:       tf,
	}
}

func (a *App) Run() {

	for {

		candles, err := a.market.GetCandles(a.symbol, a.tf)

		if err != nil {
			log.Println(err)
			time.Sleep(a.tf)
			continue
		}

		for _, c := range candles {

			if !c.Time.After(a.lastCandle) {
				continue
			}

			a.lastCandle = c.Time

			signal := a.strategy.OnCandle(c)

			if signal != strategy.HOLD {
				a.exec.Execute(signal, c)
			}
		}

		time.Sleep(a.tf)
	}
}
