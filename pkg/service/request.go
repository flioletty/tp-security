package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"proxy/models"
	"proxy/pkg/repository"
)

type RequestService struct {
	RepoRequest repository.RequestRepo
}

func NewRequestService(repoRequest repository.RequestRepo) *RequestService {
	return &RequestService{RepoRequest: repoRequest}
}

func (s *RequestService) CreateRequest(ctx context.Context,
	method,
	host,
	url string,
	headers http.Header,
	cookies []*http.Cookie,
	getParams url.Values,
	postParams url.Values) error {

	request := models.Request{Method: method,
		Url:    url,
		Host:   host,
	}
	request.Headers = make(map[string][]string)
	for key, value := range headers {
		request.Headers[key] = value
	}
	request.Cookies = make(map[string]string)
	for _, cookie := range cookies {
		request.Cookies[cookie.Name] = cookie.Value
	}
	request.GetParams = make(map[string][]string)
	for key, value := range getParams {
		request.GetParams[key] = value
	}

	request.PostParams = make(map[string][]string)
	for key, value := range postParams {
		request.PostParams[key] = value
	}

	if _, err := s.RepoRequest.Create(ctx, &request); err != nil {
		return err
	}

	return nil
}

func (s *RequestService) GetAll(ctx context.Context) ([]models.Request, error) {
	return s.RepoRequest.GetAll(ctx)
}

func (s *RequestService) GetById(ctx context.Context, id int) (models.Request, error) {
	return s.RepoRequest.GetById(ctx, id)
}


func (s *RequestService) ConvertToHttpRequest(request models.Request) (*http.Request, error) {
	body, err := json.Marshal(request.PostParams)
	if err != nil {
		return nil, err
	}
	reqHttp, err := http.NewRequest(request.Method, request.Url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	reqHttp.Host = request.Host
	reqHttp.URL.Host = request.Host

	query := reqHttp.URL.Query()
	for key, value := range request.GetParams {
		for _, item := range value {
			query.Add(key, item)
		}
	}
	reqHttp.URL.RawQuery = query.Encode()

	for key, headers := range request.Headers {
		for _, header := range headers {
			reqHttp.Header.Set(key, fmt.Sprint(header))
		}
	}
	for key, value := range request.Cookies {
		reqHttp.AddCookie(&http.Cookie{Name: key, Value: fmt.Sprint(value)})
	}

	return reqHttp, nil
}
