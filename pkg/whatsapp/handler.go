package whatsapp

import (
	"context"
	"fmt"

	"github.com/costa38r/bot/pkg/openaiclient"
	"github.com/costa38r/bot/pkg/threadcache"
	"go.mau.fi/whatsmeow/types/events"
)

func (c *Client) eventHandler(evt interface{}) {

    
    switch v := evt.(type) {

    case *events.Message:
        c.HandlerMessage(context.Background(), v)
    }
   

}

func (c *Client) HandlerMessage(ctx context.Context, v *events.Message) {

    contactNumber := v.Info.Chat.User
    contactMsg := v.RawMessage.GetConversation()
    if contactMsg == "" {
        return
    }

    rdb,err := threadcache.NewRedisClient(ctx)
    if err != nil {
        return
    }

    err = CheckIfThreadExists(ctx,rdb,contactNumber)
    if err != nil {
        fmt.Println("Thread not found")
         respAssitant, err:= openaiclient.OpenAiHandle(ctx,contactMsg )
         if err != nil {
                fmt.Println("Error creating thread")
                return
            }
        fmt.Println(respAssitant.ThreadID)
        CreateThreadOnCache(ctx,rdb,contactNumber,respAssitant.ThreadID)
    }
}

func CheckIfThreadExists(ctx context.Context, rdb *threadcache.RedisClient, contactNumber string) error {
    _,err := rdb.GetData(ctx,rdb,contactNumber)
    if err != nil {
      return err
    }
    return err
}

func CreateThreadOnCache(ctx context.Context, rdb *threadcache.RedisClient, contactNumber string,threadID string) error {
    err := rdb.StoreData(ctx,rdb,contactNumber,threadID)
    if err != nil {
        return err
    }
    return nil
}