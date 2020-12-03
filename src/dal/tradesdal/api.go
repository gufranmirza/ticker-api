package tradesdal

import (
	"github.com/gufranmirza/ticker-api/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TradesDal interface
type TradesDal interface {
	Create(txID string, job *dbmodels.Trades) (*dbmodels.Trades, error)
	Update(job *dbmodels.Trades) error
	GetByID(id primitive.ObjectID) (*dbmodels.Trades, error)
	GetByTicker(symbol string) (*dbmodels.Trades, error)
	DeleteByTicker(symbol string) error
	GetTradesByUserID(userID primitive.ObjectID) ([]dbmodels.Trades, error)
	GetByTickerAndUser(symbol string, userID primitive.ObjectID) (*dbmodels.Trades, error)
}
