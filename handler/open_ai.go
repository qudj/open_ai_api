package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qudj/open_ai_api/model"
	"github.com/qudj/open_ai_api/responses"
	"github.com/qudj/open_ai_api/service"
)

func openAiHandler(c *gin.Context) {
	var req model.OpenAIChatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.InvalidArgument.Resp(c)
		return
	}
	param := &model.HanHaiRequest{
		Query:          "1+1等于几",
		ResponseMode:   "dd",
		ConversationID: "",
		User:           "common",
	}
	dataChan := make(chan string)
	errChan := make(chan error)
	exitChan := make(chan bool)
	go service.OAIClient.StreamOpenAI(c, param, dataChan, errChan, exitChan)
	var quite bool
	for !quite {
		select {
		case err := <-errChan:
			responses.RespFromError(c, err)
		case data := <-dataChan:
			transData, err := TransResponse(param.ResponseMode, data)
			if err != nil {
				responses.RespFromError(c, err)
				break
			}
			responses.Success.RespData(c, transData)
			c.Writer.Flush()
		case <-exitChan:
			quite = true
		}
	}
	close(dataChan)
	close(errChan)
	close(exitChan)
}

func TransResponse(responseMode string, data string) (interface{}, error) {
	switch responseMode {
	case "dd":

	}
	return nil, errors.New("transData error")
}
