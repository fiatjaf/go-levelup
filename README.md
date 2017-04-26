[![](https://godoc.org/github.com/fiatjaf/levelup?status.svg)](http://godoc.org/github.com/fiatjaf/levelup)

# levelup

This package implements a few types and interfaces to be used by the actual "down" libraries, which will implement the core functionality of each database.

Some of the "down" libraries implementing the "up" interface:

  * [goleveldown](https://github.com/fiatjaf/goleveldown)
  * [rocksdown](https://github.com/fiatjaf/rocksdown)
  * [levelup-js](https://github.com/fiatjaf/levelup-js)

Any similarity with the [Node.js LevelUP](https://github.com/Level/levelup) is mere coincidence.

### some things:

  * When the value for some key is "" or blank or not found, `NotFound` will be returned.
  * `RangeOpts` treats go zero values as if you didn't pass anything, so `Limit: 0` is the same as no limit, `Start: ""` is the same as no start etc.
  * In `RangeOpts` `Start` is inclusive, `End` is non-inclusive.
  * All keys and values are `[]byte`, if you want to use strings, see [github.com/fiatjaf/levelup/stringlevelup](https://github.com/fiatjaf/levelup/tree/master/stringlevelup).
