package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/qudj/open_ai_api/model"
	"github.com/qudj/open_ai_api/responses"
)

func testHanHaiHandler(c *gin.Context) {
	var req *model.HanHaiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.InvalidArgument.Resp(c)
		return
	}
	if req.ResponseMode == "streaming" {
		handleStreamResponse(c, req)
	} else {
		handleBlockingResponse(c, req)
	}
}

func handleBlockingResponse(c *gin.Context, req *model.HanHaiRequest) {
	if req.ConversationID == "" {
		req.ConversationID = uuid.New().String()
	}
	ret := model.HanHaiCompleteResponse{
		ID:             uuid.New().String(),
		ConversationID: req.ConversationID,
		Answer:         "test block answer",
		CreatedAt:      time.Now().Unix(),
		Metadata: model.Metadata{
			Usage: model.Usage{
				PromptTokens:     12,
				CompletionTokens: 428,
				TotalTokens:      440,
			},
			RetrieverResources: []model.RetrieverResource{
				{
					DatasetID:     "101b4c97-fc2e-463c-90b1-5261a4cdcafb",
					DatasetName:   "助手知识库",
					DocumentID:    "8dd1ad74-0b5f-4175-b735-7d98bbbb4e00",
					DocumentName:  "助手文档001.pdf",
					Content:       "Hello world!",
					SequenceID:    "adb12dfj",
					SequenceIndex: 1,
				},
			},
		},
	}
	responses.Success.RespData(c, ret)
}

func handleStreamResponse(c *gin.Context, req *model.HanHaiRequest) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Status(http.StatusOK)
	for i := 0; i < 10; i++ {
		streamResponse := model.HanHaiSSEMessage{
			Event:          "message",
			ID:             uuid.New().String(),
			ConversationID: req.ConversationID,
			Answer:         fmt.Sprintf("Hello World! %d", i),
			Metadata: &model.Metadata{
				Usage: model.Usage{
					PromptTokens:     12,
					CompletionTokens: 428,
					TotalTokens:      440,
				},
				RetrieverResources: []model.RetrieverResource{
					{
						DatasetID:     "101b4c97-fc2e-463c-90b1-5261a4cdcafb",
						DatasetName:   "助手知识库",
						DocumentID:    "8dd1ad74-0b5f-4175-b735-7d98bbbb4e00",
						DocumentName:  "助手文档001.pdf",
						Content:       "Hello world!",
						SequenceID:    "adb12dfj",
						SequenceIndex: 1,
					},
				},
			},
		}
		jsonData, _ := json.Marshal(streamResponse)
		_, _ = c.Writer.Write([]byte("data: " + string(jsonData) + "\n\n"))
		c.Writer.Flush()
		time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
	}
	_, _ = c.Writer.Write([]byte("data: [DONE]\n\n"))
	c.Writer.Flush()
}
