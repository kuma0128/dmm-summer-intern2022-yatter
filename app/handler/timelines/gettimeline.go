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

	_maxID := r.FormValue("_maxID")
	if _maxID == "" {
		_maxID = "0"
	}
	if strings.Contains(_maxID, "-") {
		httperror.BadRequest(w, errors.Errorf("negative ID doesn't existe"))
		return
	}
	maxID, err := strconv.ParseInt(_maxID, 10, 64)
	if maxID == 0 {
		maxID = math.MaxInt64
	}
	//fmt.Printf("%d\n", _maxID)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	_sinceID := r.FormValue("_sinceID")
	if _sinceID == "" {
		_sinceID = "1"
	}
	if strings.Contains(_sinceID, "-") {
		httperror.BadRequest(w, errors.Errorf("negative ID doesn't existe"))
		return
	}
	sinceID, err := strconv.ParseInt(_sinceID, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	//_maxID < _sinceID はエラー
	if maxID < sinceID {
		httperror.BadRequest(w, errors.Errorf("Need that _maxID is bigger than _sinceID"))
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
	statuses, err = timelineRepo.FindPublicTimelines(ctx, maxID, sinceID, Limit)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}
