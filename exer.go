package ex

import (
	"github.com/hxzqlh/exchange/types"
)

type Exer interface {
	GetExchangeName() string

	GetCurrencies() ([]*types.Currency, error)
	GetSymbols() ([]*types.Symbol, error)

	GetAccount(accountId string) (*types.Account, error)

	PlaceOrder(o *types.Order) (*types.Order, error)
	CancelOrder(orderId string) (*types.Order, error)
	GetOneOrder(orderId string) (*types.Order, error)
	GetActiveOrders(accountId, symbol string) ([]*types.Order, error)
	// from: time(ms), direct: prev or next
	GetOrderHistorys(accountId, symbol, direct string, size, from int64) ([]*types.Order, error)

	GetTicker(symbol string) (*types.Ticker, error)
	GetDepths(symbol, depth string) (*types.Depths, error)
	GetTrades(symbol, direct string, size, from int64) ([]*types.Trade, error)
	GetTradesByOrder(orderId string) ([]*types.Trade, error)
	GetKlines(symbol, period string, size int64) ([]*types.Kline, error)

	// TODO
	// QueryDepoitRecords()
	// QueryWithdrawRecords()
	// ApplyWithDraw()
	// CancelWithDraw()
}
