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
        c.HandlerMessage(v)
    }
}

func (c *Client) HandlerMessage(v *events.Message) {
    contactMsg := v.RawMessage.GetConversation()

    if contactMsg != "" {
        response, err := openaiclient.OpenAiHandle(contactMsg)
        if err != nil {
            fmt.Printf("Error handling message: %v\n", err)
            return
        }

        // Construir o JID completo do remetente original
        senderJID := types.NewJID(v.Info.Sender.User, types.DefaultUserServer)

        // Criar a mensagem de resposta
        message := &waProto.Message{
            Conversation: proto.String(response),
        }

        if message != nil {

        // Enviar a mensagem de resposta
        _, err = c.SendMessage(context.Background(), senderJID, message)
        if err != nil {
            fmt.Printf("Error sending message: %v\n", err)
        }
    } else {
        return
    }
    }
}