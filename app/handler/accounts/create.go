package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	// domain/repository の取得
	accountRepo := h.app.Dao.Account()

	account, err := object.CreateAccountobject(req.Username, req.Password)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	if err := accountRepo.AddAccount(ctx, account); err != nil {
		httperror.InternalServerError(w, err)
	}

	var newaccount *object.Account
	newaccount, err = accountRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newaccount); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
