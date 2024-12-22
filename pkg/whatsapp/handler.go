package whatsapp

import (
	"fmt"
	"time"

	"go.mau.fi/whatsmeow/types/events"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		HandlerMessage(v)
	}
}

func HandlerMessage(v *events.Message) {
	contactMsg := v.RawMessage.GetConversation()
	fmt.Println(contactMsg)
	time.Sleep(1 * time.Second) // Dá tempo para a goroutine terminar

}