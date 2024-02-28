package repository

import (
	"context"
	"proxy/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ResponseRepo struct {
	dbConn *mongo.Database
	dbColl *mongo.Collection
}

func NewResponse(db *mongo.Database) *ResponseRepo {
	return &ResponseRepo{dbConn: db, dbColl: db.Collection("requests")}
}

func (r *ResponseRepo) Create(ctx context.Context, response *models.Response) (*models.Response, error) {
	_, err := r.dbColl.InsertOne(ctx, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
