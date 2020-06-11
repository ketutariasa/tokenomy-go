// Copyright 2019 Tokenomy Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v1

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shuLhan/share/lib/math/big"

	"github.com/tokenomy/tokenomy-go"
)

const (
	DefaultAddress = "https://exchange.tokenomy.com"
)

//
// List of APIs.
//
const (
	apiPrivate          = "/tapi"
	apiMarketSummaries  = "/api/summaries"
	apiMarketTicker     = "/api/%s/ticker"
	apiMarketTrades     = "/api/%s/trades"
	apiMarketOrdersOpen = "/api/%s/depth"
	apiMarketInfo       = "/api/market_info"
)

//
// List of "method" form value for private API.
//
const (
	MethodTrade            = "trade"
	MethodTradeCancelOrder = "cancelOrder"
	MethodUserInfo         = "getInfo"
	MethodUserOrder        = "getOrder"
	MethodUserOrdersOpen   = "openOrders"
	MethodUserOrderHistory = "orderHistory"
	MethodUserTradeHistory = "tradeHistory"
	MethodUserTransHistory = "transHistory"
	MethodUserWithdraw     = "withdrawCoin"
)

// List of common JSON field names.
const (
	fieldNameAmount            = "amount"
	fieldNameBalance           = "balance"
	fieldNameBaseCurrency      = "base_currency"
	fieldNameBaseCurrencyPrice = "base_currency_price"
	fieldNameDate              = "date"
	fieldNameError             = "error"
	fieldNameErrorCode         = "error_code"
	fieldNameFinishTime        = "finish_time"
	fieldNameIsError           = "is_error"
	fieldNameMethod            = "method"
	fieldNameOrderID           = "order_id"
	fieldNamePair              = "pair"
	fieldNamePrice             = "price"
	fieldNameStatus            = "status"
	fieldNameSubmitTime        = "submit_time"
	fieldNameSuccess           = "success"
	fieldNameTID               = "tid"
	fieldNameTradeID           = "trade_id"
	fieldNameTradeTime         = "trade_time"
	fieldNameTradeTimePrint    = "trade_time_print"
	fieldNameType              = "type"
)

const (
	responseSuccess = 1
)

var (
	// ErrUnauthenticated define an error when user did not provide token
	// and secret keys when accessing private APIs.
	ErrUnauthenticated = fmt.Errorf("unauthenticated connection")

	// ErrInvalidPairName define an error if user call API with empty,
	// invalid or unknown pair's name.
	ErrInvalidPairName = fmt.Errorf("invalid or empty pair name")
)

//
// convertBalance convert the v1's Balance to tokenomy's UserAssets.
//
func convertBalance(balances map[string]*big.Rat) (userAssets *tokenomy.UserAssets) {
	userAssets = tokenomy.NewUserAssets()

	for k, v := range balances {
		if strings.HasPrefix(k, "frozen_") {
			k = strings.TrimPrefix(k, "frozen_")
			userAssets.FrozenBalances[k] = big.NewRat(v)
		} else {
			userAssets.Balances[k] = big.NewRat(v)
		}
	}

	return userAssets
}

//
// timestamp return current time in milliseconds as integer.
//
func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//
// timestampAsString return current time in milliseconds as string.
//
func timestampAsString() string {
	return strconv.FormatInt(timestamp(), 10)
}
