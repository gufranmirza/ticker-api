package tradesservice

import (
	"net/http"
	"strings"
	"time"

	"github.com/gufranmirza/ticker-api/database/dbmodels"
	"github.com/gufranmirza/ticker-api/web/interfaces/v1/tradesinterface"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gufranmirza/ticker-api/web/renderers"
)

// @Summary Create new trade
// @Description It creates new trade details with the given ticker symbol
// @Tags Trades
// @Accept  json
// @Produce  json
// @Param * body tradesinterface.NewTradeReq{} true "Trade Details"
// @Param trade_id path string true "Ticker Symbol"
// @Success 200 {object} tradesinterface.NewTradeRes{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /trades/buy/{trade_id} [POST]
func (t *trades) Create(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	tickerSymbol := strings.ToUpper(chi.URLParam(r, "ticker_symbol"))

	tradeData := &tradesinterface.NewTradeReq{}
	if err := render.Bind(r, tradeData); err != nil {
		t.logger.Error(txID, FailedToCreateTrade).Error(err)
		render.Render(w, r, renderers.ErrorBadRequest(ErrIncompleteDetails))
		return
	}

	res, _ := t.tradesdal.GetByTicker(tickerSymbol)
	if res != nil {
		render.Render(w, r, renderers.ErrorBadRequest(ErrTradeAlreadyExists))
		return
	}

	trade := &dbmodels.Trades{}
	trade.CreatedTimestampUTC = time.Now().UTC()
	trade.CreatedTimestampUTC = time.Now().UTC()
	trade.TickerSymbol = tickerSymbol
	trade.AveragePrice = tradeData.AveragePrice
	trade.Shares = tradeData.Shares

	resp, err := t.tradesdal.Create(txID, trade)
	if err != nil {
		t.logger.Error(txID, FailedToCreateTrade).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	render.Respond(w, r, &tradesinterface.NewTradeRes{
		Trades: resp,
	})
	return
}
