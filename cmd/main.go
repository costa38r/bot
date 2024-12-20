package main

import (
	"fmt"

	"github.com/costa38r/bot/pkg/whatsapp"
)

func main() {

	err := whatsapp.RunClient()
	if err != nil {
		fmt.Println("error running client: ", err)

	}
}
