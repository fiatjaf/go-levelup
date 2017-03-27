package levelup

import "errors"

type DB interface {
	Put(string, string) error
	Get(string) (string, error)
	Del(string) error
	Batch([]Operation) error
	ReadRange(*RangeOpts) ReadIterator
}

func Put(key, value string) Operation { return Operation{"type": "put", "key": key, "value": value} }
func Del(key string) Operation        { return Operation{"type": "del", "key": key} }

var BatchPut = Put
var BatchDel = Del

type Operation map[string]string

type RangeOpts struct {
	Start   string // included, default ""
	End     string // not included, default "~~~~~"
	Reverse bool   // default false
	Limit   int    // default 9999999
}

const (
	DefaultRangeLimit = 9999999
	DefaultRangeEnd   = "~~~~~"
)

func (ro *RangeOpts) FillDefaults() {
	if ro.Limit == 0 {
		ro.Limit = DefaultRangeLimit
	}
	if ro.End == "" {
		ro.End = DefaultRangeEnd
	}
}

type ReadIterator interface {
	Valid() bool   // returns false when we have reached the end of the rows we asked for.
	Next()         // go to the next row.
	Key() string   // the key in the current row.
	Value() string // the value in the current row.
	Error() error  // if some error has happened this have it.
	Release()      // it may be necessary to call this after using, or defer it.
}

var (
	NotFound = errors.New("not found")
)
