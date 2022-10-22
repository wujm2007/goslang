package option

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type s struct {
	intp *int
	str  string
}

func TestOf(t *testing.T) {

	get := func() (*int, error) { return new(int), nil }
	getE := func() (*int, error) { return nil, errors.New("foo") }

	got, ok := OfNil[int]().Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, got)

	got, ok = OfErrorFunc(getE).Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, got)

	gotS, ok := OfNil[s]().Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, gotS)

	gotIntS, ok := OfNil[[]int]().Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, gotIntS)

	gotMap, ok := OfNil[map[string]int]().Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, gotMap)

	var nilInt *int
	got, ok = OfNillable[int](nilInt).Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, got)

	intVal := 1
	got, ok = OfNillable(&intVal).Get()
	assert.Equal(t, true, ok)
	assert.NotNil(t, got)
	assert.Equal(t, intVal, *got)

	var nilS *s
	gotS, ok = OfNillable(nilS).Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, gotS)

	var nilIntS *[]int
	gotIntS, ok = OfNillable(nilIntS).Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, gotIntS)

	var nilMap *map[string]int
	gotMap, ok = OfNillable(nilMap).Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, gotMap)

	got, ok = Of(1).Get()
	assert.Equal(t, true, ok)
	assert.NotNil(t, got)
	assert.Equal(t, 1, *got)

	gotS, ok = Of(s{nil, "foo"}).Get()
	assert.Equal(t, true, ok)
	assert.NotNil(t, got)
	assert.Equal(t, s{nil, "foo"}, *gotS)

	gotIntS, ok = Of([]int{1, 2, 3}).Get()
	assert.Equal(t, true, ok)
	assert.NotNil(t, got)
	assert.Equal(t, []int{1, 2, 3}, *gotIntS)

	gotMap, ok = Of(map[string]int{"foo": 2}).Get()
	assert.Equal(t, true, ok)
	assert.NotNil(t, got)
	assert.Equal(t, map[string]int{"foo": 2}, *gotMap)

	got, ok = OfErrorFunc(get).Get()
	assert.Equal(t, true, ok)
	assert.NotNil(t, got)
	assert.Equal(t, 0, *got)
}

func TestMap(t *testing.T) {
	mapper := func(t int) string { return fmt.Sprintf("%d", t) }
	mapperE := func(t int) (*string, error) { return nil, errors.New("bar") }

	m, ok := Map(OfNil[int](), mapper).Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, m)

	m, ok = Map(Of(123), mapper).Get()
	assert.Equal(t, true, ok)
	assert.NotNil(t, m)
	assert.Equal(t, "123", *m)

	m, ok = MapErrorFunc(Of(123), mapperE).Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, m)

	err := MapErrorFunc(Of(123), mapperE).Error()
	assert.NotNil(t, err)
	assert.Equal(t, "bar", err.Error())

	err = MapErrorFunc(OfNil[int](), mapperE).Error()
	assert.Nil(t, m)
}

func Test_option_Get(t *testing.T) {
	got, ok := OfNil[int]().Get()
	assert.Equal(t, false, ok)
	assert.Nil(t, got)

	got, ok = Of(1).Get()
	assert.Equal(t, true, ok)
	assert.NotNil(t, got)
	assert.Equal(t, 1, *got)
}

func Test_option_IsNil(t *testing.T) {
	isNil := OfNil[int]().IsNil()
	assert.Equal(t, true, isNil)

	isNil = Of(1).IsNil()
	assert.Equal(t, false, isNil)
}

func Test_option_MustGet(t *testing.T) {
	assert.Panics(t, func() {
		OfNil[int]().MustGet()
	})

	assert.NotPanics(t, func() {
		i := Of(1).MustGet()
		assert.Equal(t, 1, i)
	})
}

func Test_option_OrElse(t *testing.T) {
	got := Of(1).OrElse(2)
	assert.Equal(t, 1, got)

	got = OfNil[int]().OrElse(2)
	assert.Equal(t, 2, got)
}
