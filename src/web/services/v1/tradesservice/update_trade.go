package tradesservice

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gufranmirza/ticker-api/web/interfaces/v1/tradesinterface"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gufranmirza/ticker-api/web/renderers"
)

// @Summary Update a trade
// @Description It allows to update trade details with given ticker symbol
// @Tags Trades
// @Accept  json
// @Produce  json
// @Param * body tradesinterface.NewTradeReq{} true "Trade Details"
// @Param ticker_symbol path string true "Ticker Symbol"
// @Success 200 {object} tradesinterface.NewTradeRes{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /trades/{ticker_symbol} [PUT]
func (t *trades) Update(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	tickerSymbol := strings.ToUpper(chi.URLParam(r, "ticker_symbol"))

	tradeData := &tradesinterface.NewTradeReq{}
	if err := render.Bind(r, tradeData); err != nil {
		t.logger.Error(txID, FailedToCreateTrade).Error(err)
		render.Render(w, r, renderers.ErrorBadRequest(ErrIncompleteDetails))
		return
	}

	res, _ := t.tradesdal.GetByTicker(tickerSymbol)
	if res == nil {
		render.Render(w, r, renderers.ErrorNotFound(errors.New("Nothing found for the given ticker symbol, first create trade before updating")))
		return
	}

	res.AveragePrice = ((float64(res.Shares) * res.AveragePrice) + (float64(tradeData.Shares) * tradeData.AveragePrice)) / (float64(res.Shares) + float64(tradeData.Shares))
	res.Shares += tradeData.Shares
	res.UpdatedTimestampUTC = time.Now().UTC()

	err := t.tradesdal.Update(res)
	if err != nil {
		t.logger.Error(txID, FailedToCreateTrade).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	render.Respond(w, r, &tradesinterface.NewTradeRes{
		Trades: res,
	})
	return
}
