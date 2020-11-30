package tradesservice

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gufranmirza/ticker-api/web/interfaces/v1/tradesinterface"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gufranmirza/ticker-api/web/renderers"
)

// @Summary Get trade details
// @Description It returns the trade details for the given ticker symbol
// @Tags Trades
// @Accept  json
// @Produce  json
// @Param ticker_symbol path string true "Ticker Symbol"
// @Success 200 {object} tradesinterface.NewTradeRes{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /trades/buy/{ticker_symbol} [GET]
func (t *trades) Get(w http.ResponseWriter, r *http.Request) {
	tickerSymbol := strings.ToUpper(chi.URLParam(r, "ticker_symbol"))

	res, _ := t.tradesdal.GetByTicker(tickerSymbol)
	if res == nil {
		render.Render(w, r, renderers.ErrorNotFound(errors.New("Nothing found for the given ticker symbol")))
		return
	}

	render.Respond(w, r, &tradesinterface.NewTradeRes{
		Trades: res,
	})
	return
}
