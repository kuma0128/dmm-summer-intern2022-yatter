package accounts

import (
	"encoding/json"
	"errors"
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

	account := new(object.Account)
	if len(req.Username) > 10 {
		err := errors.New("username is too long")
		httperror.BadRequest(w, err)
		return
	}
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// domain/repository の取得
	accountRepo := h.app.Dao.Account()
	if err := accountRepo.AddAccount(ctx, account); err != nil {
		httperror.InternalServerError(w, err)
	}

	var newaccount *object.Account
	var err error
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
