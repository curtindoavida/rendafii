package main

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
	"time"
)

func main() {
	hoje := time.Now()
	antes := hoje.AddDate(-1, 0, 0)
	depois := hoje.AddDate(0, 1, 0)
	acao := "BRCR11.SA"
	//acao := "SNEL11.SA"
	//acao := "ADA-USD"
	strAnteOntem := "2024-10-01"
	strOntem := "2024-11-01"

	fmt.Println(antes.Format("2006-01-02"), depois.Format("2006-01-02"), strOntem, strAnteOntem)
	spy, _ := quote.NewQuoteFromYahoo(acao, antes.Format("2006-01-02"), depois.Format("2006-01-02"), quote.Daily, true)
	// fmt.Print(spy.CSV())
	fmt.Println(len(spy.Close))
	fmt.Println(spy.Close[len(spy.Close)-1])
	rsi := talib.Rsi(spy.Close, 30)
	fmt.Println(rsi[len(spy.Close)-1])
	rsi2 := talib.Rsi(spy.Close, 90)
	fmt.Println(rsi2[len(spy.Close)-1])
	rsi3 := talib.Rsi(spy.Close, len(spy.Close)-1)
	fmt.Println(rsi3[len(spy.Close)-1])
}