package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	fmt.Println()
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3958606077490-3986362840240-22vmk1PHjnCY5wjkfOxWf4vw")
	os.Setenv("CHANNEL_ID", "C03TZV9Q4GN")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"python.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println("%s\n", err)
			return
		}
		fmt.Println("Name: %s,URL: %s\n", file.Name, file.URL)
	}
}
