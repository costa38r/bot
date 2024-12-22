package openaiclient

type Assistant struct {
    ID string `json:"id"`
}

type ThreadResponse struct {
    ID            string         `json:"id"`
    Object        string         `json:"object"`
    CreatedAt     int64          `json:"created_at"`
    Metadata      map[string]any `json:"metadata"`
    ToolResources map[string]any `json:"tool_resources"`
}

type Content struct {
    Type string `json:"type"`
    Text Text   `json:"text"`
}

type Text struct {
    Value       string   `json:"value"`
    Annotations []string `json:"annotations"`
}

type Metadata struct{}

type Message struct {
    ID          string    `json:"id"`
    Object      string    `json:"object"`
    CreatedAt   int64     `json:"created_at"`
    AssistantID *string   `json:"assistant_id"`
    ThreadID    string    `json:"thread_id"`
    RunID       *string   `json:"run_id"`
    Role        string    `json:"role"`
    Content     []Content `json:"content"`
    Attachments []string  `json:"attachments"`
    Metadata    Metadata  `json:"metadata"`
}

type MessageListResponse struct {
    Object  string    `json:"object"`
    Data    []Message `json:"data"`
    FirstID string    `json:"first_id"`
    LastID  string    `json:"last_id"`
    HasMore bool      `json:"has_more"`
}

type CreateRunResponse struct {
    ID                string      `json:"id"`
    Object            string      `json:"object"`
    CreatedAt         int64       `json:"created_at"`
    AssistantID       string      `json:"assistant_id"`
    ThreadID          string      `json:"thread_id"`
    Status            string      `json:"status"`
    StartedAt         int64       `json:"started_at"`
    ExpiresAt         *int64      `json:"expires_at"`
    CancelledAt       *int64      `json:"cancelled_at"`
    FailedAt          *int64      `json:"failed_at"`
    CompletedAt       int64       `json:"completed_at"`
    LastError         *int64      `json:"last_error"`
    Model             string      `json:"model"`
    Instructions      *string     `json:"instructions"`
    IncompleteDetails *string     `json:"incomplete_details"`
    Tools             []Tool      `json:"tools"`
    Metadata          interface{} `json:"metadata"`
    Usage             *Usage      `json:"usage"`
    Temperature       float64     `json:"temperature"`
    TopP              float64     `json:"top_p"`
    MaxPromptTokens   int         `json:"max_prompt_tokens"`
    MaxCompletionTokens int       `json:"max_completion_tokens"`
    TruncationStrategy TruncationStrategy `json:"truncation_strategy"`
    ResponseFormat    interface{} `json:"response_format"`
    ToolChoice        string      `json:"tool_choice"`
    ParallelToolCalls bool        `json:"parallel_tool_calls"`
}
type RunResponse struct {
    ID                string      `json:"id"`
    Object            string      `json:"object"`
    CreatedAt         int64       `json:"created_at"`
    AssistantID       string      `json:"assistant_id"`
    ThreadID          string      `json:"thread_id"`
    Status            string      `json:"status"`
    StartedAt         int64       `json:"started_at"`
    ExpiresAt         *int64      `json:"expires_at"`
    CancelledAt       *int64      `json:"cancelled_at"`
    FailedAt          *int64      `json:"failed_at"`
    CompletedAt       int64       `json:"completed_at"`
    LastError         *int64      `json:"last_error"`
    Model             string      `json:"model"`
    Instructions      *string     `json:"instructions"`
    IncompleteDetails *string     `json:"incomplete_details"`
    Tools             []Tool      `json:"tools"`
    Metadata          interface{} `json:"metadata"`
    Usage             Usage       `json:"usage"`
    Temperature       float64     `json:"temperature"`
    TopP              float64     `json:"top_p"`
    MaxPromptTokens   int         `json:"max_prompt_tokens"`
    MaxCompletionTokens int       `json:"max_completion_tokens"`
    TruncationStrategy TruncationStrategy `json:"truncation_strategy"`
    ResponseFormat    interface{} `json:"response_format"`
    ToolChoice        string      `json:"tool_choice"`
    ParallelToolCalls bool        `json:"parallel_tool_calls"`
}
type Tool struct {
    Type string `json:"type"`
}

type Usage struct {
    PromptTokens     int `json:"prompt_tokens"`
    CompletionTokens int `json:"completion_tokens"`
    TotalTokens      int `json:"total_tokens"`
}

type TruncationStrategy struct {
    Type         string `json:"type"`
    LastMessages *string `json:"last_messages"`
}

type OpenAIClient struct {
    APIKey     string
    HTTPClient HttpClient
}