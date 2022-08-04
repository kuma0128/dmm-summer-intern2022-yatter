package statuses

import (
	"database/sql"
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
	Account_auth := auth.AccountOf(r)

	statusRepo := h.app.Dao.Status()
	status := new(object.Status)
	status.Content = req.Status

	status.Account = *Account_auth

	var newstatus *object.Status
	var err error
	var result sql.Result
	newstatus, result, err = statusRepo.AddStatus(ctx, status, Account_auth)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	newstatus.Sid, err = result.LastInsertId()
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	newstatus, err = statusRepo.FindStatusByID(ctx, newstatus.Sid)
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	newstatus.Account = status.Account

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newstatus); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
