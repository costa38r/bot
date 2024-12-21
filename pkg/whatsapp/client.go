// RunClient sets up, connects, and manages the WhatsApp client lifecycle.
package whatsapp

import (
	"fmt"

	"github.com/costa38r/bot/config"
)

func RunClient() error {
	cfg := config.LoadConfig()
	container, err := ConfigContainer(cfg)
	if err != nil {
		return fmt.Errorf("error configuring container: %w", err)
	}
	client, err := ConfigClient(container)
	if err != nil {
		return fmt.Errorf("error configuring client: %w", err)
	}
	if err := client.ConnectClient(); err != nil {
		return fmt.Errorf("error connecting client: %w", err)
	}
	waitForShutdown(client)
	return nil
}
