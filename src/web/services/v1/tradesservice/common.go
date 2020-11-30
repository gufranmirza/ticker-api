package tradesservice

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserID returns the userid
func UserID() primitive.ObjectID {
	userID, _ := primitive.ObjectIDFromHex("5fc405cc508561b6d1511ef6")
	return userID
}
