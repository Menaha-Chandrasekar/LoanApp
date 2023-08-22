package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Loan struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitLoan(collection *mongo.Collection, ctx context.Context) interfaces.Iloan{
	return &Loan{ctx,collection}
}

func (c* Loan)CreateLoan(user *models.Loan)(*mongo.InsertOneResult,error){

	res,err := c.mongoCollection.InsertOne(c.ctx,&user)
	 if err!= nil{
		return nil,err
	 }
	 return res, nil
}

func (c *Loan)GetLoanById(id primitive.ObjectID)([]*models.Loan,error){
	match:=bson.D{{Key:"amount",Value: id}}
	result,err:=c.mongoCollection.Find(c.ctx,match)
	if err!=nil{
		return nil, err
	}else{
		var Loan_detail[] *models.Loan
		for result.Next(c.ctx){
			detail:=&models.Loan{}
			err:=result.Decode(detail)
			if err!=nil{
				return nil, err
			}
			Loan_detail=append(Loan_detail, detail)
		}
		return Loan_detail,nil
	}
}

func (c *Loan) UpdateLoanById(id primitive.ObjectID, loan *models.Loan) (*mongo.UpdateResult, error){
	update:=bson.M{"amount":id}
	updated:=bson.M{"$Set":&loan}
	result,err:=c.mongoCollection.UpdateOne(c.ctx,update,updated)
	if err!=nil{
		
		return nil,err
	}
	return result,nil
}

func (c *Loan) DeleteLoanById(id primitive.ObjectID) (*mongo.DeleteResult, error){	
	del := bson.M{"amount": id}
    res,err:=c.mongoCollection.DeleteOne(c.ctx,del)
	if err!=nil{
		return nil,err
	}
	return res, nil
}