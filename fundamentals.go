package levelup

type Operation struct {
	Type  string
	Key   []byte
	Value []byte
}

func Put(key, value []byte) Operation { return Operation{"put", key, value} }
func Del(key []byte) Operation        { return Operation{"del", key, []byte{}} }

var BatchPut = Put
var BatchDel = Del

type RangeOpts struct {
	Start   []byte // included, default []byte{}
	End     []byte // not included, default []byte{0xff, 0xff, 0xff, 0xff, 0xff}
	Reverse bool   // default false
	Limit   int    // default 9999999
}

var (
	DefaultRangeLimit = 9999999
	DefaultRangeEnd   = []byte{0xff, 0xff, 0xff, 0xff, 0xff}
)

func (ro *RangeOpts) FillDefaults() {
	if ro.Limit == 0 {
		ro.Limit = DefaultRangeLimit
	}
	if len(ro.End) == 0 {
		ro.End = DefaultRangeEnd
	}
}

type Error string

func (e Error) Error() string { return string(e) }

const (
	NotFound = Error("levelup: not found")
)
