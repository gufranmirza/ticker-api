package tradesservice

import (
	"net/http"

	"github.com/go-chi/render"
	_ "github.com/gufranmirza/ticker-api/database/dbmodels" //swag
	"github.com/gufranmirza/ticker-api/web/renderers"
)

// @Summary Get Portfolio
// @Description It returns trades portfolio for a given user
// @Tags Trades
// @Accept  json
// @Produce  json
// @Success 200 {object} []dbmodels.Trades
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /trades/portfolio [GET]
func (t *trades) Portfolio(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]

	res, err := t.tradesdal.GetTradesByUserID(UserID())
	if err != nil {
		t.logger.Error(txID, FailedToCreateTrade).Error(err)
		render.Render(w, r, renderers.ErrorInternalServerError(err))
		return
	}

	render.Respond(w, r, res)
	return
}
