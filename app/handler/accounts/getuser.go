package accounts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type UserRequest struct {
	Username string
}

// Handle request for `GET /v1/accounts/{username}`
func (h *handler) Getuser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var req UserRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	httperror.BadRequest(w, err)
	// 	return
	// }
	username, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	account := new(object.Account)
	account.Username = string(username)

	// domain/repository の取得
	accountRepo := h.app.Dao.Account()
	if _, err := accountRepo.FindByUsername(ctx, account.Username); err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}
