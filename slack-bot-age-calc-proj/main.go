package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-303145360071-3816026909282-4Bc7bT0tMEYNLKF6jrdXblJ7")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03PX1JDM53-3839765481424-c67698451e5f531dd11ca0f99d6bd3ce4c434aaaef44171013cbed283f099dc7")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN", os.Getenv("SLACK_APP_TOKEN")))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition){
		Description: "yob Calculator",
		Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2022-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
