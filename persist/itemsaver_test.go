package persist

import (
	"context"
	"demoCrawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSaver(t *testing.T) {

	profile := model.Profile{
		Age: 35,
		Height:163,
		Income: "3001-5000元",
		Name: "安静的雪2",
	}

	id, err := save(profile)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%v", resp)

	var actual model.Profile
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	if actual != profile {
		t.Errorf("no pass")
	}
}


