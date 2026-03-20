package trading_engine

import (
	"bybit-bot_ruslan/internal/strategy"
	"log"
)

type Engine struct {
	candles    []strategy.Candle
	inPosition bool
	side       strategy.Signal
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

	// если нет позиции
	if !e.inPosition {

		if last.Close > prev.Close {
			log.Printf("SIGNAL BUY | prev=%.2f last=%.2f", prev.Close, last.Close)
			return strategy.BUY
		}

		if last.Close < prev.Close {
			log.Printf("SIGNAL SELL | prev=%.2f last=%.2f", prev.Close, last.Close)
			return strategy.SELL
		}

	}

	// если позиция открыта
	if e.inPosition {

		if e.side == strategy.BUY && last.Close < prev.Close {
			e.inPosition = false
			return strategy.EXIT
		}

		if e.side == strategy.SELL && last.Close > prev.Close {
			e.inPosition = false
			return strategy.EXIT
		}

	}

	return strategy.HOLD
}
