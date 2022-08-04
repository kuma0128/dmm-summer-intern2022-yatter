package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `GET /v1/timelines/public`
func (h *handler) Gettimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	max_id := chi.URLParam(r, "max_id")
	Max_id, err := strconv.ParseInt(max_id, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	since_id := chi.URLParam(r, "since_id")
	Since_id, err := strconv.ParseInt(since_id, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	limit := chi.URLParam(r, "limit")
	if limit == "" {
		limit = "40"
	}
	Limit, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	if 80 < Limit {
		Limit = 80
	}

	var statuses []object.Status

	timelineRepo := h.app.Dao.Timeline()
	statuses, err = timelineRepo.FindPublicTimelines(ctx, Max_id, Since_id, Limit, statuses)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}
