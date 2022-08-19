package accounts

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// Handle request for `GET /v1/accounts/{username}/follower`
func (h *handler) Follower(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")

	maxID := r.FormValue("max_id")
	if maxID == "" {
		maxID = "0"
	}
	if strings.Contains(maxID, "-") {
		httperror.BadRequest(w, errors.Errorf("negative ID doesn't existe"))
		return
	}
	MaxID, err := strconv.ParseInt(maxID, 10, 64)
	if MaxID == 0 {
		MaxID = math.MaxInt64
	}
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	sinceID := r.FormValue("since_id")
	if sinceID == "" {
		sinceID = "1"
	}
	if strings.Contains(sinceID, "-") {
		httperror.BadRequest(w, errors.Errorf("negative ID doesn't existe"))
		return
	}
	SinceID, err := strconv.ParseInt(sinceID, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	//max_id < since_id はエラー
	if MaxID < SinceID {
		httperror.BadRequest(w, errors.Errorf("Need that max_id is bigger than since_id"))
		return
	}

	limit := r.FormValue("limit")
	if limit == "" {
		limit = "40"
	}
	Limit, err := strconv.ParseInt(limit, 10, 64)
	//fmt.Printf("%d\n", Limit)
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
	accounts, err = accountRepo.FindFollowerByID(ctx, _account.ID, MaxID, SinceID, Limit)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
