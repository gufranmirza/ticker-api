package tradesservice

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gufranmirza/ticker-api/web/renderers"
)

// @Summary Delete trade details
// @Description It delete the trade details for the given ticker symbol
// @Tags Trades
// @Accept  json
// @Produce  json
// @Param ticker_symbol path string true "Ticker Symbol"
// @Success 200
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /trades/{ticker_symbol} [DELETE]
func (t *trades) Delete(w http.ResponseWriter, r *http.Request) {
	tickerSymbol := strings.ToUpper(chi.URLParam(r, "ticker_symbol"))

	res, _ := t.tradesdal.GetByTicker(tickerSymbol)
	if res == nil {
		render.Render(w, r, renderers.ErrorNotFound(errors.New("Nothing found for the given ticker symbol")))
		return
	}

	err := t.tradesdal.DeleteByTicker(tickerSymbol)
	if err != nil {
		render.Render(w, r, renderers.ErrorNotFound(ErrServerError))
		return
	}

	render.Respond(w, r, http.NoBody)
	return
}
