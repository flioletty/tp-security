package repository

import (
	"context"
	"proxy/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Request  *RequestRepo
	Response *ResponseRepo
}

func NewRepository(dbConn *mongo.Database) *Repository {
	return &Repository{
		Request:  NewRequest(dbConn),
		Response: NewResponse(dbConn),
	}
}

type IRequestRepo interface {
	Create(ctx context.Context, req []byte) (int, error)
	GetAll(ctx context.Context) ([]models.Request, error)
	GetById(ctx context.Context, id int) (models.Request, error)
}

type IResponceRepo interface {
	Create(ctx context.Context, req []byte) (int, error)
}
