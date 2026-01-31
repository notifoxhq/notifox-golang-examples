package main

import (
	"context"
	"fmt"
	"log"

	"github.com/notifoxhq/notifox-go"
)

func main() {
	// Reads from NOTIFOX_API_KEY environment variable
	client, err := notifox.NewClient()
	if err != nil {
		log.Println(err)
		return
	}

	ctx := context.Background()
	resp, err := client.SendAlert(ctx, notifox.AlertRequest{
		Alert:    "Database server is down!",
		Channel:  notifox.Email, // or notifox.SMS
		Audience: "mathis",
	})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Alert sent! Message ID: %s\n", resp.MessageID)
}
