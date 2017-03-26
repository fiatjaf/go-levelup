package levelup

import (
	"testing"
)

func BasicTests(db DB, t *testing.T) {
	value, err := db.Get("key-x")
	if err != NotFound {
		t.Error("NotFound error should happen: ", err)
	}
	if value != "" {
		t.Error("value should be a blank string: ", value)
	}
}
