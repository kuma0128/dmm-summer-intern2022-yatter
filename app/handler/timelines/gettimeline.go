package timelines

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/pkg/errors"
)

// Handle request for `GET /v1/timelines/public`
func (h *handler) Gettimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	max_id := r.FormValue("max_id")
	if max_id == "" {
		max_id = "0"
	}
	if strings.Contains(max_id, "-") {
		httperror.BadRequest(w, errors.Errorf("negative ID doesn't existe"))
		return
	}
	Max_id, err := strconv.ParseInt(max_id, 10, 64)
	if Max_id == 0 {
		Max_id = math.MaxInt64
	}
	//fmt.Printf("%d\n", Max_id)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	since_id := r.FormValue("since_id")
	if since_id == "" {
		since_id = "1"
	}
	if strings.Contains(since_id, "-") {
		httperror.BadRequest(w, errors.Errorf("negative ID doesn't existe"))
		return
	}
	Since_id, err := strconv.ParseInt(since_id, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	//max_id < since_id はエラー
	if Max_id < Since_id {
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

	var statuses []*object.Status

	timelineRepo := h.app.Dao.Timeline()
	statuses, err = timelineRepo.FindPublicTimelines(ctx, Max_id, Since_id, Limit)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}
