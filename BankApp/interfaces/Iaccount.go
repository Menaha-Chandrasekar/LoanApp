package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAccount interface {
	CreateAccount(*models.Account) (*mongo.InsertOneResult, error)
	CreateManyAccount([]*models.Account) (*mongo.InsertManyResult, error)
	GetAccountById(primitive.ObjectID) (*models.Account, error)
	UpdateAccountById(primitive.ObjectID, *models.Account) (*mongo.UpdateResult, error)
	DeleteAccountById(primitive.ObjectID) (*mongo.DeleteResult, error)
}

