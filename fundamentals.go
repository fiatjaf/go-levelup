package levelup

import "errors"

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

var (
	NotFound = errors.New("not found")
)
