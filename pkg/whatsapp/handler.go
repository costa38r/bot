package whatsapp

import (
	"fmt"
	"strings"

	"github.com/costa38r/bot/pkg/threadcache"
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
	threadcache.CheckIfThreadExists()
	fmt.Println("Número do contato:", contactNumber)
}
