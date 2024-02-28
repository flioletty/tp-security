package handler

import (
	"io"
	"log"
	"net/http"
	"proxy/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) getRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := h.services.Request.GetAll(r.Context())

	if err != nil {
		log.Println(err)
		NewResponseDto(r.Context(), w, http.StatusInternalServerError, "internal server error", "")
		return
	}

	NewResponseDto(r.Context(), w, 200, "success!", requests)
}

func (h *Handler) getRequestById(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		NewResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params", "")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		NewResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params", "")
		return
	}

	request, err := h.services.Request.GetById(r.Context(), id)

	if err != nil {
		NewResponseDto(r.Context(), w, http.StatusInternalServerError, "internal server error", "")
		return
	}

	NewResponseDto(r.Context(), w, 200, "success!", request)
}

func (h *Handler) repeatRequestById(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		NewResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params", "")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		NewResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params", "")
		return
	}

	request, err := h.services.Request.GetById(r.Context(), id)

	if err != nil {
		NewResponseDto(r.Context(), w, http.StatusInternalServerError, "internal server error", "")
		return
	}

	data, err := h.repeatRequest(request)
	if err != nil {
		NewResponseDto(r.Context(), w, http.StatusInternalServerError, "failed to repeat request", "")
		return
	}

	NewResponseDto(r.Context(), w, 200, "success!", string(data))
}

func (h *Handler) repeatRequest(request models.Request) ([]byte, error) {
	reqHttp, err := h.services.Request.ConvertToHttpRequest(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultTransport.RoundTrip(reqHttp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	h.services.Response.CreateResponse(reqHttp.Context(), resp.StatusCode, resp.Status, resp.Header, data)
	return data, nil
}

func (h *Handler) scanRequestById(w http.ResponseWriter, r *http.Request) {

}
