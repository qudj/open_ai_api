package service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/qudj/open_ai_api/model"
	"io"
	"net/http"
	"sync"
)

type OpenAIClient struct {
	host         string
	client       *http.Client
	tokens       []string
	selIndex     int
	tokenCount   int
	nextTokenGap int
	reqCount     int
	lock         *sync.Mutex
}

func (c *OpenAIClient) StreamOpenAI(ctx context.Context, param *model.HanHaiRequest, dataChan chan string, errChan chan error) {
	defer close(dataChan)
	defer close(errChan)
	openaiAPIKey := c.getReqToken()
	reqBody, _ := json.Marshal(param)
	req, err := http.NewRequestWithContext(ctx, "POST", c.host, bytes.NewReader(reqBody))
	if err != nil {
		errChan <- err
		return
	}
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	if param.ResponseMode == "streaming" {
		buf := make([]byte, 1024)
		for {
			n, err := resp.Body.Read(buf)
			if n > 0 {
				dataChan <- string(buf[:n])
			}
			if err != nil {
				if err != io.EOF {
					errChan <- err
				}
				break
			}
		}
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}
		dataChan <- string(body) // 非流式只发一次
	}
}

func (c *OpenAIClient) getReqToken() string {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.reqCount++
	if c.reqCount >= c.nextTokenGap {
		c.selIndex++
		c.reqCount = 0
	}
	if c.selIndex >= c.tokenCount {
		c.selIndex = 0
	}
	return c.tokens[c.selIndex]
}

func (c *OpenAIClient) nextToken() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.selIndex++
	c.reqCount = 0
	if c.selIndex >= c.tokenCount {
		c.selIndex = 0
	}
}
