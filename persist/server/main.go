package main

import (
	"crawler/crawler_distributed/rpcsupport"
	"demoCrawler/persist"
	"github.com/olivere/elastic"
)

func main() {
	err := ServerRpc(":1234", "datring_profile")
}

func ServerRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.NewClient(host, persist.ItemSaverService{
		Client:client,
		Index:index,
	})
}