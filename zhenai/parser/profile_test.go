package parser

import (
	"demoCrawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "安静的雪")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+ "element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age:34,
		Marriage:"离异",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", )
	}
}
