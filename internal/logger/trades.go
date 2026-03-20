package logger

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type TradeLogger struct {
	file   *os.File
	writer *csv.Writer
}

func NewTradeLogger() *TradeLogger {

	_ = os.MkdirAll("logs", os.ModePerm)

	file, _ := os.OpenFile("logs/trades.csv",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	writer := csv.NewWriter(file)

	// заголовок
	writer.Write([]string{
		"time", "type", "price",
	})

	writer.Flush()

	return &TradeLogger{
		file:   file,
		writer: writer,
	}
}

func (t *TradeLogger) Log(signal string, price float64) {

	t.writer.Write([]string{
		time.Now().Format(time.RFC3339),
		signal,
		strconv.FormatFloat(price, 'f', 2, 64),
	})

	t.writer.Flush()
}
