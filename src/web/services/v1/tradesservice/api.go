package tradesservice

import (
	"net/http"
)

// TradesService interface ...
type TradesService interface {
	Create(w http.ResponseWriter, r *http.Request)
}
