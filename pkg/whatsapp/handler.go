package whatsapp

import (
	"fmt"
	"strings"

	"github.com/costa38r/bot/pkg/threadcache"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type Sender struct{
	*types.JID
}


func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		sender, err := GetContact(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		sender.HanlderMessage()
	}
}

func GetContact(m *events.Message ) (*Sender,error) {

	senderJID := m.Info.Sender.User
	if senderJID == "" {
		return nil, fmt.Errorf("senderJID is empty")
	}
	contactNumber := strings.Split(senderJID, "@")[0]
	if contactNumber == "" {
		return nil, fmt.Errorf("contactNumber is empty")
	}

	return &Sender{&types.JID{User: contactNumber}}, nil

}

func (s *Sender) HanlderMessage() error {

	rdb,err := threadcache.NewRedisClient()
	if err != nil {
		fmt.Println("error creating redis client: ", err)
	}
	
	exists, err := threadcache.CheckValueExists(rdb, s.User)
	if err != nil {
		fmt.Println("error checking value in redis: ", err)

	}

	print(exists)

	return nil
}