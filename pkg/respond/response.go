package respond

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/imabg/responehq/pkg/errors"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/pkg/validate"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetBody(r *http.Request, param interface{}) error {
	err := json.NewDecoder(r.Body).Decode(param)
	if err != nil {
		return err
	}
	// check validation
	err = validate.Struct(param)
	return err
}

func Send(ctx context.Context, w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	databytes, err := json.Marshal(response.Data)
	if err != nil {
		logger.Error(ctx, "Failed to encode response: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	_, err = w.Write(databytes)
	if err != nil {
		logger.Error(ctx, "Failed to encode response: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}

func SendWithError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	var ex []errors.ResponseError
	switch err.(type) {
	case *errors.Error:
		if e, ok := err.(*errors.Error); ok {
			w.WriteHeader(e.Code)
			rE := errors.ResponseError{
				Code:        e.Code,
				Description: e.Message,
				Type:        e.Type,
			}
			ex = append(ex, rE)
			eList := errors.ResponseErrorArr{ex}
			json.NewEncoder(w).Encode(eList)
			return
		}
	case validator.ValidationErrors:
		if e, ok := err.(validator.ValidationErrors); ok {
			w.WriteHeader(http.StatusBadRequest)
			rE := errors.ResponseError{
				Code:        http.StatusBadRequest,
				Type:        errors.VALIDATION_ERROR,
				Description: e.Error(),
			}
			json.NewEncoder(w).Encode(rE)
			return
		}
	}
	return
}
