package handler

import (
	"fmt"
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
	go service.OAIClient.StreamOpenAI(c, param, dataChan, errChan)

	for {
		select {
		case err := <-errChan:
			fmt.Println(err)
			break
		case data := <-dataChan:
			fmt.Println(data)
			break
		}
	}
}
