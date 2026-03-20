package bybit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"bybit-bot_ruslan/internal/strategy"
)

type Bybit struct {
	client *http.Client
}

func New() *Bybit {

	return &Bybit{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

type response struct {
	Result struct {
		List [][]string `json:"list"`
	} `json:"result"`
}

func (b *Bybit) GetCandles(symbol string, tf time.Duration) ([]strategy.Candle, error) {

	url := fmt.Sprintf(
		"https://api.bybit.com/v5/market/kline?category=linear&symbol=%s&interval=5&limit=1",
		symbol,
	)

	resp, err := b.client.Get(url)
	if err != nil {
		log.Println("API ERROR:", err)
		return nil, err
	}

	defer resp.Body.Close()

	var r response

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	var candles []strategy.Candle

	for _, v := range r.Result.List {

		ts, _ := strconv.ParseInt(v[0], 10, 64)

		open, _ := strconv.ParseFloat(v[1], 64)
		high, _ := strconv.ParseFloat(v[2], 64)
		low, _ := strconv.ParseFloat(v[3], 64)
		closep, _ := strconv.ParseFloat(v[4], 64)

		c := strategy.Candle{
			Time:  time.UnixMilli(ts),
			Open:  open,
			High:  high,
			Low:   low,
			Close: closep,
		}

		candles = append(candles, c)
	}

	return candles, nil
}
