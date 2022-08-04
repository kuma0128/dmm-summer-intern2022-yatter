package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `GET /v1/statuses/{id}`
func (h *handler) Getstatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	s_id := chi.URLParam(r, "id")

	//status := new(object.Status)
	S_id, err := strconv.ParseInt(s_id, 10, 64)
	//status.S_id = S_id
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	statusRepo := h.app.Dao.Status()
	var status_info *object.Status
	status_info, err = statusRepo.FindByStatusID(ctx, S_id)
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	var account_info *object.Account
	uid := status_info.AccountID
	account_info, err = statusRepo.FindByAccountID(ctx, uid)
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	status_info.Account = *account_info

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status_info); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
