package whatsapp

import (
	"context"
	"fmt"

	"github.com/costa38r/bot/pkg/openaiclient"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func (c *Client) eventHandler(evt interface{}) {
    switch v := evt.(type) {
    case *events.Message:
        c.HandlerMessage(context.Background(), v)
    }
}

func (c *Client) HandlerMessage(ctx context.Context, v *events.Message) {
    contactMsg := v.RawMessage.GetConversation()

    if contactMsg != "" {
        response, err := openaiclient.OpenAiHandle(ctx, contactMsg)
        if err != nil {
            fmt.Printf("Error handling message: %v\n", err)
            return
        }

        senderJID := types.NewJID(v.Info.Sender.User, types.DefaultUserServer)

        message := &waProto.Message{
            Conversation: proto.String(response),
        }

        if message != nil {
            if _, err = c.SendMessage(ctx, senderJID, message); err != nil {
                fmt.Printf("Error sending message: %v\n", err)
            }
        }
    }
}