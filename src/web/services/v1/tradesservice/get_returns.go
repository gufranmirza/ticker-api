package tradesservice

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gufranmirza/ticker-api/web/interfaces/v1/tradesinterface"
	"github.com/gufranmirza/ticker-api/web/renderers"
)

// @Summary Get Portfolio returns
// @Description It returns total portfolio returns for a given user
// @Tags Trades
// @Accept  json
// @Produce  json
// @Success 200 {object} tradesinterface.ReturnsRes{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /trades/portfolio/returns [GET]
func (t *trades) PortfolioReturns(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	returns := tradesinterface.ReturnsRes{}

	res, err := t.tradesdal.GetTradesByUserID(UserID())
	if err != nil {
		t.logger.Error(txID, FailedToCreateTrade).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	for _, item := range res {
		total := float64(float64(item.Shares) * item.AveragePrice)
		returns.TotalReturnAmount += total
		returns.Trades = append(returns.Trades, tradesinterface.TradesReturn{Name: item.TickerSymbol, Return: total})
	}

	render.Respond(w, r, returns)
	return
}
