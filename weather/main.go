package main

import (
	"context"
	"fmt"
	"log"

	"github.com/notifoxhq/notifox-go"
)

func main() {
	ctx := context.Background()

	client, err := notifox.NewClientFromEnv()
	if err != nil {
		log.Panic(err)
	}

	weatherObject, err := GetSanFranciscoWeather(ctx)
	if err != nil {
		log.Println(err)
	}

	message := fmt.Sprintf("The weather in SF is currently: %v F", weatherObject.Current.Temperature2m)
	log.Println(message)

	if weatherObject.Current.Temperature2m >= 50 {
		resp, err := client.SendAlert(ctx, "mathis", message)
		if err != nil {
			log.Println(err)
		}

		log.Println(resp.Cost, resp.Currency, resp.MessageID)
	} else {
		log.Println("Weather is too cold, not sending a text message!")
	}

}
