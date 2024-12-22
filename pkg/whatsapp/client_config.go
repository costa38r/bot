package whatsapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/costa38r/bot/config"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)
type Client struct {
    *whatsmeow.Client
}

type Container struct {
    *sqlstore.Container
}


// ConfigContainer sets up the SQL store container for WhatsApp data persistence.
func ConfigContainer() (*Container, error) {
    cfg := config.GetConfig()
    dbLog := waLog.Stdout("Database", "DEBUG", true)
    container, err := sqlstore.New(cfg.WhatsApp.Dialect, cfg.WhatsApp.DSN, dbLog)
    if err != nil {
        return nil, fmt.Errorf("failed to create container: %w", err)
    }
    return &Container{Container: container}, nil
}


func ConfigClient(container *Container) (*Client, error) {
    deviceStore, err := container.GetFirstDevice()
    if err != nil {
        return nil, fmt.Errorf("failed to get device store: %w", err)
    }
    client := whatsmeow.NewClient(deviceStore, nil)
    client.AddEventHandler((&Client{Client: client}).eventHandler)
    return &Client{Client: client}, nil
}


func (c *Client) ConnectClient(ctx context.Context) error {
    if c.Store.ID == nil {
        qrChan, _ := c.GetQRChannel(ctx)
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


func (c *Client) CloseClient() {
    fmt.Println("Disconnecting client...")
    c.Disconnect()
    fmt.Println("Client disconnected.")
}

func waitForShutdown(client *Client) {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
    <-sigChan
    client.CloseClient()
}