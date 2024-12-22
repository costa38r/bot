package openaiclient

import (
	"context"
	"fmt"
	"time"

	"github.com/costa38r/bot/config"
)

func OpenAiHandle(ctx context.Context, contactMsg string) (*Message, error) {
    cfg := config.GetConfig()
    client := NewOpenAIClient(cfg.OpenAiConfig.APIKey)

    thread, err := client.CreateThread(ctx)
    if err != nil {
        return nil, fmt.Errorf("error creating thread: %w", err)
    }

    _, err = client.AddMsgToThread(ctx, thread.ID, contactMsg)
    if err != nil {
        return nil, fmt.Errorf("error adding message to thread: %w", err)
    }

    run, err := client.CreateRun(thread.ID, cfg.OpenAiConfig.AssistantID)
    if err != nil {
        return nil, fmt.Errorf("error creating run: %w", err)
    }

    time.Sleep(3 * time.Second)
    runResponse, err := client.GetRunResponse(thread.ID, run.ID)
    if err != nil {
        return nil, fmt.Errorf("error getting run response: %w", err)
    }

    if runResponse.Status == "completed" {
        msg, err := client.GetMessages(thread.ID)
        if err != nil {
            return nil, fmt.Errorf("error getting messages: %w", err)
        }

        assistantMessages := client.GetAssistantMessages(msg.Data)
        if len(assistantMessages) == 0 {
            return nil, fmt.Errorf("no assistant messages found")
        }

       return &assistantMessages[0], nil
    }

    return nil, fmt.Errorf("run not completed")
}