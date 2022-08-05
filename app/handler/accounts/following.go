package accounts

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `GET /v1/accounts/{username}/following`
func (h *handler) Following(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")

	limit := r.FormValue("limit")
	if limit == "" {
		limit = "40"
	}
	Limit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	if 80 < Limit {
		Limit = 80
	}

	accountRepo := h.app.Dao.Account()
	var _account *object.Account
	_account, err = accountRepo.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	var accounts []*object.Account
	accounts, err = accountRepo.FingFollowerByName(ctx, _account.ID, Limit)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}
