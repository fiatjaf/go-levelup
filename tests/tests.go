package levelup_test

import (
	"testing"

	"github.com/facebookgo/ensure"
	"github.com/fiatjaf/levelup"
)

func BasicTests(db levelup.DB, t *testing.T) {
	/*** basic put, get, del ***/
	value, err := db.Get("key-x")
	ensure.DeepEqual(t, err, levelup.NotFound)
	ensure.DeepEqual(t, value, "")

	err = db.Put("key-x", "some value")
	ensure.Nil(t, err)
	value, _ = db.Get("key-x")
	ensure.DeepEqual(t, value, "some value")

	err = db.Del("key-x")
	ensure.Nil(t, err)
	value, err = db.Get("key-x")
	ensure.DeepEqual(t, err, levelup.NotFound)
	ensure.DeepEqual(t, value, "")

	/*** batch put ***/
	somevalues := map[string]string{
		"letter:a": "a",
		"letter:b": "b",
		"letter:c": "c",
		"number:1": "1",
		"number:2": "2",
		"number:3": "3",
	}
	batch := []levelup.Operation{}
	for k, v := range somevalues {
		batch = append(batch, levelup.Put(k, v))
	}
	err = db.Batch(batch)
	ensure.Nil(t, err)

	/*** reading ranges ***/
	// start-end
	iter := db.ReadRange(&levelup.RangeOpts{
		Start: "letter:b",
		End:   "letter:~",
	})
	retrieved := []string{}
	for ; iter.Valid(); iter.Next() {
		ensure.Nil(t, iter.Error())
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	ensure.DeepEqual(t, retrieved, []string{"letter:b", "b", "letter:c", "c"})
	iter.Release()

	// *-end
	iter = db.ReadRange(&levelup.RangeOpts{
		End: "letter:c", /* non-inclusive */
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		ensure.Nil(t, iter.Error())
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	ensure.DeepEqual(t, retrieved, []string{"letter:a", "a", "letter:b", "b"})
	iter.Release()

	// start-* limit
	iter = db.ReadRange(&levelup.RangeOpts{
		Start: "letter:c",
		Limit: 2,
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		ensure.Nil(t, iter.Error())
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	ensure.DeepEqual(t, retrieved, []string{"letter:c", "c", "number:1", "1"})
	iter.Release()

	// reverse
	iter = db.ReadRange(&levelup.RangeOpts{
		Reverse: true,
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		ensure.Nil(t, iter.Error())
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	ensure.DeepEqual(t, retrieved, []string{
		"number:3", "3", "number:2", "2", "number:1", "1",
		"letter:c", "c", "letter:b", "b", "letter:a", "a",
	})
	iter.Release()

	// reverse start-end
	iter = db.ReadRange(&levelup.RangeOpts{
		Start:   "letter:c",
		End:     "number:1~",
		Reverse: true,
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		ensure.Nil(t, iter.Error())
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	ensure.DeepEqual(t, retrieved, []string{"number:1", "1", "letter:c", "c"})
	iter.Release()

	// reverse *-end limit
	iter = db.ReadRange(&levelup.RangeOpts{
		End:     "number:3", /* non-inclusive */
		Reverse: true,
		Limit:   3,
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		ensure.Nil(t, iter.Error())
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	ensure.DeepEqual(t, retrieved, []string{"number:2", "2", "number:1", "1", "letter:c", "c"})
	iter.Release()
}
