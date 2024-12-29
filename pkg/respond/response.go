package respond

import (
	"context"
	"encoding/json"
	"github.com/imabg/responehq/pkg/logger"
	"net/http"
)

func GetBody(ctx context.Context, r *http.Request, param interface{}) error {
	err := json.NewDecoder(r.Body).Decode(param)
	if err != nil {
		logger.Error(ctx, "Failed to decode body", err)
		return err
	}
	return nil
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
func StatusInternalServerError(ctx context.Context, w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	logger.Error(ctx, "error", response)
	databytes, err := json.Marshal(response)
	if err != nil {
		logger.Error(ctx, "Failed to encode response", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	_, err = w.Write(databytes)
	if err != nil {
		logger.Error(ctx, "Failed to encode response: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// 401 - Unauthorised
