package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `POST /v1/accounts/{username}/follow`
func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")

	accountRepo := h.app.Dao.Account()
	var _account *object.Account
	var err error
	_account, err = accountRepo.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	Account_auth := auth.AccountOf(r)

	relation := new(object.Relationship)
	var flag bool
	flag, err = accountRepo.FindRelationByID(ctx, Account_auth.ID, _account.ID)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	if !flag {
		err = accountRepo.FollowAccount(ctx, Account_auth.ID, _account.ID)
		if err != nil {
			httperror.InternalServerError(w, err)
		}

		relation.ID = Account_auth.ID
		relation.Following = flag
		relation.Followed_by = true // This func is Follow

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(relation); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}

	relation.ID = Account_auth.ID
	relation.Following = flag
	relation.Followed_by = true // This func is Follow

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relation); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
