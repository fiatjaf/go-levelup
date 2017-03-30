
### some things:

  * When the value for some key is "" or blank or not found, `NotFound` will be returned.
  * `RangeOpts` treats go zero values as if you didn't pass anything, so `Limit: 0` is the same as no limit, `Start: ""` is the same as no start etc.
  * In `RangeOpts` `Start` is inclusive, `End` is non-inclusive.
  * All keys and values are `[]byte`, if you want to use strings, import `github.com/fiatjaf/levelup/stringlevelup` and call `stringlevelup.StringDB()` on your `levelup.DB` object. It will return an exact copy of the `levelup.DB`, only it will use `string` for keys and values everywhere.
