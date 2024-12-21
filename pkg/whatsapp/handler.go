package whatsapp

import (
	"context"

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
	//senderJID := m.Info.Sender.User
	//contactNumber := strings.Split(senderJID, "@")[0]

	ctx := context.Background()

	cfg := &threadcache.RedisClientConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	threadcache.CheckIfThreadExists(ctx, cfg, "thread1", "this is a test value")

}
