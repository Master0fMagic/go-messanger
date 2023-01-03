package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	contentTypeJSON = "application/json;charset=utf-8"
)

func sendResponse(ctx *gin.Context, statusCode int, respBody interface{}) {
	ctx.Writer.Header().Set("Content-Type", contentTypeJSON)

	binRespBody, err := json.Marshal(respBody)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	ctx.Writer.WriteHeader(statusCode)
	_, _ = ctx.Writer.Write(binRespBody)
}

func getRequestBody(ctx *gin.Context, obj any) error {
	reqBodyReader, err := ctx.Request.GetBody()
	if err != nil {
		log.Errorf("Error parsing request body: %+v", err)
		return err
	}

	data, err := ioutil.ReadAll(reqBodyReader)
	if err != nil {
		log.Errorf("Error reading request body: %+v", err)
		return err
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		log.Errorf("Error unmarshaling request body: %+v", err)
		return err
	}

	return nil
}
