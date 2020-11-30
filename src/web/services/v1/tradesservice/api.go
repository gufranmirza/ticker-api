package tradesservice

import (
	"errors"
	"net/http"
)

// TradesService interface ...
type TradesService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

// The list of error types presented to the end user as error message.
var (
	ErrIncompleteDetails  = errors.New("Incorrect details provided, please provice correct details")
	ErrTradeAlreadyExists = errors.New("Trade with the given ticker symbol already exists, please try updating instead")
	ErrServerError        = errors.New("Something went wrong at server side, please try again")
)

// List of error codes used in Trades service/model
var (
	FailedToCreateTrade = "Failed-To-Create-Trade"
	FailedToGetTrade    = "Failed-To-Get-Trade"
	FailedToUpdateTrade = "Failed-To-Update-Trade"
	FailedToDeleteTrade = "Failed-To-Delete-Trade"
)
