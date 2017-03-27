package levelup_test

import (
	"testing"

	"github.com/fiatjaf/levelup"
	. "gopkg.in/check.v1"
)

var db levelup.DB
var err error

func Test(sdb levelup.DB, t *testing.T) {
	db = sdb
	TestingT(t)
}

type BasicSuite struct{}

var _ = Suite(&BasicSuite{})

// the tests must be run in order. yes.
// each suite depends on the the previous, that's why they have numbers in their names.

func (s *BasicSuite) Test1PutGetDel(c *C) {
	value, err := db.Get("key-x")
	c.Assert(err, DeepEquals, levelup.NotFound)
	c.Assert(value, DeepEquals, "")

	err = db.Put("key-x", "some value")
	c.Assert(err, IsNil)
	value, _ = db.Get("key-x")
	c.Assert(value, DeepEquals, "some value")

	err = db.Del("key-x")
	c.Assert(err, IsNil)
	value, err = db.Get("key-x")
	c.Assert(err, DeepEquals, levelup.NotFound)
	c.Assert(value, DeepEquals, "")
}

func (s *BasicSuite) Test2BatchPut(c *C) {
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
	c.Assert(err, IsNil)

	iter := db.ReadRange(nil)
	retrieved := []string{}
	for ; iter.Valid(); iter.Next() {
		c.Assert(iter.Error(), IsNil)
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	c.Assert(iter.Error(), IsNil)
	c.Assert(retrieved, DeepEquals, []string{
		"letter:a", "a",
		"letter:b", "b",
		"letter:c", "c",
		"number:1", "1",
		"number:2", "2",
		"number:3", "3",
	})
	iter.Release()
}

func (s *BasicSuite) Test3ReadRange(c *C) {
	// start-end
	iter := db.ReadRange(&levelup.RangeOpts{
		Start: "letter:b",
		End:   "letter:~",
	})
	retrieved := []string{}
	for ; iter.Valid(); iter.Next() {
		c.Assert(iter.Error(), IsNil)
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	c.Assert(iter.Error(), IsNil)
	c.Assert(retrieved, DeepEquals, []string{"letter:b", "b", "letter:c", "c"})
	iter.Release()

	// *-end
	iter = db.ReadRange(&levelup.RangeOpts{
		End: "letter:c", /* non-inclusive */
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		c.Assert(iter.Error(), IsNil)
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	c.Assert(iter.Error(), IsNil)
	c.Assert(retrieved, DeepEquals, []string{"letter:a", "a", "letter:b", "b"})
	iter.Release()

	// start-* limit
	iter = db.ReadRange(&levelup.RangeOpts{
		Start: "letter:c",
		Limit: 2,
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		c.Assert(iter.Error(), IsNil)
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	c.Assert(iter.Error(), IsNil)
	c.Assert(retrieved, DeepEquals, []string{"letter:c", "c", "number:1", "1"})
	iter.Release()

	// reverse
	iter = db.ReadRange(&levelup.RangeOpts{
		Reverse: true,
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		c.Assert(iter.Error(), IsNil)
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	c.Assert(iter.Error(), IsNil)
	c.Assert(retrieved, DeepEquals, []string{
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
		c.Assert(iter.Error(), IsNil)
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	c.Assert(iter.Error(), IsNil)
	c.Assert(retrieved, DeepEquals, []string{"number:1", "1", "letter:c", "c"})
	iter.Release()

	// reverse *-end limit
	iter = db.ReadRange(&levelup.RangeOpts{
		End:     "number:3", /* non-inclusive */
		Reverse: true,
		Limit:   3,
	})
	retrieved = []string{}
	for ; iter.Valid(); iter.Next() {
		c.Assert(iter.Error(), IsNil)
		retrieved = append(retrieved, iter.Key(), iter.Value())
	}
	c.Assert(iter.Error(), IsNil)
	c.Assert(retrieved, DeepEquals, []string{"number:2", "2", "number:1", "1", "letter:c", "c"})
	iter.Release()
}

func (s *BasicSuite) Test4MoreBatches(c *C) {
	batch := []levelup.Operation{
		levelup.Del("number:2"),
		levelup.Del("number:1"),
		levelup.Put("number:3", "33"),
		levelup.Del("number:4"),
		levelup.Del("letter:a"),
		levelup.Del("number:3"),
		levelup.Del("letter:b"),
		levelup.Del("letter:c"),
		levelup.Put("number:3", "333"),
		levelup.Del("letter:d"),
		levelup.Put("letter:d", "dd"),
		levelup.Del("letter:e"),
	}
	err = db.Batch(batch)
	c.Assert(err, IsNil)

	value, err := db.Get("number:1")
	c.Assert(err, DeepEquals, levelup.NotFound)
	c.Assert(value, DeepEquals, "")

	value, err = db.Get("letter:e")
	c.Assert(err, DeepEquals, levelup.NotFound)
	c.Assert(value, DeepEquals, "")

	value, _ = db.Get("number:3")
	c.Assert(value, DeepEquals, "333")

	value, _ = db.Get("letter:d")
	c.Assert(value, DeepEquals, "dd")
}
