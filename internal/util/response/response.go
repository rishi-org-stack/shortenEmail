package response

import (
	"encoding/json"
	"net/http"
	utilError "shortenEmail/internal/util/error"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Respond(w http.ResponseWriter, res *Response) (err error) {

	w.WriteHeader(res.Status)
	w.Header().Set("Content-Type", "application/json")
	jsonRes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	w.Write(jsonRes)

	return nil
}
func RespondError(w http.ResponseWriter, res utilError.ApiErrorInterface) (err error) {
	cres := res.(utilError.ApiError)
	w.WriteHeader(cres.Status)
	w.Header().Set("Content-Type", "application/json")
	jsoncres, err := json.Marshal(cres)
	if err != nil {
		return err
	}
	w.Write(jsoncres)

	return nil
}
