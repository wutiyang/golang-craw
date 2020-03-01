package persist

import (
	"demoCrawler/engine"
	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	id, err := save(s.Client, s.Index, item)

}
