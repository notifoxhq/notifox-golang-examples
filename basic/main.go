package main

import (
	"context"
	"fmt"

	"github.com/notifoxhq/notifox-go"
)

func main() {
	// Reads from NOTIFOX_API_KEY environment variable
	client, err := notifox.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	resp, err := client.SendAlert(ctx, "mike", "Database server is down!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Alert sent! Message ID: %s\n", resp.MessageID)
}
