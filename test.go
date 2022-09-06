package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://azurecomcdn.azureedge.net/en-us/updates/feed/")
	for i := 0; i < len(feed.Items); i++ {
		fmt.Println(feed.Items[i].Title)
	}
}
