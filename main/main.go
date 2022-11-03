package main

import (
	"newsapi"
)

func main() {
	client := newsapi.NewClient(newsapi.ApiKey, newsapi.Agent)
	params := &newsapi.EverythingRequest{
		Q: "Golang",
	}
	client.Everything(params)
}
