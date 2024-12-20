package main

import (
	"fmt"

	"github.com/costa38r/bot/pkg/whatsapp"
)

func main() {
	container, err := whatsapp.ConfigContainer()
	if err != nil {
		fmt.Println("Error configuring container:", err)
		return
	}

	client, err := whatsapp.ConfigClient(container)
	if err != nil {
		fmt.Println("Error configuring client:", err)
		return
	}

	err = client.RunClient()
	if err != nil {
		fmt.Println("Error running client:", err)
	}
}
