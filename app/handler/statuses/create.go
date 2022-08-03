package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
)

type StatusRequest struct {
	Status string
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req StatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	//account := new(object.Account)

	return
}
