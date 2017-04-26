[![](https://godoc.org/github.com/fiatjaf/levelup/stringlevelup?status.svg)](http://godoc.org/github.com/fiatjaf/levelup/stringlevelup)

## how to use

Just call `stringlevelup.StringDB()` on your `levelup.DB` object (from [github.com/fiatjaf/levelup](https://github.com/fiatjaf/levelup)). It will return an exact copy of the `levelup.DB`, the only difference is that it will use `string` for keys and values everywhere.
