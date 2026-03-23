package executor

import (
	"bybit-bot_ruslan/internal/logger"
	"log"

	"bybit-bot_ruslan/internal/strategy"
)

type Executor interface {
	Execute(signal strategy.Signal, candle strategy.Candle)
}

type MockExecutor struct {
	logger *logger.TradeLogger
}

func NewMockExecutor() *MockExecutor {
	return &MockExecutor{
		logger: logger.NewTradeLogger(),
	}
}

func (m *MockExecutor) Execute(signal strategy.Signal, c strategy.Candle) {

	switch signal {

	case strategy.BUY:
		log.Printf("ENTRY BUY | price=%.2f", c.Close)
		m.logger.Log("BUY", c.Close)

	case strategy.SELL:
		log.Printf("ENTRY SELL | price=%.2f", c.Close)
		m.logger.Log("SELL", c.Close)

	case strategy.EXIT:
		log.Printf("EXIT | price=%.2f", c.Close)
		m.logger.Log("EXIT", c.Close)
	}
}
