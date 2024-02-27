package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/roblesdotdev/movies-ms/rating/internal/controller/rating"
	"github.com/roblesdotdev/movies-ms/rating/pkg/model"
)

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl}
}

// Handle PUT and GET /rating requests.
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	recordId := model.RecordId(r.FormValue("id"))
	if recordId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordType := model.RecordType(r.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(r.Context(), recordId, recordType)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	case http.MethodPut:
		userId := model.UserId(r.FormValue("userId"))
		v, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := h.ctrl.PutRating(r.Context(), recordId, recordType, &model.Rating{
			UserId: userId,
			Value:  model.RatingValue(v),
		}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
