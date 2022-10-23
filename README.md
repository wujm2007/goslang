# goslang

[![Code Coverage](https://codecov.io/gh/wujm2007/goslang/branch/main/graph/badge.svg)](https://codecov.io/gh/wujm2007/goslang)

## Using goslang

### `Option[T]`
Example:
```go
package main

import (
	"errors"
	"fmt"

	"github.com/wujm2007/goslang/option"
)

func main() {
	o1 := option.Of(1) // Option[int] with value=1
	o1.Get()           // returns (*int)(1), true
	o1.MustGet()       // returns 1

	int2str := func(t int) string { return fmt.Sprintf("%d", t) }
	option.Map(o1, int2str) // Option[string] with value="1"
	
	badMapper := func(t int) (*string, error) { return nil, errors.New("boom") }
	o2 := option.MapE(o1, badMapper) // Option[string] with nil value & boom error
	o2.Get() // returns (*int)(nil), false
	o2.Error() // returns error="boom"
	
	o3 := option.OfNil[int]() // Option[int] with nil value
	o3.Get()                  // returns (*int)(nil), false
	o3.MustGet()              // panic!
	o3.OrElse(1)              // Option[int] with value=1
	option.Map(o3, int2str)   // Option[string] with nil value

	intVal := 1
	_ = option.OfNillable(&intVal) // Option[int] with value=1
	
	getter := func() (*int, error) { return new(int), nil }
	_ = option.OfFuncE(getter) // Option[int] with value=0
	
	badGetter := func() (*int, error) { return nil, errors.New("boom") }
	_ = option.OfFuncE(badGetter) // Option[int] with nil value & boom error
}
```