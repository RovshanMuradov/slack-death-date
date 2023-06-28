package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
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
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-1501102524275-3300719242437-MCJRDAIhh7XV71NikR5aTIqc")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A039MA5N72L-3305927315140-9216e2c4ab14d3b2dc97028f08e95d780583f31f12085e1e8d453e97d69438ea")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my age is <age>", &slacker.CommandDefinition{
		Description: "my date of death",
		Example:     "my date of death is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			age := request.Param("age")
			date, err := strconv.Atoi(age)
			if err != nil {
				fmt.Println("error")
			}

			seconds := time.Now().Unix()
			rand.Seed(seconds)
			target := rand.Intn(50) + 1

			dateOfDeath := date + target

			r := fmt.Sprintf("your date of death is %d", dateOfDeath)
			response.Reply(r)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
