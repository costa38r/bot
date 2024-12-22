// The main package initializes and starts the WhatsApp client.
package main

import (
	"github.com/costa38r/bot/pkg/openaiclient"
)


func main() {
	// RunClient handles the entire lifecycle of the WhatsApp client.
	// If an error occurs, it will be logged to the console.
	/*err := whatsapp.RunClient()
	if err != nil {
		fmt.Println("error running client: ", err)
	}*/
	
	openaiclient.OpenAiHandle()
	
}


