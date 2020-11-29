package tradesservice

import (
	"github.com/gufranmirza/ticker-api/dal/tradesdal"
	"github.com/gufranmirza/ticker-api/logging"
)

type trades struct {
	tradesdal tradesdal.TradesDal
	logger    logging.Logger
}

// NewTradesService returns service impl
func NewTradesService() TradesService {
	return &trades{
		tradesdal: tradesdal.NewTradesDal(),
		logger:    logging.NewLogger(),
	}
}
