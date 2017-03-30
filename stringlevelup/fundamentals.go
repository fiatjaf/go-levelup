package stringlevelup

import (
	"github.com/fiatjaf/levelup"
)

func Put(key, value string) levelup.Operation { return levelup.Put([]byte(key), []byte(value)) }
func Del(key string) levelup.Operation        { return levelup.Del([]byte(key)) }

var BatchPut = Put
var BatchDel = Del

type RangeOpts struct {
	Start   string // included, default ""
	End     string // not included, default "~~~~~"
	Reverse bool   // default false
	Limit   int    // default 9999999
}
