package openaiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/costa38r/bot/config"
)


const (
    createThreadParam     = "/threads"
    AddMsgToThreadParam   = "/threads/%s/messages"
    CreateRunParam        = "/threads/%s/runs"
    GetRunResponseParam   = "/threads/%s/runs/%s"
    GetMessagesParam      = "/threads/%s/messages"
)

func NewOpenAIClient(apiKey string) *OpenAIClient {
    return &OpenAIClient{
        APIKey:     apiKey,
        HTTPClient: &HttpClientImpl{Client: &http.Client{}},
    }
}

func (c *OpenAIClient) CreateThread(ctx context.Context) (*ThreadResponse, error) {

    cfg := config.GetConfig()
    url := cfg.OpenAiConfig.URLBase + createThreadParam

    req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer([]byte(``)))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    c.setHeaders(req)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to create thread: %s", resp.Status)
    }

    var threadResponse ThreadResponse
    if err := json.NewDecoder(resp.Body).Decode(&threadResponse); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    fmt.Println("Thread created successfully with ID:", threadResponse.ID)
    return &threadResponse, nil
}

func (c *OpenAIClient) AddMsgToThread(ctx context.Context, threadID, content string) (*Message, error) {
    cfg := config.GetConfig()
    url := fmt.Sprintf(cfg.OpenAiConfig.URLBase+AddMsgToThreadParam, threadID)

    message := map[string]string{
        "role":    "user",
        "content": content,
    }

    messageJSON, err := json.Marshal(message)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal message: %w", err)
    }

    req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(messageJSON))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    c.setHeaders(req)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to add message to thread: %s", resp.Status)
    }

    var messageResponse Message
    if err := json.NewDecoder(resp.Body).Decode(&messageResponse); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    fmt.Println("Message added successfully with ID:", messageResponse.ID)
    return &messageResponse, nil
}

func (c *OpenAIClient) CreateRun(threadID, assistantID string) (*CreateRunResponse, error) {
    cfg := config.GetConfig()
    url := fmt.Sprintf(cfg.OpenAiConfig.URLBase+CreateRunParam, threadID)
    payload := map[string]string{
        "assistant_id": assistantID,
    }
    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal payload: %w", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    c.setHeaders(req)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    var createRunResponse CreateRunResponse
    if err := json.NewDecoder(resp.Body).Decode(&createRunResponse); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    return &createRunResponse, nil
}

func (c *OpenAIClient) GetRunResponse(threadID, runID string) (*RunResponse, error) {
    cfg:= config.GetConfig()
    url := fmt.Sprintf(cfg.OpenAiConfig.URLBase+GetRunResponseParam, threadID, runID)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    c.setHeaders(req)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    var runResponse RunResponse
    if err := json.NewDecoder(resp.Body).Decode(&runResponse); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    return &runResponse, nil
}

func (c *OpenAIClient) GetMessages(threadID string) (*MessageListResponse, error) {
    cfg := config.GetConfig()
    url := fmt.Sprintf(cfg.OpenAiConfig.URLBase+GetMessagesParam, threadID)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    c.setHeaders(req)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    var messageListResponse MessageListResponse
    if err := json.NewDecoder(resp.Body).Decode(&messageListResponse); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    return &messageListResponse, nil
}

func (c *OpenAIClient) GetAssistantMessages(messages []Message) []Message {
    var assistantMessages []Message
    for _, message := range messages {
        if message.Role == "assistant" {
            assistantMessages = append(assistantMessages, message)
        }
    }
    return assistantMessages
}

func (c *OpenAIClient) setHeaders(req *http.Request) {
    req.Header.Set("Authorization", "Bearer "+c.APIKey)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("OpenAI-Beta", "assistants=v2")
}