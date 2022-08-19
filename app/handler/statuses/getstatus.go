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
func (h *handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sID := chi.URLParam(r, "id")

	Sid, err := strconv.ParseInt(sID, 10, 64)
	//status.S_id = S_id
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	statusRepo := h.app.Dao.Status()
	accountRepo := h.app.Dao.Account()
	var status_info *object.Status
	status_info, err = statusRepo.FindByID(ctx, Sid)
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	var account_info *object.Account
	uID := status_info.AccountID
	account_info, err = accountRepo.FindByID(ctx, uID)
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
