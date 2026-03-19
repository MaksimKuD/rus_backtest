package strategy

import "time"

type Signal int

const (
	HOLD Signal = iota
	BUY
	SELL
	EXIT
)

type Candle struct {
	Time  time.Time
	Open  float64
	High  float64
	Low   float64
	Close float64
}

type Strategy interface {
	OnCandle(c Candle) Signal
}
