package service

import (
	"context"
	"net/http"
	"proxy/models"
	"proxy/pkg/repository"
)

type ResponseService struct {
	RepoResponse repository.ResponseRepo
}

func NewResponseService(repoResponse repository.ResponseRepo) *ResponseService {
	return &ResponseService{RepoResponse: repoResponse}
}

func (s *ResponseService) CreateResponse(ctx context.Context,
	status int,
	message string,
	headers http.Header,
	body []byte,
) error {

	response := models.Response{Status: status,
		Message: message,
	}
	response.Headers = make(map[string]any)
	for key, value := range headers {
		response.Headers[key] = value
	}
	response.Body = string(body)

	if _, err := s.RepoResponse.Create(ctx, &response); err != nil {
		return err
	}

	return nil
}
