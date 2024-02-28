package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
	"proxy/models"
	"proxy/pkg/service"
	"time"
)

type Proxy struct {
	services *service.Service
}

func NewProxy(services *service.Service) *Proxy {
	return &Proxy{services: services}
}

var customTransport = http.DefaultTransport

func (p *Proxy) InitRoutes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodConnect {
			p.handleRequestTLS(w, r)
		} else {
			p.handleRequest(w, r)
		}
	})
}

func (p *Proxy) handleRequest(w http.ResponseWriter, r *http.Request) {
	r.Header.Del("Proxy-Connection")
	r.RequestURI = ""

	resp, err := customTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, "Error sending proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to parse response body")
		return
	}

	respDb, err := p.services.Response.CreateResponse(resp.StatusCode, resp.Status, resp.Header, body)
	if err != nil {
		log.Println("Failed to add response to db")
		return
	}

	err = p.services.Request.CreateRequest(r.Context(), r.Method, r.Host, r.URL.Path, r.Header, r.Cookies(), r.URL.Query(), r.Form, respDb)
	if err != nil {
		log.Println("Failed to add request to db")
		return
	}
}

func (p *Proxy) handleRequestTLS(w http.ResponseWriter, r *http.Request) {
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	err = p.services.Request.CreateRequest(r.Context(), r.Method, r.Host, r.URL.Path, r.Header, r.Cookies(), r.URL.Query(), r.Form, models.Response{})
	if err != nil {
		log.Println("Failed to add request to db")
		return
	}
	go p.transfer(dest_conn, client_conn)
	go p.transfer(client_conn, dest_conn)
}

func (p *Proxy) transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
