package http

import (
	"encoding/json"
	"fmt"
	"github.com/eunicebjm/gc/pkg/models"
	"io/ioutil"

	"github.com/eunicebjm/gc/internal/service"
	"net/http"
)

const AuthHeader = "Authorization"

// Handler is the http handler that will enable calls to this service via HTTP REST
type Handler struct {
	service service.Geocoder
}

// NewHandler will create a new instance of httpHandler
func NewHandler(service service.Geocoder) (*Handler, error) {
	if service == nil {
		return nil, fmt.Errorf("invalid_param: service")
	}

	return &Handler{
		service: service,
	}, nil
}

func (h *Handler) GeocodeOne(w http.ResponseWriter, r *http.Request) {
	// todo: handle auth here
	if r.Header.Get(AuthHeader) == ""  {
		http.Error(w, "unauthorized", 401)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
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

	if req.Address == "" {
		http.Error(w, "bad_request", 400)
		return
	}

	location, err := h.service.GeocodeOne(req.Address)
	if err != nil {
		http.Error(w, "internal_server_error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	JSONResponse, err := json.Marshal(location)
	if err != nil {
		http.Error(w, "internal_server_error", 500)
		return
	}

	//_ = json.NewEncoder(w).Encode(location)
	w.Write(JSONResponse)
}
