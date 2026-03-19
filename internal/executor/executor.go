package executor

import (
	"log"

	"bybit-bot_ruslan/internal/strategy"
)

type Executor interface {
	Execute(signal strategy.Signal, candle strategy.Candle)
}

type MockExecutor struct {
}

func NewMockExecutor() *MockExecutor {
	return &MockExecutor{}
}

func (m *MockExecutor) Execute(signal strategy.Signal, c strategy.Candle) {

	switch signal {

	case strategy.BUY:
		log.Println("BUY at", c.Close)

	case strategy.SELL:
		log.Println("SELL at", c.Close)

	case strategy.EXIT:
		log.Println("EXIT at", c.Close)
	}
}
