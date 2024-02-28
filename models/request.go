package models

type Request struct {
	Id         int
	Method     string
	Url        string
	Host       string
	GetParams  map[string][]string
	Headers    map[string][]string
	Cookies    map[string]string
	PostParams map[string][]string
	Response   Response
}
