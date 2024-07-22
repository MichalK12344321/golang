package server

import (
	"encoding/json"
	"io/ioutil"
	"lca/internal/pkg/util"
	"net/http"
)

type ControllerBase struct {
}

func (c *ControllerBase) WriteResponse(writer http.ResponseWriter, result any, err error, contentType string) error {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
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

func (c *ControllerBase) ParseBody(writer http.ResponseWriter, request *http.Request, target any) error {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	if err != nil {
		c.WriteResponse(writer, "Error reading request body", err, "application/json")
		return err
	}

	err = util.FromJson(target, bodyBytes)
	if err != nil {
		c.WriteResponse(writer, "Error parsing body", err, "application/json")
		return err
	}

	return nil
}
