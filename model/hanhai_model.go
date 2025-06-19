package model

type HanHaiRequest struct {
	Query          string `json:"query"`
	ResponseMode   string `json:"response_mode"`
	ConversationID string `json:"conversation_id"`
	User           string `json:"user"`
}

type RetrieverResource struct {
	DatasetID     string `json:"dataset_id"`
	DatasetName   string `json:"dataset_name"`
	DocumentID    string `json:"document_id"`
	DocumentName  string `json:"document_name"`
	Content       string `json:"content"`
	SequenceID    string `json:"sequence_id"`
	SequenceIndex int    `json:"sequence_index"`
}

type Metadata struct {
	Usage              Usage               `json:"usage"`
	RetrieverResources []RetrieverResource `json:"retriever_resources,omitempty"` // 可能不存在
}

type HanHaiCompleteResponse struct {
	ID             string   `json:"id"`
	ConversationID string   `json:"conversation_id"`
	Answer         string   `json:"answer"`
	Reason         string   `json:"reason"`
	CreatedAt      int64    `json:"created_at"`
	Metadata       Metadata `json:"metadata"`
}

type HanHaiSSEMessage struct {
	Event          string    `json:"event"` // "message" 或 "message_end"
	ID             string    `json:"id"`
	ConversationID string    `json:"conversation_id"`
	Answer         string    `json:"answer"`
	Reason         string    `json:"reason"`
	Metadata       *Metadata `json:"metadata,omitempty"` // 仅 message_end 带，有时省略
}
