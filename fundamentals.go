package levelup

import "errors"

type DB interface {
	Put(string, string) error
	Get(string) (string, error)
	Del(string) error
	Batch([]Operation) error
	ReadRange(RangeOpts) ReadIterator
}

func OpPut(key, value string) Operation { return Operation{"type": "put", "key": key, "value": value} }
func OpDel(key string) Operation        { return Operation{"type": "del", "key": key} }

type Operation map[string]string

type RangeOpts struct {
	Start   string // included
	End     string // not included
	Reverse bool
	Limit   int
}

type ReadIterator interface {
	Next() bool
	Key() string
	Value() string
	Error() error
	Release()
}

var (
	NotFound = errors.New("not found")
)
