package repository

import (
	"context"
	"proxy/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestRepo struct {
	dbConn *mongo.Database
	dbColl *mongo.Collection
}

func NewRequest(db *mongo.Database) *RequestRepo {
	return &RequestRepo{dbConn: db, dbColl: db.Collection("requests")}
}

func (r *RequestRepo) Create(ctx context.Context, request *models.Request) (*models.Request, error) {
	_, err := r.dbColl.InsertOne(ctx, request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (r *RequestRepo) GetAll(ctx context.Context) ([]models.Request, error){
	var result []models.Request
	cur, err := r.dbColl.Find(ctx, bson.M{}, nil)
	if(err!=nil){
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var req models.Request
		err := cur.Decode(&req)
		if err != nil {
			return nil, err
		}
	
		result = append(result, req)
	}
	
	if err := cur.Err(); err != nil {
		return nil, err
	}
	
	cur.Close(ctx)
	return result, nil
}
	
func (r *RequestRepo) GetById(ctx context.Context, id int) (models.Request, error){
	var req models.Request
	err := r.dbColl.FindOne(ctx, bson.M{"id": id}, nil).Decode(&req)
	if(err!=nil){
		return req, err
	}
	return req, nil
}