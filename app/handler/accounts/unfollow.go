package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `POST /v1/accounts/{username}/unfollow`
func (h *handler) UnFollow(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	username := chi.URLParam(r, "username")
	accountRepo := h.app.Dao.Account()
	var _account *object.Account
	var err error
	_account, err = accountRepo.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	accountAuth := auth.AccountOf(r)

	relation := new(object.Relationship)
	var flag bool
	//check whether following now
	flag, err = accountRepo.FindRelationByID(ctx, accountAuth.ID, _account.ID)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	if flag {
		err = accountRepo.UnFollowAccount(ctx, accountAuth.ID, _account.ID)
		if err != nil {
			httperror.InternalServerError(w, err)
		}

		relation.ID = accountAuth.ID
		relation.Following = flag
		relation.Followed_by = false // This func is UnFollow

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(relation); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	} else {
		relation.ID = accountAuth.ID
		relation.Following = flag
		relation.Followed_by = false // This func is UnFollow

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(relation); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}

}
