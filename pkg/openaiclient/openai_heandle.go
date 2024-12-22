package openaiclient

import (
	"fmt"
	"os"
	"time"
)

func OpenAiHandle(contactMsg string)(string,error){
	client := NewOpenAIClient(os.Getenv("OPENAI_API_KEY"))

    thread, err := client.CreateThread()
    if err != nil {
        fmt.Println("error creating thread: ", err)
        return "", err
    }
    client.AddMsgToThread(thread.ID, contactMsg)
    run, err := client.CreateRun(thread.ID, os.Getenv("ASSISTANT_ID"))
    if err != nil {
        fmt.Println("error creating run: ", err)
        return "", err
    }

    time.Sleep(3 * time.Second)
    runResponse, err := client.GetRunResponse(thread.ID, run.ID)
    if err != nil {
        fmt.Println("error getting run response: ", err)
       return   "", err
    }

    if runResponse.Status == "completed" {
        msg, err := client.GetMessages(thread.ID)
        if err != nil {
            fmt.Println("error getting messages: ", err)
            return "", err
        }

        assistantMessages := client.GetAssistantMessages(msg.Data)
        for _, message := range assistantMessages {
            fmt.Println(message.Content[0].Text.Value)
        }

        mensagem :=  assistantMessages[len(assistantMessages)-1].Content[0].Text.Value
        time.Sleep(3*time.Second)
        return mensagem, nil

    }

    return "", nil
}