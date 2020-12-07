package tradesservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gufranmirza/ticker-api/web/interfaces/v1/tradesinterface"
	"github.com/gufranmirza/ticker-api/web/renderers"
)

type Price struct {
	Price float64 `json:"price"`
}

type Response struct {
	Data []Price `json:"data"`
}

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
		price, err := GetPrice(item.TickerSymbol)
		if err != nil {
			t.logger.Error(txID, "ERROR").Errorf("Failed to reterive current price with error %v", err)
			continue
		}

		total := float64(float64(item.Shares) * (price - item.AveragePrice))
		returns.TotalReturnAmount += total
		returns.Trades = append(returns.Trades, tradesinterface.TradesReturn{Name: item.TickerSymbol, Return: total})
	}

	render.Respond(w, r, returns)
	return
}

func GetPrice(ticker string) (float64, error) {
	url := fmt.Sprintf("https://quotes-api.tickertape.in/quotes?sids=%s", ticker)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Non ok response from server")
	}

	var data Response

	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, fmt.Errorf("Failed to unmarshal with error %v", err)
	}

	fmt.Println(data.Data[0].Price)

	return data.Data[0].Price, nil
}
