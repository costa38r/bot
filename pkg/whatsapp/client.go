package whatsapp

import (
	"context"
	"fmt"
)

func RunClient(ctx context.Context) error {

    container, err := ConfigContainer()
    if err != nil {
        return fmt.Errorf("error configuring container: %w", err)
    }
    client, err := ConfigClient(container)
    if err != nil {
        return fmt.Errorf("error configuring client: %w", err)
    }
    if err := client.ConnectClient(ctx); err != nil {
        return fmt.Errorf("error connecting client: %w", err)
    }
    waitForShutdown(client)
    return nil
}