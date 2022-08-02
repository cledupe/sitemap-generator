package output_test

import (
	"fmt"
	"github.com/cledupe/sitemap-generator/adapter/output"
	"math/rand"
	"sync"
	"testing"
)

const ConcurrencyValue = 100

func TestInsertAndLoad(t *testing.T) {
	db := output.MappingDB{}
	db.Save("line 1")
	db.Save("line 2")
	array := db.FindAll()

	if len(array) != 2 {
		t.Errorf("Wrong Size array should be %d but is %d", 2, len(array))
	}
}

func TestInsertInConcurrency(t *testing.T) {
	db := output.MappingDB{}
	wg := &sync.WaitGroup{}
	wg.Add(ConcurrencyValue)
	for i := 0; i < ConcurrencyValue; i++ {
		go func(w *sync.WaitGroup) {
			value := fmt.Sprintf("%d", rand.Int())
			db.Save(value)
			wg.Done()
		}(wg)
	}
	wg.Wait()

	arraySize := len(db.FindAll())
	if arraySize != ConcurrencyValue {
		t.Errorf("Wrong size array with length %d but should be %d", arraySize, ConcurrencyValue)
	}
}
