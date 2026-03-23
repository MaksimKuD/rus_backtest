package trading_engine

import (
	"bybit-bot_ruslan/internal/strategy"
	"log"
	"time"
)

type Engine struct {
	candles        []strategy.Candle
	inPosition     bool
	side           strategy.Signal
	lastSignalTime time.Time
}

func New() *Engine {
	return &Engine{
		candles: make([]strategy.Candle, 0, 200),
	}
}

func (e *Engine) OnCandle(c strategy.Candle) strategy.Signal {

	e.candles = append(e.candles, c)

	if len(e.candles) < 20 {
		return strategy.HOLD
	}

	last := e.candles[len(e.candles)-1]
	prev := e.candles[len(e.candles)-2]

	// ⛔ анти-спам сигналов (5 минут)
	if time.Since(e.lastSignalTime) < 5*time.Minute {
		return strategy.HOLD
	}

	// 📈 ВХОД
	if !e.inPosition {

		if last.Close > prev.Close {
			log.Printf("SIGNAL BUY | prev=%.2f last=%.2f", prev.Close, last.Close)

			e.inPosition = true
			e.side = strategy.BUY
			e.lastSignalTime = time.Now()

			return strategy.BUY
		}

		if last.Close < prev.Close {
			log.Printf("SIGNAL SELL | prev=%.2f last=%.2f", prev.Close, last.Close)

			e.inPosition = true
			e.side = strategy.SELL
			e.lastSignalTime = time.Now()

			return strategy.SELL
		}
	}

	// 📉 ВЫХОД
	if e.inPosition {

		if e.side == strategy.BUY && last.Close < prev.Close {
			log.Printf("EXIT BUY")

			e.inPosition = false
			e.lastSignalTime = time.Now()

			return strategy.EXIT
		}

		if e.side == strategy.SELL && last.Close > prev.Close {
			log.Printf("EXIT SELL")

			e.inPosition = false
			e.lastSignalTime = time.Now()

			return strategy.EXIT
		}
	}

	return strategy.HOLD
}
