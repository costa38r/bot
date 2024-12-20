package whatsapp

import (
	"fmt"
	"strings"

	"go.mau.fi/whatsmeow/types/events"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		GetContact(v)
	}
}

func GetContact(m *events.Message) {
	senderJID := m.Info.Sender.User
	contactNumber := strings.Split(senderJID, "@")[0]
	fmt.Println("Número do contato:", contactNumber)
}
