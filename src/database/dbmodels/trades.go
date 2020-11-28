package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Trades represents a trade info
type Trades struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID              primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"created_timestamp_utc,omitempty" bson:"created_timestamp_utc,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"updated_timestamp_utc,omitempty" bson:"updated_timestamp_utc,omitempty"`
	Name                string             `json:"name,omitempty" bson:"name,omitempty"`
	TickerSymbol        string             `json:"ticker_symbol,omitempty" bson:"ticker_symbol,omitempty"`
	AveragePrice        float64            `json:"average_price,omitempty" bson:"average_price,omitempty"`
	Shares              int                `json:"shares,omitempty" bson:"shares,omitempty"`
}
