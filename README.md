


### what happens:

  * keys and values are always `string`. This is sad, but the truth.
  * When the value for some key is "" or blank or not found, `NotFound` will be returned.
  * `RangeOpts` treats go zero values as if you didn't pass anything, so `Limit: 0` is the same as no limit, `Start: ""` is the same as no start etc.
  * In `RangeOpts` `Start` is inclusive, `End` is non-inclusive.
