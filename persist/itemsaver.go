package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item saver: Got item " + "#%d: %v", itemCount, item)

			_, err := save(item)
			if err != nil {
				log.Printf("Item saver: error " + "saveing item %v : %v", item, err)
				continue
			}

		}
	}()

	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) // 内网
	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenhui").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	fmt.Printf("%v", resp)
	return resp.Id, nil
}