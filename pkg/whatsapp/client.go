package whatsapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

// Client wraps the WhatsApp client to add custom methods.
type Client struct {
	*whatsmeow.Client
}

// Container wraps the SQL store container for custom usage.
type Container struct {
	*sqlstore.Container
}

// ConfigContainer sets up the SQL store container for WhatsApp data persistence.
func ConfigContainer() (*Container, error) {
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", waLog.Stdout("Database", "DEBUG", true))
	if err != nil {
		return nil, fmt.Errorf("failed to create storage container: %w", err)
	}
	return &Container{Container: container}, nil
}

// ConfigClient initializes a WhatsApp client using the provided container.
func ConfigClient(container *Container) (*Client, error) {
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return nil, fmt.Errorf("failed to get device store: %w", err)
	}
	client := whatsmeow.NewClient(deviceStore, nil)
	client.AddEventHandler(eventHandler)
	return &Client{Client: client}, nil
}

// ConnectClient connects the WhatsApp client, generating a QR code if not logged in.
func (c *Client) ConnectClient() error {
	if c.Store.ID == nil {
		qrChan, _ := c.GetQRChannel(context.Background())
		fmt.Println("Generating QR code...")
		if err := c.Connect(); err != nil {
			return fmt.Errorf("failed to connect client: %w", err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login Event:", evt.Event)
			}
		}
	} else {
		if err := c.Connect(); err != nil {
			return fmt.Errorf("failed to connect client: %w", err)
		}
		fmt.Println("Client connected successfully.")
	}
	return nil
}

// CloseClient disconnects the WhatsApp client.
func (c *Client) CloseClient() {
	fmt.Println("Disconnecting client...")
	c.Disconnect()
	fmt.Println("Client disconnected.")
}

// waitForShutdown waits for termination signals and disconnects the client.
func waitForShutdown(client *Client) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	client.CloseClient()
}

// RunClient sets up, connects, and manages the WhatsApp client lifecycle.
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