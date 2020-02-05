// Copyright 2019 Tokenomy Technologies Pte. Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1

import (
	"fmt"
	"log"
	"os"

	"github.com/tokenomy/tokenomy-go"
)

func ExampleClient_Buy() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	tres, err := cl.Buy(tokenomy.PairTokenomyBitcoin, 100, 0.00005)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Buy response: %+v\n", tres)
}

func ExampleClient_BuyByMarket() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	tres, err := cl.BuyByMarket(tokenomy.PairTokenomyBitcoin, 0.0001)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Buy by market response: %+v\n", tres)
}

//nolint:dupl
func ExampleClient_CancelBuy() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	openOrders, err := cl.ListOpenOrders(tokenomy.PairTokenomyBitcoin)
	if err != nil {
		log.Fatal(err)
	}
	if len(openOrders) == 0 {
		fmt.Println("No open orders to cancel")
		return
	}

	var openBuy *OrderHistory

	fmt.Println("Open orders:")
	for pairName, list := range openOrders {
		fmt.Printf("[%s]\n", pairName)
		for x, order := range list {
			if order.Type == tokenomy.TradeTypeBid && openBuy == nil {
				openBuy = order
			}
			fmt.Printf("  [%d] - %+v\n", x, order)
		}
	}

	if openBuy == nil {
		fmt.Println("There is no open bid to cancel.")
		return
	}

	fmt.Printf("Canceling the first open bid with ID '%d'\n", openBuy.OrderID)

	cancelRes, err := cl.CancelBuy(tokenomy.PairTokenomyBitcoin, openBuy.OrderID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cancel bid response: %+v\n", cancelRes)
}

//nolint:dupl
func ExampleClient_CancelSell() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	openOrders, err := cl.ListOpenOrders(tokenomy.PairTokenomyBitcoin)
	if err != nil {
		log.Fatal(err)
	}
	if len(openOrders) == 0 {
		fmt.Println("No open orders to cancel")
		return
	}

	var openSell *OrderHistory

	fmt.Println("Open orders:")
	for pairName, list := range openOrders {
		fmt.Printf("[%s]\n", pairName)
		for x, order := range list {
			if order.Type == tokenomy.TradeTypeAsk && openSell == nil {
				openSell = order
			}
			fmt.Printf("  [%d] - %+v\n", x, order)
		}
	}
	if openSell == nil {
		fmt.Println("There is no open sell (ask) to cancel.")
		return
	}

	fmt.Printf("Canceling the open sell (ask) with ID '%d'\n", openSell.OrderID)

	cancelRes, err := cl.CancelSell(tokenomy.PairTokenomyBitcoin, openSell.OrderID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cancel sell (ask) response: %+v\n", cancelRes)
}

func ExampleClient_GetOrder() {
	env := tokenomy.NewEnvironment("", "")
	orderID := int64(1023965)

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	order, err := cl.GetOrder(tokenomy.PairTokenomyBitcoin, orderID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Order detail for %d: %+v\n", orderID, order)
}

func ExampleClient_GetTicker() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	pair, err := cl.GetTicker(tokenomy.PairTokenomyBitcoin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pair information for %q: %+v\n", tokenomy.PairTokenomyBitcoin, pair)
}

func ExampleClient_ListOpenOrders() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	openOrders, err := cl.ListOpenOrders(tokenomy.PairTokenomyBitcoin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Open orders:")
	for pairName, list := range openOrders {
		fmt.Printf("[%s]\n", pairName)
		for x, order := range list {
			fmt.Printf("  [%d] - %+v\n", x, order)
		}
	}
}

func ExampleClient_ListOrderHistory() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	orderHistory, err := cl.ListOrderHistory(tokenomy.PairTokenomyBitcoin, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Order history:")
	for x, order := range orderHistory {
		fmt.Printf("  [%d] %+v\n", x, order)
	}
}

func ExampleClient_ListTrades() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	trades, err := cl.ListTrades(tokenomy.PairTokenomyBitcoin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Trades information for pair %q:", tokenomy.PairTokenomyBitcoin)
	for _, trade := range trades {
		fmt.Printf("  %+v\n", trade)
	}
}

func ExampleClient_ListTradeHistory() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	tradeHistory, err := cl.ListTradeHistory(tokenomy.PairTokenomyBitcoin,
		1, 0, 0, "asc", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Trade history:")
	for x, trade := range tradeHistory {
		fmt.Printf("  [%d]: %+v\n", x, trade)
	}
}

func ExampleClient_ListTransactionHistory() {
	env := tokenomy.NewEnvironment(
		os.Getenv(tokenomy.EnvNameToken),
		os.Getenv(tokenomy.EnvNameSecret),
	)

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	transHistory, err := cl.ListTransactionHistory()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction history: %+v\n", transHistory)
}

func ExampleClient_MarketInfo() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	marketInfos, err := cl.MarketInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Market Info: %+v\n", marketInfos)
}

func ExampleClient_OrderBook() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	orderBook, err := cl.OrderBook(tokenomy.PairTokenomyBitcoin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Order book buys:")
	for _, buy := range orderBook.Buys {
		fmt.Printf("  %+v\n", buy)
	}

	fmt.Println("Order book sells:")
	for _, sell := range orderBook.Sells {
		fmt.Printf("  %+v\n", sell)
	}
}

func ExampleClient_SellByMarket() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	tres, err := cl.SellByMarket(tokenomy.PairTokenomyBitcoin, 20)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sell by market response: %+v\n", tres)
}

func ExampleClient_Summaries() {
	env := tokenomy.NewEnvironment("", "")

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	summary, err := cl.Summaries()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Summaries:")
	fmt.Println("  Tickers:")
	for name, ticker := range summary.Pairs {
		fmt.Printf("    %s: %+v\n", name, ticker)
	}
	fmt.Printf("  Prices for the last 24 hours: %v\n", summary.PricesLast7Days)
	fmt.Printf("  Prices for the last 7 days  : %v\n", summary.PricesLast24Hours)
}

func ExampleClient_UserInfo() {
	env := tokenomy.NewEnvironment(
		os.Getenv(tokenomy.EnvNameToken),
		os.Getenv(tokenomy.EnvNameSecret),
	)

	cl, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	userInfo, err := cl.UserInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", userInfo)
}
