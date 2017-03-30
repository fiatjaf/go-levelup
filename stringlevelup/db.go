package stringlevelup

import "github.com/fiatjaf/levelup"

type DB struct {
	levelup.DB
}

func StringDB(db levelup.DB) DB {
	return DB{db}
}

func (db DB) Put(key string, value string) error {
	return db.DB.Put([]byte(key), []byte(value))
}

func (db DB) Get(key string) (string, error) {
	value, err := db.DB.Get([]byte(key))
	return string(value), err
}

func (db DB) Del(key string) error {
	return db.DB.Del([]byte(key))
}

func (db DB) ReadRange(opts *RangeOpts) ReadIterator {
	var upopts levelup.RangeOpts
	if opts != nil {
		upopts = levelup.RangeOpts{
			Start:   []byte(opts.Start),
			End:     []byte(opts.End),
			Reverse: opts.Reverse,
			Limit:   opts.Limit,
		}
	}
	upiter := db.DB.ReadRange(&upopts)
	return ReadIterator{upiter}
}

type ReadIterator struct {
	levelup.ReadIterator
}

func (ri ReadIterator) Key() string {
	return string(ri.ReadIterator.Key())
}

func (ri ReadIterator) Value() string {
	return string(ri.ReadIterator.Value())
}
