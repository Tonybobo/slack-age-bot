package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"



	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main (){
	os.Setenv("SLACK_BOT_TOKEN" , "xoxb-4017635101060-4015209514514-AwJA3oXIkxLgzPzdYgqSengC")
	os.Setenv("SLACK_APP_TOKEN" , "xapp-1-A040C6BCMTP-4017640163764-6f56661010fd52b72a3f1598da0108bc26ec500b0d26b9480d57c262023b2cf6")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())


	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx , cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}