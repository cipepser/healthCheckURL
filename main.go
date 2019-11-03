package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cipepser/healthCheckURL/slack"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context) (string, error) {
	c, err := slack.NewClient("アドベントカレンダー")
	if err != nil {
		panic(err)
	}

	return c.PostMessage("2019年公開されたよ！", "alert")
}
