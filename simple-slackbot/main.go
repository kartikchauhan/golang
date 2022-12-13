package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		log.Println(event.Timestamp)
		log.Println(event.Command)
		log.Println(event.Parameters)
		log.Println(event.Event)
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "<xoxb-token>")
	os.Setenv("SLACK_APP_TOKEN", "<xapp-token>")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	definition := &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 1990"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Fatal("Error while parsing")
			}

			r := fmt.Sprintf("Your age is %d", 2022-yob)
			response.Reply(r)
		},
	}

	bot.Command("my yob is <year>", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
