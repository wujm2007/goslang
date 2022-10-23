package option

type Option[T any] interface {
	Get() (*T, bool)
	MustGet() T

	OrElse(T) T

	IsNil() bool
	Error() error
}

type option[T any] struct {
	val *T
	err error
}

func (o option[T]) Get() (*T, bool) {
	return o.val, o.val != nil
}

func (o option[T]) MustGet() T {
	if o.IsNil() {
		panic("nil")
	}
	return *o.val
}

func (o option[T]) IsNil() bool {
	return o.val == nil
}

func (o option[T]) Error() error {
	return o.err
}

func (o option[T]) OrElse(elseV T) T {
	if o.IsNil() {
		return elseV
	}
	return *o.val
}

func OfE[T any](v *T, err error) Option[T] {
	if (v != nil) != (err == nil) {
		panic("either val==nil or error!=nil")
	}
	return &option[T]{val: v, err: err}
}

func OfFuncE[T any](f func() (v *T, err error)) Option[T] {
	return OfE(f())
}

func Of[T any](v T) Option[T] {
	return &option[T]{val: &v}
}

func OfNil[T any]() Option[T] {
	return &option[T]{}
}

func OfNillable[T any](v *T) Option[T] {
	return &option[T]{val: v}
}

func Map[O Option[T], T, R any](o O, mapper func(T) R) Option[R] {
	if o.IsNil() {
		return OfNil[R]()
	}
	return Of[R](mapper(o.MustGet()))
}

func MapE[O Option[T], T, R any](o O, mapper func(T) (*R, error)) Option[R] {
	if o.IsNil() {
		return OfNil[R]()
	}
	return OfFuncE(func() (*R, error) { return mapper(o.MustGet()) })
}

func BinaryOp[O Option[T], T, R any](o1, o2 O, op func(v1, v2 T) R) Option[R] {
	if o1.IsNil() || o2.IsNil() {
		return OfNil[R]()
	}
	return Of(op(o1.MustGet(), o2.MustGet()))
}

func BinaryOpE[O Option[T], T, R any](o1, o2 O, op func(v1, v2 T) (*R, error)) Option[R] {
	if o1.IsNil() || o2.IsNil() {
		return OfNil[R]()
	}
	return OfE(op(o1.MustGet(), o2.MustGet()))
}
