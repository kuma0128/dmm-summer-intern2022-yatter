package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
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
	//autholize
	accountAuth := auth.AccountOf(r)

	statusRepo := h.app.Dao.Status()

	status, err := object.CreateStatusobject(req.Status, accountAuth)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	newstatus, err := statusRepo.AddStatus(ctx, status, accountAuth.ID)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	//get create_at time
	tmpstatus, err := statusRepo.FindByID(ctx, newstatus.Sid)
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	newstatus.CreateAt = tmpstatus.CreateAt

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newstatus); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
