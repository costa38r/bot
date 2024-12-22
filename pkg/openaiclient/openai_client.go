package openaiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// NewOpenAIClient cria uma nova instância de OpenAIClient.
func NewOpenAIClient(apiKey string) *OpenAIClient {
    return &OpenAIClient{
        APIKey:     apiKey,
        HTTPClient: &HttpClientImpl{Client: &http.Client{}},
    }
}

// CreateThread cria uma nova thread.
func (c *OpenAIClient) CreateThread() (*ThreadResponse, error) {
    url := "https://api.openai.com/v1/threads"

    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(``)))
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

// AddMsgToThread adiciona uma mensagem a uma thread.
func (c *OpenAIClient) AddMsgToThread(threadID, content string) (*Message, error) {
    url := fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", threadID)

    message := map[string]string{
        "role":    "user",
        "content": content,
    }
    messageJSON, err := json.Marshal(message)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal message: %w", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(messageJSON))
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

// CreateRun cria uma nova execução.
func (c *OpenAIClient) CreateRun(threadID, assistantID string) (*CreateRunResponse, error) {
    url := fmt.Sprintf("https://api.openai.com/v1/threads/%s/runs", threadID)
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

// GetRunResponse obtém a resposta de uma execução.
func (c *OpenAIClient) GetRunResponse(threadID, runID string) (*RunResponse, error) {
    url := fmt.Sprintf("https://api.openai.com/v1/threads/%s/runs/%s", threadID, runID)
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

// GetMessages obtém as mensagens de uma thread.
func (c *OpenAIClient) GetMessages(threadID string) (*MessageListResponse, error) {
    url := fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", threadID)
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

// GetAssistantMessages filtra as mensagens do assistente.
func (c *OpenAIClient) GetAssistantMessages(messages []Message) []Message {
    var assistantMessages []Message
    for _, message := range messages {
        if message.Role == "assistant" {
            assistantMessages = append(assistantMessages, message)
        }
    }
    return assistantMessages
}

// setHeaders define os cabeçalhos comuns para as requisições.
func (c *OpenAIClient) setHeaders(req *http.Request) {
    req.Header.Set("Authorization", "Bearer "+c.APIKey)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("OpenAI-Beta", "assistants=v2")
}