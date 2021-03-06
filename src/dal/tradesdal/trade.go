package tradesdal

import (
	"context"
	"fmt"
	"time"

	"github.com/gufranmirza/ticker-api/database/connection"
	"github.com/gufranmirza/ticker-api/database/dbmodels"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type trade struct {
	db connection.MongoStore
}

// NewTradesDal returns the implementation
func NewTradesDal() TradesDal {
	return &trade{
		db: connection.NewMongoStore(),
	}
}

// Create creates a new trade details.
func (r *trade) Create(txID string, job *dbmodels.Trades) (*dbmodels.Trades, error) {
	rc := r.db.Database().Collection(viper.GetString("db.trades_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	insertResult, err := rc.InsertOne(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("Failed to create trade with error %v", err)
	}

	job.ID = insertResult.InsertedID.(primitive.ObjectID)
	return job, nil
}

func (r *trade) Update(job *dbmodels.Trades) error {
	rc := r.db.Database().Collection(viper.GetString("db.trades_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if _, err := rc.ReplaceOne(ctx, bson.M{"_id": job.ID}, job); err != nil {
		return err
	}

	return nil
}

func (r *trade) GetByID(id primitive.ObjectID) (*dbmodels.Trades, error) {
	rc := r.db.Database().Collection(viper.GetString("db.trades_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	var rec dbmodels.Trades
	if err := rc.FindOne(ctx, bson.M{"_id": id}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}

func (r *trade) GetByTicker(symbol string) (*dbmodels.Trades, error) {
	rc := r.db.Database().Collection(viper.GetString("db.trades_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	var rec dbmodels.Trades
	if err := rc.FindOne(ctx, bson.M{"ticker_symbol": symbol}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}

func (r *trade) DeleteByTicker(symbol string) error {
	rc := r.db.Database().Collection(viper.GetString("db.trades_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	if _, err := rc.DeleteOne(ctx, bson.M{"ticker_symbol": symbol}); err != nil {
		return err
	}

	return nil
}

func (r *trade) GetTradesByUserID(userID primitive.ObjectID) ([]dbmodels.Trades, error) {
	rc := r.db.Database().Collection(viper.GetString("db.trades_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	filterCursor, err := rc.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return []dbmodels.Trades{}, fmt.Errorf("Failed to search trades with error %v", err)
	}

	var trades []dbmodels.Trades
	if err = filterCursor.All(ctx, &trades); err != nil {
		return []dbmodels.Trades{}, fmt.Errorf("Failed to search trades with error %v", err)
	}

	return trades, nil
}

func (r *trade) GetByTickerAndUser(symbol string, userID primitive.ObjectID) (*dbmodels.Trades, error) {
	rc := r.db.Database().Collection(viper.GetString("db.trades_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	var rec dbmodels.Trades
	if err := rc.FindOne(ctx, bson.M{"ticker_symbol": symbol, "user_id": userID}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}
