package server

import (
	"encoding/json"
	"net/http"
)

type ControllerBase struct {
}

func (c *ControllerBase) WriteResponse(writer http.ResponseWriter, result any, err error, contentType string) error {
	writer.Header().Add("Content-Type", contentType)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, err = writer.Write([]byte(err.Error()))
		return err
	}

	writer.WriteHeader(http.StatusOK)
	var jsonResult []byte
	jsonResult, _ = json.MarshalIndent(result, "", "  ")
	_, err = writer.Write([]byte(jsonResult))
	return err
}
