package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Iloan interface{
	CreateLoan(*models.Loan)(*mongo.InsertOneResult,error)
	//CreateManyLoan([]*models.Loan)(*mongo.InsertManyResult,error)
	GetLoanById(primitive.ObjectID) ([]*models.Loan, error)
	UpdateLoanById(primitive.ObjectID, *models.Loan) (*mongo.UpdateResult, error)
	DeleteLoanById(primitive.ObjectID) (*mongo.DeleteResult, error)
}