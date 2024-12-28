package respond

import (
	"context"
	"encoding/json"
	"github.com/imabg/responehq/pkg/logger"
	"net/http"
)

func GetBody(ctx context.Context, w http.ResponseWriter, param interface{}) {
	err := json.NewEncoder(w).Encode(param)
	if err != nil {
		logger.Error(ctx, "Failed to encode body: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// 200 - StatusOk
func StatusOk(ctx context.Context, w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	databytes, err := json.Marshal(response)
	if err != nil {
		logger.Error(ctx, "Failed to encode response: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	_, err = w.Write(databytes)
	if err != nil {
		logger.Error(ctx, "Failed to encode response: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// 400 - BadRequest
// 500 - Internal server error
// 401 - Unauthorised
