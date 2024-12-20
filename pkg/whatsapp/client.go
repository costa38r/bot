// RunClient sets up, connects, and manages the WhatsApp client lifecycle.
package whatsapp

import "fmt"

func RunClient() error {
	container, err := ConfigContainer()
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
