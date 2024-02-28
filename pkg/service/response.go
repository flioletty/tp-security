package service

import (
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

func (s *ResponseService) CreateResponse(status int,
	message string,
	headers http.Header,
	body []byte,
) (models.Response, error) {

	response := models.Response{Status: status,
		Message: message,
	}
	response.Headers = make(map[string]any)
	for key, value := range headers {
		response.Headers[key] = value
	}
	response.Body = string(body)

	return response, nil
}
