package main

import (
	"fmt"
	"time"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {

	acoes := []string{"MALL11", "HGBS11", "HSML11", "VISC11", "XPML11", "HGLG11", "GGRC11", "BTLG11", "VILG11", "XPIN11", "XPLG11", "BRCR11", "HGBS11", "JSRE11", "HGRE11", "VINO11", "SNEL11", "FIIB11", "MXRF11"}
	for _, acao := range acoes {
		fmt.Print(acao + ": ")
		fmt.Println(getPrice(acao))
	}
}

func getPrice(ticker string) (float64, float64, float64, float64) {
	hoje := time.Now()
	antes := hoje.AddDate(-1, 0, 0)
	depois := hoje.AddDate(0, 1, 0)
	acao := ticker + ".SA"

	spy, err := quote.NewQuoteFromYahoo(acao, antes.Format("2006-01-02"), depois.Format("2006-01-02"), quote.Daily, true)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Print(spy.CSV())
	// fmt.Println(len(spy.Close))
	//	fmt.Println(spy.Close[len(spy.Close)-1])
	rsi := talib.Rsi(spy.Close, 30)
	//	fmt.Println(rsi[len(spy.Close)-1])
	rsi2 := talib.Rsi(spy.Close, 90)
	//	fmt.Println(rsi2[len(spy.Close)-1])
	rsi3 := talib.Rsi(spy.Close, len(spy.Close)-1)
	//	fmt.Println(rsi3[len(spy.Close)-1])
	return spy.Close[len(spy.Close)-1], rsi[len(spy.Close)-1], rsi2[len(spy.Close)-1], rsi3[len(spy.Close)-1]
}
