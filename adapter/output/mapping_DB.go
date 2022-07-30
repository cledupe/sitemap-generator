package output

import (
	"sync"
)

type MappingDB struct {
	dbLinks map[string]int
	monitor sync.Mutex
}

func (mappingDB *MappingDB) Save(link string) {
	mappingDB.monitor.Lock()
	defer mappingDB.monitor.Unlock()
	if mappingDB.dbLinks == nil {
		mappingDB.dbLinks = make(map[string]int)
	}
	if mappingDB.dbLinks[link] != 1 {
		mappingDB.dbLinks[link] = 1
	}

}

func (mappingDB MappingDB) FindAll() []string {
	var arrayLinks []string

	for link := range mappingDB.dbLinks {
		arrayLinks = append(arrayLinks, link)
	}

	return arrayLinks
}
