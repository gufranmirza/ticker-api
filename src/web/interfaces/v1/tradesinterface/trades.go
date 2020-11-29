package tradesinterface

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gufranmirza/ticker-api/database/dbmodels"
)

// NewTradeReq holds the new trade request data
type NewTradeReq struct {
	AveragePrice float64 `json:"average_price,omitempty"`
	Shares       int     `json:"shares,omitempty"`
}

// NewTradeRes holds the a trade data
type NewTradeRes struct {
	*dbmodels.Trades
}

// ============== Validation ============== //

// Bind valide the request interface with rules given
func (body *NewTradeReq) Bind(r *http.Request) error {
	return validation.ValidateStruct(body,
		validation.Field(&body.AveragePrice, validation.Required, validation.Min(0.0)),
		validation.Field(&body.Shares, validation.Required, validation.Min(1)),
	)
}
