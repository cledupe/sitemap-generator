package output

import (
	"testing"
)

func InsertAndLoad(t *testing.T) {
	db := MappingDB{}
	db.Save("line 1")
	db.Save("line 2")
	array := db.FindAll()

	if len(array) != 2 {
		t.Errorf("Wrong Size array should be %d but is %d", 2, len(array))
	}
}
