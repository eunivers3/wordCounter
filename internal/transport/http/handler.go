package http

import (
	"encoding/json"
	"errors"
	"github.com/eunicebjm/wordCounter/internal/service"
	"github.com/eunicebjm/wordCounter/pkg/models"
	"io"
	"net/http"
)

const AuthHeader = "Authorization"

// Handler is the http handler that will enable calls to this service via HTTP REST
type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) (*Handler, error) {
	if service == nil {
		return nil, errors.New("invalid_param: service")
	}
	return &Handler{
		service: service,
	}, nil
}

func (h *Handler) CountWords(w http.ResponseWriter, r *http.Request) {
	// todo: handle auth here
	if r.Header.Get(AuthHeader) == "" {
		http.Error(w, "unauthorized", 401)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "bad_request", 400)
		return
	}

	var req models.Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "bad_request", 400)
		return
	}

	// url passed in body so can fit more diverse data
	// - more complex objects if needed in future
	// long urls which won't err on req uri too long
	if req.URL == "" {
		http.Error(w, "bad_request", 400)
		return
	}

	words, err := h.service.CountWords(req.URL)
	if err != nil {
		http.Error(w, "internal_server_error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	JSONResponse, err := json.Marshal(words)
	if err != nil {
		http.Error(w, "internal_server_error", 500)
		return
	}
	// TODO: implement pagination of response
	w.Write(JSONResponse)
}
