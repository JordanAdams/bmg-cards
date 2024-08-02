package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	api "github.com/jordanadams/bmg-cards/pkg"
)

var CLI struct {
	Download struct{} `cmd:"" help:"Download BMG card images"`
	Process  struct{} `cmd:"" help:"Process data from BMG cards"`
}

func downloadCommand() {
	resp, err := api.FetchGameData()
	if err != nil {
		panic(err)
	}

	var imageUrls []string
	for _, card := range resp.Cards {
		imageUrls = append(imageUrls, card.Image)
	}

	fmt.Println(imageUrls)
}

func main() {
	ctx := kong.Parse(&CLI)

	switch ctx.Command() {
	case "download":
		downloadCommand()
	case "process":
	default:
		panic(ctx.Command())
	}
}
