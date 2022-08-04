package statuses

import (
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `Delete /v1/statuses/{id}`
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sid := r.FormValue("id")
	Sid, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	Account_auth := auth.AccountOf(r)

	statusRepo := h.app.Dao.Status()
	err = statusRepo.DeleteStatus(ctx, Sid, Account_auth)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

}
