package levelup

type DB interface {
	Close()
	Erase()
	Put([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Del([]byte) error
	Batch([]Operation) error
	ReadRange(*RangeOpts) ReadIterator
}

type ReadIterator interface {
	Valid() bool   // returns false when we have reached the end of the rows we asked for.
	Next()         // go to the next row.
	Key() []byte   // the key in the current row.
	Value() []byte // the value in the current row.
	Error() error  // if some error has happened this have it.
	Release()      // it may be necessary to call this after using, or defer it.
}
